/*
Copyright © 2024 ALY ASHOUR <ALYASHOUR1@GMAIL.COM>
*/

package cmd

import (
	"fmt"
	"time"

	"recurse-core/task"
	core "recurse-core/task"

	"github.com/spf13/cobra"
)

type TaskManager interface {
	CreateTask(title, description, assignee string, dueDate time.Time) (*task.Task, error)
	GetTask(id int) (*task.Task, error)
	UpdateTask(id int, updates task.Task) error
	DeleteTask(id int) error
	ListTasks() ([]task.Task, error)
}

var manager = core.JSONTaskManager{}

var TasksCmd = &cobra.Command{
	Use:   "tasks",
	Short: "Manage a list of tasks",
	RunE: func(cmd *cobra.Command, args []string) error {
		return fmt.Errorf("no command provided")
	},
}

var newTaskCmd = &cobra.Command{
	Use:     "new",
	Aliases: []string{"create"},
	Run: func(cmd *cobra.Command, args []string) {
		if task, err := manager.CreateTask("title", "description", "assignee", time.Now()); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(task)
		}
	},
}

var listTasksCmd = &cobra.Command{
	Use:   "list",
	Short: "lists all tasks in the currently selected org",
	Run: func(cmd *cobra.Command, args []string) {
		if list, err := manager.ListTasks(); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(list)
		}
	},
}

func init() {
	TasksCmd.AddCommand(newTaskCmd)
	TasksCmd.AddCommand(listTasksCmd)
}
