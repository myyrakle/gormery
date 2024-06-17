package pkg

import (
	"sync"

	models "github.com/myyrakle/gormery/example"
	gormSchema "gorm.io/gorm/schema"
)

func GetGormSchemaFromValue() (*gormSchema.Schema, error) {
	schema, err := gormSchema.ParseWithSpecialTableName(
		&models.PersonSoMany{},
		&sync.Map{},
		&gormSchema.NamingStrategy{},
		"",
	)

	for _, field := range schema.Fields {
		if field.DBName == "id" {
			field.TagSettings["AUTO_INCREMENT"] = "true"
		}
	}

	return schema, err
}
