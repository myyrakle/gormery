package run

import (
	"fmt"

	steps "github.com/myyrakle/gormery/internal/steps"
)

func RunGenerate() {
	fmt.Println(">>> Running LoadConfigFile")
	configFile := steps.LoadConfigFile()

	fmt.Println(">>> Running ReadAllTargets")
	targets := steps.ReadAllTargets(configFile)

	fmt.Println(">>> Running GenerateRunner")
	steps.GenerateRunner(configFile, targets)

	fmt.Println(">>> Running RunRunner")
	steps.RunRunner(configFile)
}
