package main

import (
	 "sync"

	 "os"

	 "fmt"

	target "github.com/myyrakle/gormery/example"
	gormSchema "gorm.io/gorm/schema"
)

func main() {

	target_0, err := gormSchema.ParseWithSpecialTableName(
		&target.Order{},
		&sync.Map{},
		&gormSchema.NamingStrategy{},
		"",
	)

	if err == nil {
		createGormFile(target_0, "example/order.go", "models", "Order")
	}


	target_1, err := gormSchema.ParseWithSpecialTableName(
		&target.Person{},
		&sync.Map{},
		&gormSchema.NamingStrategy{},
		"",
	)

	if err == nil {
		createGormFile(target_1, "example/person.go", "models", "Person")
	}


	target_2, err := gormSchema.ParseWithSpecialTableName(
		&target.PersonSoMany{},
		&sync.Map{},
		&gormSchema.NamingStrategy{},
		"",
	)

	if err == nil {
		createGormFile(target_2, "example/person.go", "models", "PersonSoMany")
	}

}

var basedir = "example"
var outputSuffix = "_gorm.go"
func createGormFile(schema *gormSchema.Schema, filename string, packageName string, structName string) {
	gormFilePath := basedir + "/" + filename + outputSuffix
	code := ""
	code += "func (t " + structName + ") TableName() string {\n"
	code += "\treturn \"" + schema.Table + "\"\n"
	code += "}\n\n"
	f, err := os.OpenFile(gormFilePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	_, err = fmt.Fprintln(f, code)
}