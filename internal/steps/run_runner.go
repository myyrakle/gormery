package steps

import (
	"os/exec"

	config "github.com/myyrakle/gormery/internal/config"
)

func RunRunner(configFile config.ConfigFile) {
	filePath := configFile.RunnerPath + "/main.go"

	// go run으로 실행
	// go run runnerPath/main.go
	cmd := exec.Command("go", "run", filePath)
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}
