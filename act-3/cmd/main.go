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
		todo("Design initial database interface"),
		todo("Create in memory implementation"),
		todo("Create data dictionary table").breakdown(
			todo("Create data model"),
			todo("Insert data from CSV"),
		),
		todo("Create order table").breakdown(
			todo("Create data model"),
			todo("Insert data from CSV"),
		),
		todo("Create order_details table").breakdown(
			todo("Create data model"),
			todo("Insert data from CSV"),
		),
		todo("Create pizzas table").breakdown(
			todo("Create data model"),
			todo("Insert data from CSV"),
		),
		todo("Create pizza_types table").breakdown(
			todo("Create data model"),
			todo("Insert data from CSV"),
		),
		todo("Import SQLite library"),
		todo("Create SQLite database").breakdown(
			todo("Create tables"),
		),
		todo("Insert data into new SQLite database").breakdown(
			todo("Insert data dictionary").breakdown(
				todo("Insert data"),
				todo("Read data"),
			),
			todo("Insert orders").breakdown(
				todo("Insert data"),
				todo("Read data"),
			),
			todo("Insert order_details").breakdown(
				todo("Insert data"),
				todo("Read data"),
			),
			todo("Insert pizzas").breakdown(
				todo("Insert data"),
				todo("Read data"),
			),
			todo("Insert pizza_types").breakdown(
				todo("Insert data"),
				todo("Read data"),
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
