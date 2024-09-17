package guess_it

import "math"

// LinearRegression calculates the slope, intercept, mean of x, and SSX.
func LinearRegression(x, y []float64) (float64, float64, float64, float64, float64) {
	N := float64(len(x))
	sumX, sumY, sumXY, sumX2 := 0.0, 0.0, 0.0, 0.0

	for i := 0; i < len(x); i++ {
		sumX += x[i]
		sumY += y[i]
		sumXY += x[i] * y[i]
		sumX2 += x[i] * x[i]
	}

	meanX := sumX / N
	meanY := sumY / N

	slope := (N*sumXY - sumX*sumY) / (N*sumX2 - sumX*sumX)
	intercept := meanY - slope*meanX

	// Calculate SSX (sum of squares of x deviations)
	ssx := 0.0
	for i := 0; i < len(x); i++ {
		ssx += (x[i] - meanX) * (x[i] - meanX)
	}
	// Calculates the residual sum of squares and standard error
	residualSum := 0.0
	for i := 0; i < len(x); i++ {
		yPred := intercept + slope*x[i]
		residualSum += (y[i] - yPred) * (y[i] - yPred)
	}
	stdError := math.Sqrt(residualSum/float64(len(x) - 2))

	return slope, intercept, meanX, ssx, stdError
}

func Predictiveinterval(x []float64, slope, intercept, meanX, ssx, stdError float64) (float64, float64) {
	N := float64(len(x))
	xNew := N
	// Predict y for new x
	ypred := intercept + slope*xNew

	// Calculate critical value for Z
	zvalue := 1.96

	// Calculate the prediction interval for the next value(x + 1)
	predictionInterv := zvalue * stdError * math.Sqrt(1.0+1.0/float64(N)+(xNew-meanX)*(xNew-meanX)/ssx)

	lowerbound := ypred - predictionInterv
	upperbound := ypred + predictionInterv

	return lowerbound, upperbound

}
