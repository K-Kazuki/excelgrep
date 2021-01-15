/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/K-Kazuki/excel_grep/logger"
	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string
var useVerboseLogger bool

// rootCmd represents the base command when called without any subcommands
func NewCmdRoot() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "eg",
		Short: "excel_grep (eg) recursively searches your current directory and grep the xlsx files.",
		Long: `
USAGE:
		eg [OPTIONS] PATTERN [PATH ...]
		command | eg [OPTIONS] PATTERN

ARGS:
    <PATTERN>    A glob pattern used for  searching.
		<PATH>...    A file or directory to search.`,
		Args: cobra.MinimumNArgs(1),
		Run:  runRootCmd,
	}

	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.excel_grep.yaml)")
	rootCmd.PersistentFlags().BoolVar(&useVerboseLogger, "verbose", false, "Verbose log enable flag")

	return rootCmd
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	rootCmd := NewCmdRoot()
	rootCmd.SetOutput(os.Stdout)

	if err := rootCmd.Execute(); err != nil {
		rootCmd.SetOutput(os.Stderr)
		rootCmd.Println(err)
		os.Exit(1)
	}

	logger.Debugln("DONE")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".excel_grep" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".excel_grep")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}

	// ロガーを設定
	if useVerboseLogger {
		logger.SetLogger(logger.Verbose)
	}
}
