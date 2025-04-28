package main

import "fmt"

var usrinput string

func start() {
	fmt.Println("miku-go")
	fmt.Println("A Compiler built for turning .frog files into valid python configs")
	fmt.Println("===============")
	fmt.Print("")
}

func read() {
	for i := 0; i < len(usrinput); i++ {
		if usrinput[i] == `"` {

		}
	}
}

func main() {
	start()
	fmt.Scan(&usrinput)
	fmt.Println("===============")
}