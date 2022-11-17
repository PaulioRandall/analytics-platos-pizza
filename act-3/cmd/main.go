package main

import (
	"fmt"
	"os"
)

func main() {
	printArgs()
}

func printArgs() {
	for i, v := range os.Args {
		fmt.Print(i, ": ")
		fmt.Println(v)
	}
}
