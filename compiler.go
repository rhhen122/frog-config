package main

import "fmt"

var usrinput string
var varcomf bool = false

func start() {
	fmt.Println("miku-go")
	fmt.Println("A Compiler built for turning .frog files into valid python configs")
	fmt.Println("===============")
	fmt.Print("")
}

func read() {
	for i := 0; i < len(usrinput); i++ {
		if varcomf == true {
			if string(usrinput[i]) == "*" {
				varcomf := false
			} else {
				fmt.Print(usrinput[i])
			}
		}
		if string(usrinput[i]) == "*" {
			varcomf := true
		}
	}
}

func main() {
	start()
	fmt.Scan(&usrinput)
	fmt.Println("===============")
	read()
}