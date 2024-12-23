package main

import "fmt"

// import (
// 	"fmt"
// 	"os"
// 	"strings"
// )

var tasks []string

func addTask(task string) {
	tasks = append(tasks, task)
	fmt.Println("Added:", task)
}

func listTasks() {
	fmt.Println("Your tasks:")
	for i, task := range tasks {
		fmt.Printf("%d. %s\n", i+1, task)
	}
}

// func main() {
// 	if len(os.Args) < 2 {
// 		fmt.Println("Please provide a command: add/list")
// 		return
// 	}

// 	command := os.Args[1]

// 	switch command {
// 	case "add":
// 		if len(os.Args) < 3 {
// 			fmt.Println("Please provide a task to add.")
// 		} else {
// 			task := strings.Join(os.Args[2:], " ")
// 			addTask(task)
// 		}
// 	case "list":
// 		listTasks()
// 	default:
// 		fmt.Println("Unknown command:", command)
// 	}
// }
