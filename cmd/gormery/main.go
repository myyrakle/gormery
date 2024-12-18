package main

import (
	 "sync"
	 "os"
	 "fmt"
	 "strings"

	target "github.com/myyrakle/gormery/example"
	gormSchema "gorm.io/gorm/schema"
)

func main() {
	code := ""
	gormFilePath0 := "example/order_gorm.go"
	f0, err := os.Create(gormFilePath0)
	if err != nil {
		panic(err)
	}
	code = ""
	code += "// Code generated by gormery. DO NOT EDIT.\n"
	code += "package " + "models" + "\n\n"
	_, err = f0.WriteString(code)
	if err != nil {
		panic(err)
	}
	f0.Close()
	gormFilePath1 := "example/person_gorm.go"
	f1, err := os.Create(gormFilePath1)
	if err != nil {
		panic(err)
	}
	code = ""
	code += "// Code generated by gormery. DO NOT EDIT.\n"
	code += "package " + "models" + "\n\n"
	_, err = f1.WriteString(code)
	if err != nil {
		panic(err)
	}
	f1.Close()
	gormFilePath2 := "example/clothes_gorm.go"
	f2, err := os.Create(gormFilePath2)
	if err != nil {
		panic(err)
	}
	code = ""
	code += "// Code generated by gormery. DO NOT EDIT.\n"
	code += "package " + "models" + "\n\n"
	_, err = f2.WriteString(code)
	if err != nil {
		panic(err)
	}
	f2.Close()

	target_0, err := gormSchema.ParseWithSpecialTableName(
		&target.Order{},
		&sync.Map{},
		&gormSchema.NamingStrategy{},
		"",
	)

	if err == nil {
		createGormFile(target_0, "example/order.go", "Order", "order__")
	}


	target_1, err := gormSchema.ParseWithSpecialTableName(
		&target.Person{},
		&sync.Map{},
		&gormSchema.NamingStrategy{},
		"",
	)

	if err == nil {
		createGormFile(target_1, "example/person.go", "Person", "")
	}


	target_2, err := gormSchema.ParseWithSpecialTableName(
		&target.PersonSoMany{},
		&sync.Map{},
		&gormSchema.NamingStrategy{},
		"",
	)

	if err == nil {
		createGormFile(target_2, "example/person.go", "PersonSoMany", "")
	}


	target_3, err := gormSchema.ParseWithSpecialTableName(
		&target.PackingClothes{},
		&sync.Map{},
		&gormSchema.NamingStrategy{},
		"",
	)

	if err == nil {
		createGormFile(target_3, "example/clothes.go", "PackingClothes", "")
	}

}

var basedir = "example"
var outputSuffix = "_gorm.go"
func createGormFile(schema *gormSchema.Schema, filename string, structName string, tableName string) {
	gormFilePath := strings.Replace(filename, ".go", "", 1) + outputSuffix
	code := ""
	code += "func (t " + structName + ") TableName() string {\n"
	if tableName != "" {
	code += "\treturn \""+ tableName + "\"\n"
	} else {
	code += "\treturn \"" + schema.Table + "\"\n"
	}
	code += "}\n\n"

	code += "func (t " + structName + ") StructName() string {\n"
	code += "\treturn \"" + structName + "\"\n"
	code += "}\n\n"

	columnConstantNames := []string{}
	for _, field := range schema.Fields {
		columnConstantName := structName + "_" + field.Name
		columnConstantExpression := "const " + columnConstantName + " = " + "\"" + field.DBName + "\"" + "\n"
		columnConstantNames = append(columnConstantNames, "\t\t"+columnConstantName+",")
		code += columnConstantExpression
	}

	code += "\nfunc (t " + structName + ") Columns() []string {\n"
	code += "\treturn []string{\n" + strings.Join(columnConstantNames, "\n") + "\n\t}\n"
	code += "}\n\n"

	sliceTypeName := gormSchema.NamingStrategy{ NoLowerCase: true }.TableName(structName)
	if sliceTypeName == structName {
		sliceTypeName += "List"
	}
	code += "type " + sliceTypeName + " []" + structName + "\n\n"
	code += "func (t " + sliceTypeName + ") Len() int {\n"
	code += "\treturn len(t)\n"
	code += "}\n\n"
	code += "func (t " + sliceTypeName + ") IsEmpty() bool {\n"
	code += "\treturn len(t) == 0\n"
	code += "}\n\n"
	code += "func (t " + sliceTypeName + ") First() " + structName + " {\n"
	code += "\tif t.IsEmpty() {\n"
	code += "\t\treturn " + structName + "{}\n"
	code += "\t}\n"
	code += "\treturn t[0]\n"
	code += "}\n\n"
	f, err := os.OpenFile(gormFilePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	_, err = fmt.Fprintln(f, code)
}