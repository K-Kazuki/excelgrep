package cmd

import (
	"path/filepath"

	"github.com/K-Kazuki/excel_grep/logger"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(newFindCmd())
}

func newFindCmd() *cobra.Command {
	echoCmd := &cobra.Command{
		Use:   "find [Glob pattern]",
		Short: "Find xlsx files.",
		Long:  `Find xlsx files. Search recursively with "**/".`,
		Args:  cobra.MaximumNArgs(1),
		Run:   runCmd,
	}
	return echoCmd
}

func runCmd(cmd *cobra.Command, args []string) {
	if len(args) > 0 {
		path := filepath.Join(args[0], "*.xlsx")
		logger.Debugln(path)
	} else {
		logger.Debugln("non args")
		path := "./**/*.xlsx"
		logger.Debugln(path)
	}
}
