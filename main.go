package main

import (
	"fmt"
	"sync"

	models "github.com/myyrakle/gormery/example"
	gorm "gorm.io/gorm/schema"
)

func main() {
	schema, err := gorm.ParseWithSpecialTableName(
		&models.PersonSoMany{},
		&sync.Map{},
		&gorm.NamingStrategy{},
		"",
	)
	if err != nil {
		panic(err)
	}

	fmt.Println(schema.Table)

	for _, field := range schema.Fields {
		fmt.Println(field)
	}
}
