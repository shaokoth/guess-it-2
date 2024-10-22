package guess_it

import "math"


// Function to calculate the mean of a slice
func mean(values []float64) float64 {
	sum := 0.0
	for _, v := range values {
		sum += v
	}
	return sum / float64(len(values))
}

// Function to calculate the standard deviation of the y-values
func StandardDeviation(values []float64) float64 {
	meanValue := mean(values)
	var sumSquares float64
	for _, v := range values {
		sumSquares += (v - meanValue) * (v - meanValue)
	}
	variance := sumSquares / float64(len(values))
	return math.Sqrt(variance)
}


// LinearRegression calculates the slope, intercept, mean of x, and SSX.
func LinearRegression(x, y []float64) (float64, float64) {
xMean := mean(x)
yMean := mean(y)

// calculate slope (b)
var num, denom float64
for i := 0; i < len(x); i++ {
	num += (x[i] - xMean) * (y[i] - yMean)
		denom += (x[i] - xMean) * (x[i] - xMean)
}
b := num/denom

// calculate intercept (a)
a := yMean - (b * xMean)

return a, b
}

// Function to calculate the range for the next number
func PredictiveInterval(x, a, b, stdDev float64) (float64, float64) {
	ypred := a + b * x

	// Set a buffer for the range (can be based on standard deviation or a fixed value)
	lowerbound := ypred - math.Max(1, ypred - 3 * stdDev)
	upperbound := ypred + 3 * stdDev

	return lowerbound, upperbound

}
// Function to manage the sliding window of the last 5 values
func MaintainLastFive(slice []float64, newValue float64) []float64 {
	if len(slice) < 5.0 {
		return append(slice, newValue)
	}
	return append(slice[1:], newValue)
}
