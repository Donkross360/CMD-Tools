package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/CMD-Tools/interacting/todo"
)

// Default the file name
var todoFileName = ".todo.json"

func main() {

	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "%s tool. Developed for MidasTech\n", os.Args[0])
		fmt.Fprintf(flag.CommandLine.Output(), "Copyright 2025\n")
		fmt.Fprintf(flag.CommandLine.Output(), "Usage information:\n")
		flag.PrintDefaults()
	}

	//Parse command line flags
	add := flag.Bool("add", false, "Add task to the todo list: With -add you can add task\n to the list via STDIN or args after the -add option\n input the task directly or use a file with tasks in it\n Ex. ./todo -add args")
	list := flag.Bool("list", false, "list all task - Ex. to list all task use ./todo -list")
	complete := flag.Int("complete", 0, "Item to be completed: use the -complete option\n with the task id to mark a task as Done.\n Ex. to mark the nth number use: ./todo -complete n ")
	delete := flag.Int("del", 0, "Delete an item form list: use the -del option\n with the task id to delete an item from list.\n Ex. to delete the nth item use: ./todo -del n")
	verbose := flag.Bool("vlist", false, "List all task with date and time\n - to list all task with created date and time use: ./todo -vlist")
	compNoShow := flag.Bool("xcomp", false, "Remove completed task from listed items \n- Ex. with ./todo -xcomp you can remove all completed item from list")

	flag.Parse()

	//Check if the user defined the ENV VAR for a custom file name
	if os.Getenv("TODO_FILENAME") != "" {
		todoFileName = os.Getenv("TODO_FILENAME")
	}

	// define an items list
	l := &todo.List{}

	// Use the Get method to read to do items from file
	if err := l.Get(todoFileName); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// Decide what to do based on the number of arguments provided
	switch {
	// for no extra arguments, print the list
	case *list:
		//list current todo items
		fmt.Print(l)

	case *complete > 0:
		// Complete the given item
		if err := l.Compelete(*complete); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(0)
		}
		//Save the new list
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	case *delete > 0:
		if err := l.Delete(*delete); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	case *add:
		// When any argument (excluding flags) are provided, they will be
		// used as the new task
		t, err := getTask(os.Stdin, flag.Args()...)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		l.Add(t)

		//save the new list
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	case *verbose:

		for i, item := range *l {
			formattedTime := item.CreatedAt.Format("2006-01-02 15:04:05")

			fmt.Printf("%d %s %q finished:%v\n", i+1, item.Task, formattedTime, item.Done)
		}
	case *compNoShow:
		for i, item := range *l {
			if !item.Done {
				fmt.Printf("%d %s\n", i+1, item.Task)
			}
		}

	default:
		// Invalid flag provided
		fmt.Fprintln(os.Stderr, "Invalid option")
		os.Exit(1)
	}
}

func getTask(r io.Reader, args ...string) (string, error) {
	if len(args) > 0 {
		return strings.Join(args, " "), nil
	}
	s := bufio.NewScanner(r)
	s.Scan()
	if err := s.Err(); err != nil {
		return "", err
	}
	if len(s.Text()) == 0 {
		return "", fmt.Errorf("task cannot be blank")
	}
	return s.Text(), nil
}
