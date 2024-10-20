package main

import (
	"fmt"
)

func main() {
	for {
		var input string
		fmt.Printf("pokdex > ")
		fmt.Scan(&input)
		fmt.Printf("Input: %v \n", input)
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}
