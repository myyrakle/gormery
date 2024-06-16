package steps

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"path"
	"reflect"
	"strings"

	"github.com/myyrakle/gormery/internal/annotation"
	config "github.com/myyrakle/gormery/internal/config"
	"github.com/myyrakle/gormery/pkg/cast"
	slice "github.com/myyrakle/gormery/pkg/slice"
)

func ReadAllTargets(configFile config.ConfigFile) ProecssFileContexts {
	contexts := readFileRecursive(configFile.Basedir, configFile)

	return contexts
}

type ProcessFileField struct {
	fieldName       string
	bsonName        string
	isPointer       bool
	typePackageName string
	typeName        string
	comment         *string
}

type ProecssFileContext struct {
	packageName string
	file        *ast.File
	filename    string
	structName  string
	entityParam *string
	fields      []ProcessFileField
}

type ProecssFileContexts []ProecssFileContext

func (c ProecssFileContexts) UniquedFileNames() []string {
	uniqueFileName := make(slice.Strings, 0, len(c))

	for _, context := range c {
		if !uniqueFileName.Contains(context.filename) {
			uniqueFileName = append(uniqueFileName, context.filename)
		}
	}

	return uniqueFileName
}

func getDirList(basePath string) []string {
	dirs, err := os.ReadDir(basePath)
	if err != nil {
		log.Fatal(err)
	}

	var dirList []string
	for _, dir := range dirs {
		if dir.IsDir() {
			dirList = append(dirList, dir.Name())
		}
	}

	return dirList
}

// 단일 파일을 읽어서 형식화하는 단위 함수입니다.
func readFile(_ config.ConfigFile, packageName string, filename string, file *ast.File) []ProecssFileContext {
	contexts := make([]ProecssFileContext, 0)

	for _, declare := range file.Decls {
		if genDecl, ok := declare.(*ast.GenDecl); ok {
			for _, spec := range genDecl.Specs {
				if typeSpec, ok := spec.(*ast.TypeSpec); ok {
					structDecl, _ := typeSpec.Type.(*ast.StructType)

					if structDecl == nil {
						continue
					}

					if !isEntityStruct(genDecl) {
						continue
					}

					entityParam := getEntityParam(genDecl)

					structName := typeSpec.Name.Name

					processFileContext := ProecssFileContext{
						packageName: packageName,
						file:        file,
						filename:    filename,
						structName:  structName,
						entityParam: entityParam,
					}

					// 구조체 필드를 순회하면서 필요한 정보를 추출합니다.
					for _, field := range structDecl.Fields.List {
						processFileField := convertFieldToProcessFileField(structName, packageName, field)

						if processFileField != nil {
							processFileContext.fields = append(processFileContext.fields, *processFileField)
						}
					}

					contexts = append(contexts, processFileContext)
				}
			}
		}
	}

	return contexts
}

// 패키지 목록을 가져옵니다.
func getPackageList(basePath string) map[string]*ast.Package {
	fset := token.NewFileSet()

	packages, err := parser.ParseDir(fset, basePath, nil, parser.ParseComments)

	if err != nil {
		panic(err)
	}

	return packages
}

// 필드 정보를 받아서 ProcessFileField로 변환합니다.
func convertFieldToProcessFileField(_ string, packageName string, field *ast.Field) *ProcessFileField {
	processFileField := ProcessFileField{}

	if len(field.Names) == 0 {
		return nil
	}

	name := field.Names[0].Name
	processFileField.fieldName = name

	if field.Tag == nil {
		return nil
	}

	tag := strings.ReplaceAll(field.Tag.Value, "`", "")

	bson := reflect.StructTag(tag).Get("bson")

	if bson == "" {
		return nil
	}

	if bson == "-" {
		return nil
	}

	bsonTokens := strings.Split(bson, ",")
	bsonName := bsonTokens[0]

	processFileField.bsonName = bsonName

	if field.Type == nil {
		return nil
	}

	if field.Comment != nil {
		comment := field.Comment.Text()
		processFileField.comment = cast.ToPointer(comment)
	}

	// 필드 타입이 포인터인 경우
	if starExpr, ok := field.Type.(*ast.StarExpr); ok {
		processFileField.isPointer = true

		// 패키지가 명시되어 있는 경우
		if selectorExpr, ok := starExpr.X.(*ast.SelectorExpr); ok {
			if xIdent, ok := selectorExpr.X.(*ast.Ident); ok {
				processFileField.typePackageName = xIdent.Name
				processFileField.typeName = selectorExpr.Sel.Name
			}
		} else /* 패키지가 명시되어있지 않은 경우 */ if ident, ok := starExpr.X.(*ast.Ident); ok {
			processFileField.typePackageName = packageName
			processFileField.typeName = ident.Name
		}
	} else /* 필드 타입이 non-pointer에 패키지 명시가 있는 경우 */ if selectorExpr, ok := field.Type.(*ast.SelectorExpr); ok {
		if xIdent, ok := selectorExpr.X.(*ast.Ident); ok {
			processFileField.typePackageName = xIdent.Name
			processFileField.typeName = selectorExpr.Sel.Name
		} else {
			panic("unexpected error")
		}
	} else /* 필드 타입이 non-pointer에 패키지 명시도 없는 경우 */ if ident, ok := field.Type.(*ast.Ident); ok {
		processFileField.typeName = ident.Name
		processFileField.typePackageName = packageName
	}

	return &processFileField
}

func readFileRecursive(basedir string, configFile config.ConfigFile) []ProecssFileContext {
	contexts := make([]ProecssFileContext, 0)

	packages := getPackageList(basedir)

	for packageName, asts := range packages {
		for filename, file := range asts.Files {
			if strings.HasSuffix(filename, "_test.go") {
				continue
			}

			fmt.Printf(">> scan [%s]...\n", filename)
			eachContexts := readFile(configFile, packageName, filename, file)
			contexts = append(contexts, eachContexts...)
		}
	}

	dirList := getDirList(basedir)

	for _, dir := range dirList {
		eachContexts := readFileRecursive(path.Join(basedir, dir), configFile)
		contexts = append(contexts, eachContexts...)
	}

	return contexts
}

// 주석을 읽어와서 @Gorm 구조체인지 검증합니다.
func isEntityStruct(genDecl *ast.GenDecl) bool {
	if genDecl.Doc == nil {
		return false
	}

	if genDecl.Doc.List == nil {
		return false
	}

	for _, comment := range genDecl.Doc.List {
		if strings.Contains(comment.Text, "@Gorm") {
			return true
		}
	}

	return false
}

// @Gorm의 파라미터를 가져옵니다.
func getEntityParam(genDecl *ast.GenDecl) *string {
	if genDecl.Doc == nil {
		return nil
	}

	if genDecl.Doc.List == nil {
		return nil
	}

	for _, comment := range genDecl.Doc.List {
		if strings.Contains(comment.Text, "@Gorm") {
			params := annotation.ParseParameters(comment.Text)

			if len(params) > 0 {
				return cast.ToPointer(params[0])
			}
		}
	}

	return nil
}
