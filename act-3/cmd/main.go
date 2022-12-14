package main

import (
	"fmt"
	"os"

	"github.com/PaulioRandall/go-trackerr"

	"github.com/PaulioRandall/analytics-platos-pizza/act-3/scene-2/workflow"
)

func main() {
	fmt.Println()

	var e error
	defer trackerr.DebugPanic(&e)

	if e = workflow.Execute(); e != nil {
		trackerr.Debug(e)
		os.Exit(1)
	}

	//printArgs()
}

func printArgs() {
	for i, v := range os.Args {
		fmt.Print(i, ": ")
		fmt.Println(v)
	}
}
