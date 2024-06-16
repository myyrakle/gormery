package steps

import (
	"fmt"
	"os"
	"strings"

	config "github.com/myyrakle/gormery/internal/config"
)

func GenerateRunner(configFile config.ConfigFile, targets ProecssFileContexts) {
	// Runner-Path 디렉토리가 존재하지 않는다면 생성
	if _, err := os.Stat(configFile.RunnerPath); os.IsNotExist(err) {
		os.MkdirAll(configFile.RunnerPath, 0755)
	}

	var code string

	code += "package main\n\n"

	code += "import (\n"

	code += "\t " + `"sync"` + "\n"
	code += "\t " + `"os"` + "\n"
	code += "\t " + `"fmt"` + "\n"
	code += "\t " + `"strings"` + "\n\n"

	targetImport := `target "` + configFile.ModuleName + "/" + configFile.Basedir + `"`
	code += "\t" + targetImport + "\n"

	gormImport := `gormSchema "gorm.io/gorm/schema"`
	code += "\t" + gormImport + "\n"

	code += ")\n\n"

	code += "func main() {\n"

	for i, target := range targets {
		code += generateCodeForTarget(i, target)
		code += "\n"
	}

	code += "}\n\n"

	// Gorm 파일 생성 코드 주입
	code += generateCreateGormFileFunction(configFile)

	// 파일 생성
	filePath := configFile.RunnerPath + "/main.go"
	file, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	_, err = file.WriteString(code)
	if err != nil {
		panic(err)
	}

	// _gorm.go 파일 미리 생성
	filenames := targets.UniquedFileNames()
	for _, filename := range filenames {
		gormFilePath := strings.Replace(filename, ".go", "", 1) + configFile.OutputSuffix
		f, err := os.Create(gormFilePath)
		if err != nil {
			panic(err)
		}
		defer f.Close()

		code := ""
		code += "package " + targets[0].packageName + "\n\n"

		_, err = f.WriteString(code)
		if err != nil {
			panic(err)
		}
	}
}

func generateCodeForTarget(i int, target ProecssFileContext) string {
	id := fmt.Sprintf("target_%d", i)
	targetTypename := target.structName
	filename := target.filename
	structName := target.structName

	code := fmt.Sprintf(`
	%s, err := gormSchema.ParseWithSpecialTableName(
		&target.%s{},
		&sync.Map{},
		&gormSchema.NamingStrategy{},
		"",
	)

	if err == nil {
		createGormFile(%s, "%s", "%s")
	}
`, id, targetTypename, id, filename, structName)

	return code
}

func generateCreateGormFileFunction(configFile config.ConfigFile) string {
	code := ""

	code += `var basedir = ` + `"` + configFile.Basedir + `"` + "\n"
	code += `var outputSuffix = ` + `"` + configFile.OutputSuffix + `"` + "\n"

	code += `func createGormFile(schema *gormSchema.Schema, filename string, structName string) {` + "\n"
	code += "\t" + `gormFilePath := strings.Replace(filename, ".go", "", 1) + outputSuffix` + "\n"

	code += "\t" + `code := ""` + "\n"

	code += "\t" + `code += "func (t " + structName + ") TableName() string {\n"` + "\n"
	code += "\t" + `code += "\treturn \"" + schema.Table + "\"\n"` + "\n"
	code += "\t" + `code += "}\n"` + "\n"

	code += "\t" + `f, err := os.OpenFile(gormFilePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)` + "\n"
	code += "\t" + `if err != nil {` + "\n"
	code += "\t\t" + `panic(err)` + "\n"
	code += "\t}\n"
	code += "\tdefer f.Close()\n"
	code += "\t" + `_, err = fmt.Fprintln(f, code)` + "\n"
	code += `}`

	return code
}
