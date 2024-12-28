package main

import (
	"fmt"
	"os"
	"strings"
	"github.com/V1nic1us/task-trace-go/function"
)

func main() {
	if len(os.Args) < 2 && os.Args[1] != "list" {
		fmt.Println("expected 'add' or 'update' or 'delete' or 'mark-done' or 'mark-in-progress' or 'list' subcommands")
		os.Exit(1)
	}

	dat :=function.VerifyFile()

	command := strings.ToLower(os.Args[1])

	switch command {
	case "add":
		function.AddTask(dat)
	case "update":
		function.UpdateTask(dat)
	case "delete":
		function.DeleteTask(dat)
	case "mark-done":
		function.MarkDoneTask(dat)
	case "mark-in-progress":
		function.MarkInProgressTask(dat)
	case "list":
		function.ListTasks(dat)
	default:
		fmt.Println("expected 'add' or 'update' or 'delete' or 'mark-done' or 'mark-in-progress' or 'list' subcommands")
		os.Exit(1)
	}
}


