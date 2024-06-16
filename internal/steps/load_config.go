package steps

import (
	"fmt"
	"log"
	"os"

	config "github.com/myyrakle/gormery/internal/config"
	yaml "gopkg.in/yaml.v2"
)

func LoadConfigFile() config.ConfigFile {
	bytes, err := os.ReadFile(".gormery.yaml")

	if err != nil {
		fmt.Println("Error: .gormery.yaml file not found.")
	}

	decoded := &config.ConfigFile{}
	err = yaml.Unmarshal(bytes, decoded)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	fmt.Println(">>> Config file loaded")

	return *decoded
}
