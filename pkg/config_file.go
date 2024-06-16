package pkg

type ConfigFile struct {
	Basedir      string `yaml:"basedir"`       // 기본 경로 (ex: domain)
	BaseImport   string `yaml:"base-import"`   // 기본 import (ex: github.com/myyrakle/gormery/example)
	OutputSuffix string `yaml:"output-suffix"` // 출력 파일 접미사 (ex: _gorm.go)
	RunnerPath   string `yaml:"runner-path"`   // 실행시킬 파일 경로
}
