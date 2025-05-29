/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	log "github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var Path string

// buildCmd represents the build command
var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Build the application",
	// A longer description that spans multiple lines and likely contains examples
	Long: `Build the application using the specified configuration.
This command compiles the source code and prepares the application for deployment.`,
	Run: func(cmd *cobra.Command, args []string) {
		for key, value := range viper.GetViper().AllSettings() {
			log.WithFields(log.Fields{
				key: value,
			}).Info("Command Flag")
		}
	},
}

func init() {
	rootCmd.AddCommand(buildCmd)
	buildCmd.PersistentFlags().StringVarP(&Path, "path", "p", ".", "Define the path to scan.")
	// rootCmd.MarkPersistentFlagRequired("path")
	viper.BindPFlag("path", buildCmd.PersistentFlags().Lookup("path"))

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// buildCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// buildCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
