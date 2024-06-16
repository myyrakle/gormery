package steps

import (
	"fmt"
	"log"
	"os"
	"strings"

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

	// go.mod 파싱해서 모듈명 가져오기
	moduleName := GetModuleNameFromGoMod()
	decoded.ModuleName = moduleName

	fmt.Println(">>> Config file loaded")

	return *decoded
}

func GetModuleNameFromGoMod() string {
	modBytes, err := os.ReadFile("go.mod")
	if err != nil {
		log.Fatalf("go.mod file not found")
	}

	modLines := string(modBytes)
	modLines = modLines[:len(modLines)-1]

	modLines = modLines[:len(modLines)-1]
	modLines = modLines[:len(modLines)-1]

	modLines = strings.Replace(modLines, "\n", " ", -1)
	modLines = strings.Replace(modLines, "\t", " ", -1)
	modLines = strings.Replace(modLines, "\r", " ", -1)

	modLines = strings.TrimSpace(modLines)

	modParts := strings.Split(modLines, " ")

	return modParts[1]
}
