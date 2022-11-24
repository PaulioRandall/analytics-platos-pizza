package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/PaulioRandall/analytics-platos-pizza/act-3/pkg/workflow"
)

func main() {
	fmt.Println()

	if e := workflow.Execute(); e != nil {
		println("ERROR:", e.Error())
		os.Exit(1)
	}

	//printArgs()
	fmt.Println()
	printTasks(todos, 0)
}

func printArgs() {
	for i, v := range os.Args {
		fmt.Print(i, ": ")
		fmt.Println(v)
	}
}

var todos = []task{
	todo("Create SQLite implementation of the database interface").breakdown(
		todo("Find and import SQLite Go package & any drivers").breakdown(
			todo("Create database as a file"),
			todo("Add the tables & data:").breakdown(
				todo("Design & create tables based upon content models"),
				todo("Write SQL to insert data from content models into tables"),
				todo("Write SQL to read data to ensure correct storage"),
				todo("Insert data & read back to check success"),
			),
		),
	),
}

type task struct {
	desc string
	subs []task
}

func todo(desc string) task {
	return task{
		desc: desc,
	}
}

func (t task) breakdown(subs ...task) task {
	t.subs = subs
	return t
}

func printTasks(tasks []task, indent int) {
	prefix := strings.Repeat("\t", indent)

	for i, t := range tasks {
		fmt.Printf("%s%d: %s\n", prefix, i+1, t.desc)
		printTasks(t.subs, indent+1)
	}
}
