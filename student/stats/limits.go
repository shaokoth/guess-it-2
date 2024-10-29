package guess_it

import "math"

func Range(yValues []float64) (int, int) {
	startCalc := len(yValues) - 4
	if startCalc < 0 {
		startCalc = 0
	}

	chunk := yValues[startCalc:]

	// predict the next value
	predicted := PredictNextValue(chunk)

	stdDev := StandardDev(chunk, Mean(chunk))

	lowerRange := predicted - stdDev*(2)
	upperRange := predicted + stdDev*(2)

	lowRange := int(math.Round(lowerRange))
	highRange := int(math.Round(upperRange))
	return lowRange, highRange
}
