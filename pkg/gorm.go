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

	return schema, err
}
