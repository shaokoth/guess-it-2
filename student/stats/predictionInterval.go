package guess_it

import "math"

// Function to calculate the mean of a slice
func Mean(values []float64) float64 {
	sum := 0.0
	for _, v := range values {
		sum += v
	}
	return sum / float64(len(values))
}

// Function to calculate the standard deviation of the y-values
func StandardDev(input []float64, mean float64) float64 {
	sum := 0.0
	for _, val := range input {
		sum += (val - mean) * (val - mean)
	}
	return math.Sqrt(sum / float64(len(input)))
}

// LinearRegression calculates the slope, intercept, mean of x, and SSX.
func LinearRegression(y []float64) (m, c float64) {
	// Generate x axis values as indices of y
	x := make([]float64, len(y))
	for i := range y {
		x[i] = float64(i)
	}

	// Get means of x and y
	meanX := Mean(x)
	meanY := Mean(y)

	// Calculate variances and standard deviations for x and y
	stddevX := StandardDev(x, meanX)
	stddevY := StandardDev(y, meanY)

	// Calculate Pearson’s correlation coefficient
	corelation := PearsonsCorrelationCoefficient(y)

	// Calculate slope (m)
	m = corelation * stddevY / stddevX

	// Calculate intercept (c) where the regression line crosses the y-axis
	c = meanY - m*meanX
	return m, c
}

// Calculates the Pearson Correlation Coefficient for y with indices as x
func PearsonsCorrelationCoefficient(y []float64) float64 {
	// Generate x values as indices of y
	x := make([]float64, len(y))
	for i := range y {
		x[i] = float64(i)
	}

	// Calculate means of x and y
	meanX := Mean(x)
	meanY := Mean(y)

	// Calculate numerator and the sums of squares
	var sumXY, sumXSquare, sumYSquare float64
	for i, yi := range y {
		xi := float64(i)
		diffX := xi - meanX
		diffY := yi - meanY
		sumXY += diffX * diffY
		sumXSquare += diffX * diffX
		sumYSquare += diffY * diffY
	}

	// Calculate the denominator
	denominator := math.Sqrt(sumXSquare * sumYSquare)
	if denominator == 0 {
		return 0 // Avoid division by zero
	}

	// Calculate and return correlation coefficient
	return sumXY / denominator
}

// Function to calculate the range for the next number
func PredictNextValue(y []float64) float64 {
	slope, intercept := LinearRegression(y)
	nextIndex := float64(len(y))
	// y=mx+c
	return slope*nextIndex + intercept
}
