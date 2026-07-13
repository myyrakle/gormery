package run

import (
	"fmt"

	steps "github.com/myyrakle/gormery/internal/steps"
)

// RunGenerate generates gormery code. When selected is non-empty, only the
// given source files are processed; otherwise the whole basedir is scanned.
func RunGenerate(selected []string) {
	fmt.Println(">>> Running LoadConfigFile")
	configFile := steps.LoadConfigFile()

	fmt.Println(">>> Running ReadAllTargets")
	targets := steps.ReadAllTargets(configFile, selected)

	fmt.Println(">>> Running GenerateRunner")
	steps.GenerateRunner(configFile, targets)

	fmt.Println(">>> Running RunRunner")
	steps.RunRunner(configFile)
}
