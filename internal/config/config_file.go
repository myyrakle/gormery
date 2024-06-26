package config

type Feature string

type Features []Feature

func (f Features) Contains(feature Feature) bool {
	for _, v := range f {
		if v == feature {
			return true
		}
	}
	return false
}

const (
	FeatureSlice Feature = "SLICE"
)

type ConfigFile struct {
	Basedir      string   `yaml:"basedir"`       // 기본 경로 (ex: domain)
	OutputSuffix string   `yaml:"output-suffix"` // 출력 파일 접미사 (ex: _gorm.go)
	RunnerPath   string   `yaml:"runner-path"`   // 실행시킬 파일 경로
	Features     Features `yaml:"features"`      // 사용할 기능 목록

	// internal config
	ModuleName string // 프로젝트 모듈명 (go.mod에서 파싱)
}
