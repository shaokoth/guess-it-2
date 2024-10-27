package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	guess_it "guess-it-2/stats"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	yValues := []float64{}
	// Loop over each line of input
	for scanner.Scan() {
		// Read the value from current line
		input := scanner.Text()
		num, err := strconv.ParseFloat(input, 64)
		if err != nil {
			log.Fatal("Error parsing data")
		} else {
			yValues = append(yValues, num)
		}
		// Check if there are at least 4 values
		if len(yValues) > 1 {
			lowerLImit, upperLImit := guess_it.Range(yValues)
			fmt.Printf("%d %d\n", lowerLImit, upperLImit)
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error", err)
		return
	}
}
