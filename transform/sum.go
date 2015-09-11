package transform

import (
	"github.com/WatchBeam/mino"
)

// Sum is a transform which simply totals all the values in the data.
type Sum struct{}

var _ mino.Transform = Sum{}

// Sums all the data values and returns a float64.
func (s Sum) Transform(analyzer *mino.Analyzer, data []float64) (interface{}, error) {
	return baseSum(data), nil
}

func baseSum(data []float64) float64 {
	total := float64(0)
	for _, value := range data {
		total += value
	}

	return total
}

// Sum is a transform that calculates the average of a data set.
type Average struct{}

// Calculates the average of a data set.
func (a Average) Transform(analyzer *mino.Analyzer, data []float64) (interface{}, error) {
	average := float64(0)
	n := float64(0)
	for _, x := range data {
		n += 1
		average = average + (x-average)/n
	}

	return average, nil
}
