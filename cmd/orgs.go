/*
Copyright © 2025 ALY ASHOUR <ALYASHOUR1@GMAIL.COM>
*/

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// orgsCmd represents the orgs command
var orgsCmd = &cobra.Command{
	Use:   "orgs",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("orgs called")
	},
}

func init() {
	rootCmd.AddCommand(orgsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// orgsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// orgsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
