/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// ruleCmd represents the rule command
var ruleCmd = &cobra.Command{
	Use:   "rule",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		for _, v := range args {
			fmt.Println(v)
		}
		//没有输入name时
		if len(name) == 0 {
			name = "default"
		}
		fmt.Println(name)
	},
}

var name string

func init() {
	createCmd.AddCommand(ruleCmd)

	ruleCmd.Flags().StringVarP(&name, "name", "n", "", "rule name")
}
