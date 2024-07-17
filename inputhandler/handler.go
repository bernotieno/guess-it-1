package handler

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	calc "guess-it-1/calculation"
)

func HandleInput() {
	var data []float64
	reader := bufio.NewReader(os.Stdin)
	for {
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "" {
			fmt.Println("Exiting...")
			break
		}

		num, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input. Please enter a valid number.")
			continue
		}

		data = append(data, float64(num))
		if len(data) > 1 {
			lower, upper := calc.Range(data)
			fmt.Printf("%d %d\n", lower, upper)
		}
	}
}
