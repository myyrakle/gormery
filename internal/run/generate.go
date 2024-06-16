package run

import (
	"github.com/myyrakle/gormery/internal/config"
	"github.com/myyrakle/gormery/pkg"
)

func RunGenerate() {
	configFile := config.Load()

	pkg.Generate(configFile)
}
