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

	target_0, err := gormSchema.ParseWithSpecialTableName(
		&target.Person{},
		&sync.Map{},
		&gormSchema.NamingStrategy{},
		"",
	)

	if err == nil {
		createGormFile(target_0, "example/person.go", "Person")
	}


	target_1, err := gormSchema.ParseWithSpecialTableName(
		&target.PersonSoMany{},
		&sync.Map{},
		&gormSchema.NamingStrategy{},
		"",
	)

	if err == nil {
		createGormFile(target_1, "example/person.go", "PersonSoMany")
	}


	target_2, err := gormSchema.ParseWithSpecialTableName(
		&target.Order{},
		&sync.Map{},
		&gormSchema.NamingStrategy{},
		"",
	)

	if err == nil {
		createGormFile(target_2, "example/order.go", "Order")
	}

}

var basedir = "example"
var outputSuffix = "_gorm.go"
func createGormFile(schema *gormSchema.Schema, filename string, structName string) {
	gormFilePath := strings.Replace(filename, ".go", "", 1) + outputSuffix
	code := ""
	code += "func (t " + structName + ") TableName() string {\n"
	code += "\treturn \"" + schema.Table + "\"\n"
	code += "}\n"

	code += "func (t " + structName + ") StructName() string {\n"
	code += "\treturn \"" + structName + "\"\n"
	code += "}\n"

	columnConstantNames := []string{}
	for _, field := range schema.Fields {
		columnConstantName := structName + "_" + field.Name
		columnConstantExpression := "const " + columnConstantName + " = " + "\"" + field.DBName + "\"" + "\n"
		columnConstantNames = append(columnConstantNames, "\t\t"+columnConstantName+",")
		code += columnConstantExpression
	}

	code += "func (t " + structName + ") Columns() []string {\n"
	code += "\treturn []string{\n" + strings.Join(columnConstantNames, "\n") + "\n\t}\n"
	code += "}\n"

	f, err := os.OpenFile(gormFilePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	_, err = fmt.Fprintln(f, code)
}