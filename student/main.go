package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"

	"guess-it-2/stats"
)

func main() {
	if len(os.Args) != 1 {
		fmt.Println("Usage: go run .")
	}
	scanner := bufio.NewScanner(os.Stdin)
	xValues := []float64{}
	yValues := []float64{}
	count := 0
	// Loop over each line of input
	for scanner.Scan() {
		// Read the value from current line
		input := scanner.Text()
		num, err := strconv.ParseFloat(input, 64)
		if err != nil {
			log.Fatal("Error parsing data")
		} else {
			yValues = append(yValues, num)
			xValues = append(xValues, float64(count))
		}
		count++

		if err := scanner.Err(); err != nil {
			fmt.Println("Error", err)
			return
		}
		slope, intercept, meanX, ssx, stdError := guess_it.LinearRegression(xValues, yValues)
		lower, upper := guess_it.Predictiveinterval(xValues, slope, intercept, meanX, ssx, stdError)
		fmt.Println(math.Round(lower), math.Round(upper))
	}
}
