package main

import (
	"fmt"
	"os"

	handler "guess-it-1/inputhandler"
)

func main() {
	if len(os.Args) != 1 {
		fmt.Println("Usage: go run main.go ")
		return
	}
	handler.HandleInput()
}
