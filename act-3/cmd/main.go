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
		println("ERROR:\n ", e.Error())
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
		todo("Add the tables & data:").breakdown(
			todo("Design, create, insert, & read back orders table & data"),
			todo("Design, create, insert, & read back order_details table & data"),
			todo("Design, create, insert, & read back pizzas table & data"),
			todo("Design, create, insert, & read back pizza_types table & data"),
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
