package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/CMD-Tools/interacting/todo"
)

// hardcoding the file name
const todoFilename = ".todo.json"

func main() {
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
	case len(os.Args) == 1:
		//list current todo items
		for _, item := range *l {
			fmt.Println(item.Task)
		}

		//Concatenate all provided arguments with a space and
		//add to the list an an item
	default:
		// concatenate all args with a space
		item := strings.Join(os.Args[1:], " ")

		//Add the task
		l.Add(item)

		// Save the new list
		if err := l.Save(todoFilename); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}
}
