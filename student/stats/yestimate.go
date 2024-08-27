package guess_it

func Y_estimate(x, y []float64) (float64) {
	yestimate := 0.0
	sumx := 0.0
	sumy := 0.0
	sumxy := 0.0
	sumx2 := 0.0
	sumy2 := 0.0
	a := 0.0
	b := 0.0
	for i := 0; i < len(x); i++ {
		N := float64(len(x))
		sumx += float64(x[i])
		sumy += float64(y[i])
		sumxy += float64(x[i]) * float64(y[i])
		sumx2 += float64(x[i]) * float64(x[i])
		sumy2 += float64(y[i]) * float64(y[i])
		// Calculates the slope
		b = (N*sumxy - sumx*sumy) / (N*sumx2 - sumx*sumx)
		// Calculates the y-intercept
		a = (sumy / N) - (b * ((sumx) / N))
		// Calculates the estimate of y
		yestimate = a * float64(x[i]) + b
	}
	return yestimate
}
