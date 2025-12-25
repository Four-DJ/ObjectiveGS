package main

import (
	"fmt"
	"gso/transpiler"
	"os"
)

func main() {
	args := os.Args[1:]
	text, err := os.ReadFile(args[0])
	if err != nil {
		fmt.Printf("(os.ReadFile) failed:\n%v", err)
		return
	}
	fmt.Println(string(text))
	fmt.Println(transpiler.Tokenize(string(text)))
}
