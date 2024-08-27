package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"guess-it-2/student/stats"
)

func main() {
	if len(os.Args) != 1 {
		fmt.Println("Usage: go run .")
	}

	var x []float64
	var y []float64
	scanner := bufio.NewScanner(os.Stdin)
	count := 0
	for scanner.Scan() {
		data := scanner.Text()
		yvalues, err := strconv.ParseFloat(data, 64)
		if err != nil {
			fmt.Println("error parsing data")
			return
		} else {
			x = append(x, float64(count))
			y = append(y, yvalues)
		}
		count++
		if err := scanner.Err(); err != nil {
			fmt.Println("Error reading input:", err)
			return
		}
		yEst := guess_it.Y_estimate(x, y)
		lowerL, upperL := guess_it.PredictInt(x, y, yEst)
		fmt.Println(lowerL, upperL)
		
	}
}
