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
	count := 1
	// Loop over each line of input
	for scanner.Scan() {
		// Read the value from current line
		input := scanner.Text()
		num, err := strconv.ParseFloat(input, 64)
		if err != nil {
			log.Fatal("Error parsing data")
		} else {
			yValues = guess_it.MaintainLastFive(yValues, num)
			xValues = guess_it.MaintainLastFive(xValues, float64(count))
		}
		count++

		if err := scanner.Err(); err != nil {
			fmt.Println("Error", err)
			return
		}
		if len(yValues) > 1 {
			a, b := guess_it.LinearRegression(xValues, yValues)
			stdDev := guess_it.StandardDeviation(yValues)
			lower, upper := guess_it.PredictiveInterval(float64(count), a, b, stdDev)
			fmt.Println(math.Round(lower), math.Round(upper))
		}
	}
}