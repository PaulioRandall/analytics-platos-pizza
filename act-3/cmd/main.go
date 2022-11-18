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
	todo("Put data into SQLite database").breakdown(
		todo("Insert data into in mmeory database").breakdown(
			todo("Insert orders CSV"),
			todo("Insert order_details CSV"),
			todo("Insert pizzas CSV"),
			todo("Insert pizza_types CSV"),
		),
		todo("Import SQLite library"),
		todo("Create SQLite database").breakdown(
			todo("Create tables"),
		),
		todo("Insert data into new SQLite database").breakdown(
			todo("Insert data dictionary"),
			todo("Insert orders"),
			todo("Insert order_details"),
			todo("Insert pizzas"),
			todo("Insert pizza_types"),
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
