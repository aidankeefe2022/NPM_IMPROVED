/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "NPM_IMPROVED",
	Short: "A tool to allow npm audit to be run down a project file path and allow for targeting of projects and specified output.",
	Long: `A tool to allow npm audit to be run down a project file path and allow for targeting of projects and specified output.

NPM_IMPROVED rAudit --project "PATH" --output "PATH"`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err.Error(), "Error in root")
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.NPM_IMPROVED.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
}
