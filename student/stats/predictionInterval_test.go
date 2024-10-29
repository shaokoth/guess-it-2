package guess_it

import (
	"testing"
)

func Test_mean(t *testing.T) {
	type args struct {
		values []float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Test1",
			args: args{values: []float64{1, 2, 3, 4, 5}},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Mean(tt.args.values); got != tt.want {
				t.Errorf("mean() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStandardDev(t *testing.T) {
	type args struct {
		input []float64
		mean  float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Test2",
			args: args{input: []float64{1, 2, 3, 4, 5}},
			want: 3.3166247903554,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StandardDev(tt.args.input, tt.args.mean); got != tt.want {
				t.Errorf("StandardDev() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLinearRegression(t *testing.T) {
	type args struct {
		y []float64
	}
	tests := []struct {
		name  string
		args  args
		wantM float64
		wantC float64
	}{
		{
			name:  "Simple positive slope",
			args:  args{y: []float64{1, 2, 3, 4, 5}},
			wantM: 1,
			wantC: 1,
		},
		{
			name:  "Simple negative slope",
			args:  args{y: []float64{5, 4, 3, 2, 1}},
			wantM: -1,
			wantC: 5,
		},
		{
			name:  "Horizontal line",
			args:  args{y: []float64{3, 3, 3, 3, 3}},
			wantM: 0,
			wantC: 3,
		},
		{
			name:  "Vertical shift",
			args:  args{y: []float64{5, 5, 5, 5, 5}},
			wantM: 0,
			wantC: 5,
		},
		{
			name:  "Random points",
			args:  args{y: []float64{2, 4, 5, 4, 5}},
			wantM: 0.6,
			wantC: 2.8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotM, gotC := LinearRegression(tt.args.y)
			if gotM != tt.wantM {
				t.Errorf("LinearRegression() gotM = %v, want %v", gotM, tt.wantM)
			}
			if gotC != tt.wantC {
				t.Errorf("LinearRegression() gotC = %v, want %v", gotC, tt.wantC)
			}
		})
	}
}

func TestPearsonsCorrelationCoefficient(t *testing.T) {
	type args struct {
		y []float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name:     "Perfect Positive Correlation",
			args:   args{y: []float64{1, 2, 3, 4, 5}},
			want: 1, // when rounded its 1
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PearsonsCorrelationCoefficient(tt.args.y); got != tt.want {
				t.Errorf("PearsonsCorrelationCoefficient() = %v, want %v", got, tt.want)
			}
		})
	}
}
