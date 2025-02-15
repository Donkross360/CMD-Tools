package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/CMD-Tools/interacting/todo"
)

// hardcoding the file name
const todoFilename = ".todo.json"

func main() {

	//Parse command line flags
	task := flag.String("task", "", "Task to be included in the ToDo list")
	list := flag.Bool("list", false, "list all task")
	complete := flag.Int("complete", 0, "Item to be completed")

	flag.Parse()

	// define an items list
	l := &todo.List{}

	// Use the Get method to read to do items from file
	if err := l.Get(todoFilename); err != nil {
		fmt.Fscanln(os.Stderr, err)
		os.Exit(1)
	}

	// Decide what to do based on the number of arguments provided
	switch {
	// for no extra arguments, print the list
	case *list:
		//list current todo items
		for _, item := range *l {
			if !item.Done {
				fmt.Println(item.Task)
			}
		}

	case *complete > 0:
		// Complete the given item
		if err := l.Compelete(*complete); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(0)
		}
		//Save the new list
		if err := l.Save(todoFilename); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	case *task != "":
		//Add the task
		l.Add(*task)

		//save the new list
		if err := l.Save(todoFilename); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	default:
		// Invalid flag provided
		fmt.Fprintln(os.Stderr, "Invalid option")
		os.Exit(1)
	}
}
