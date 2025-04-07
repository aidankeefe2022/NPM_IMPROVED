/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"NPM_IMPROVED/unifyPackageLock"
	"fmt"
	"github.com/spf13/cobra"
)

// rAuditCmd represents the rAudit command
var rAuditCmd = &cobra.Command{
	Use:   "rAudit",
	Short: "Searches down file tree for dependencies of the project and audits them",
	Long:  `The rAudit command searches down file tree for dependencies of the project and audits them.`,
	Run: func(cmd *cobra.Command, args []string) {
		projectPath, _ := cmd.Flags().GetString("project")
		outputPath, _ := cmd.Flags().GetString("out")
		packageLockRequired, _ := cmd.Flags().GetBool("rpackagelock")
		var targetFileName string
		if packageLockRequired {
			targetFileName = "package-lock.json"
		} else {
			targetFileName = "package.json"
		}
		if dirs, err := unifyPackageLock.GetTargetDirs(projectPath, targetFileName); err != nil {
			fmt.Println(err, "error while getting target files")
			return
		} else if err := unifyPackageLock.CombinedOutputNpmAudit(dirs, outputPath, packageLockRequired); err != nil {
			fmt.Println(err, "error while combining npm outputs")
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(rAuditCmd)

	rAuditCmd.Flags().String("project", "", "The target project root.")
	rAuditCmd.Flags().String("out", "", "The location for the json report to be outputted.")
	rAuditCmd.Flags().Bool("rpackagelock", true, "When true npm audit will only be run on package lock file found in the target project. When false a package lock file will be create and then npm audit will be run.")

	rAuditCmd.MarkFlagRequired("project")
	rAuditCmd.MarkFlagRequired("out")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// rAuditCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// rAuditCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
