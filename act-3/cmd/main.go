package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println()
	//printArgs()
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
			todo("Insert data dictionary CSV"),
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
