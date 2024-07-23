package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	calc "guess-it-1/calculation"
)

func main() {
	if len(os.Args) != 1 {
		fmt.Println("Usage: go run main.go")
		return
	}

	var data []float64
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		num, err := strconv.ParseFloat(input, 64)
		if err != nil {
			fmt.Println("Invalid input. Please enter a valid number.")
			continue
		}

		data = append(data, num)
		if len(data) > 1 {
			lower, upper := calc.Range(data)
			fmt.Printf("%d %d\n", lower, upper)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input:", err)
	}
}
