/*
Copyright © 2025 sub zero <EMAIL ADDRESS>
*/
package cmd

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var Verbose bool
var Debug bool
var config string //
var daemon bool   //
var version bool  //

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "pici",
	Short: "pipeline CLI tool",
	Long: `Pici is a CLI tool for managing and executing pipelines no matter where they are running.
It provides a simple interface to interact with pipelines, allowing users to create, manage, and execute them seamlessly across different environments.`,
	// parse the config if one is provided, or use the defaults. Set the backend
	// driver to be used
	PersistentPreRun: func(ccmd *cobra.Command, args []string) {

		// if --config is passed, attempt to parse the config file
		if config != "" {

			// get the filepath
			abs, err := filepath.Abs(config)
			if err != nil {
				logrus.Error("Error reading filepath: ", err.Error())
			}

			// get the config name
			base := filepath.Base(abs)

			// get the path
			path := filepath.Dir(abs)

			//
			viper.SetConfigName(strings.Split(base, ".")[0])
			viper.AddConfigPath(path)

			// Find and read the config file; Handle errors reading the config file
			if err := viper.ReadInConfig(); err != nil {
				logrus.Fatal("Failed to read config file: ", err.Error())
				os.Exit(1)
			}
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.pici.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "Display more verbose output in console output. (default: false)")
	viper.BindPFlag("verbose", rootCmd.PersistentFlags().Lookup("verbose"))

	rootCmd.PersistentFlags().BoolVarP(&Debug, "debug", "d", false, "Display debugging output in the console. (default: false)")
	viper.BindPFlag("debug", rootCmd.PersistentFlags().Lookup("debug"))

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
