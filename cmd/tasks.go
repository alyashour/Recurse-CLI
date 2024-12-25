/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"time"

	"recurse-core/task"

	"github.com/spf13/cobra"
)

type TaskManager interface {
	CreateTask(title, description, assignee, string, dueDate time.Time) (task.Task, error)
	GetTask(id int) (task.Task, error)
	UpdateTask(id int, updates task.Task) error
	DeleteTask(id int) error
	ListTasks() ([]task.Task, error)
}

// tasksCmd represents the tasks command
var tasksCmd = &cobra.Command{
	Use:   "tasks",
	Short: "Manage a list of tasks",
	RunE: func(cmd *cobra.Command, args []string) error {
		return fmt.Errorf("no command provided")
	},
}

var taskNewCmd = &cobra.Command{
	Use:     "new",
	Aliases: []string{"create"},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Creating task")
	},
}

func init() {
	rootCmd.AddCommand(tasksCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// tasksCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// tasksCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
