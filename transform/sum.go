package transform

import (
	"github.com/WatchBeam/mino"
)

// Sum is a transform which simply totals all the values in the data.
type Sum struct{}

var _ mino.Transform = Sum{}

// Sums all the data values and returns a float64.
func (s Sum) Transform(analyzer *mino.Analyzer, data mino.Collection) (interface{}, error) {
	return baseSum(data), nil
}

func baseSum(data mino.Collection) float64 {
	total := float64(0)
	for data.HasMore() {
		x := data.Next()
		total += x.Value * x.Weight
	}

	return total
}

// Sum is a transform that calculates the average of a data set.
type Average struct{}

// Calculates the weighted average of a data set.
func (a Average) Transform(analyzer *mino.Analyzer, data mino.Collection) (interface{}, error) {
	average := float64(0)
	n := float64(0)
	for data.HasMore() {
		x := data.Next()
		n += x.Weight
		average = average + (x.Value-average)/(n/x.Weight)
	}

	return average, nil
}
