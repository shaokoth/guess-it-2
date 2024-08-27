package guess_it

import "math"

func PredictInt(x, y []float64, ye float64) (float64, float64) {
	z := 1.96
	yestimate := 0.0
	stdError := 0.0
	marginError := 0.0
	sumys := 0.0
	xval := 0.0
	sumx := 0.0
	sumx2 := 0.0
	xline := 0.0
	for i := 0; i < len(x); i++ {
		xline = x[i]
		N := len(x)
		sumys += (y[i] - yestimate) * (y[i] - yestimate)
		xval = xline - sumx/float64(N)
		sumx += x[i]
		sumx2 += x[i] - sumx/float64(N)
		stdError = math.Sqrt((sumys)/float64(N-2))
		marginError = z * stdError * math.Sqrt(1 + 1/float64(N) + xval/sumx2) 
	}
	return yestimate - marginError, yestimate + marginError
}
