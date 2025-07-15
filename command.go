package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CmdFlags struct {
	Add    string
	Del    int
	Edit   string
	Toggle int
	List   bool
}

func NewCmdFlags() *CmdFlags {
	cf := CmdFlags{}
	flag.StringVar(&cf.Add, "a", "", "Add a new todo")
	flag.StringVar(&cf.Edit, "e", "", "Edit a todo by index and specify a new title. Format: index:new_title")
	flag.IntVar(&cf.Del, "d", -1, "Delete a todo by index")
	flag.IntVar(&cf.Toggle, "t", -1, "Toggle a todo by index")
	flag.BoolVar(&cf.List, "l", false, "List all todos")

	flag.Usage = func() {
		fmt.Println("Todo CLI Usage:")
		flag.PrintDefaults()
	}

	flag.Parse()

	// If no arguments are provided (only the binary is run)
	if len(os.Args) == 1 {
		flag.Usage()
		os.Exit(0)
	}

	return &cf
}


func (cf *CmdFlags) Execute(todos *Todos) {
	switch {
	case cf.List:
		todos.print()
	case cf.Add != "":
		todos.add(cf.Add)
	case cf.Edit != "":
		parts := strings.SplitN(cf.Edit, ":", 2)
		if len(parts) != 2 {
			fmt.Println("Error: Invalid format for edit. Please use index:new_title")
			os.Exit(1)
		}
		index, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Println("Error: Invalid index for edit.")
			os.Exit(1)

		}
		todos.edit(index, parts[1])
	case cf.Toggle != -1:
		todos.toggle(cf.Toggle)

	case cf.Del != -1:
		todos.delete(cf.Del)

	default:
		fmt.Println("Invalid command")
	}
}
