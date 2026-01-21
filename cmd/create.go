/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var projectType string

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create [project-name]",
	Short: "Create a new project scaffold",
	Long:  "Create generates a new Go project with the required folders and starter files.",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println(" Error: project name required")
			return
		}

		projectName := args[0]
		fmt.Println(" Creating project:", projectName)
		fmt.Println(" Type:", projectType)

		err := CreateProject(projectName, projectType)
		if err != nil {
			fmt.Print(err)
		}

	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	createCmd.Flags().StringVarP(&projectType, "type", "t", "default", "type of  application you want to build")

}
