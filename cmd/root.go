package cmd

import (
	"os"

	"github.com/myyrakle/gormery/internal/run"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gormery [files...]",
	Short: "generate gormery codes",
	Long: "generate gormery codes\n\n" +
		"With no arguments, every file under basedir is processed.\n" +
		"Pass one or more file paths to generate only for those files, e.g.:\n" +
		"  gormery example/clothes.go example/order.go",
	// Accept positional file paths; without this cobra's legacyArgs rejects
	// them as unknown subcommands because a "version" subcommand is registered.
	Args: cobra.ArbitraryArgs,
	Run: func(cmd *cobra.Command, args []string) {
		run.RunGenerate(args)
	},
}

func Execute() {
	err := rootCmd.Execute()

	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
