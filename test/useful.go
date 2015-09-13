package test

import (
	"github.com/WatchBeam/mino"
)

var weightedFib = []mino.DataPoint{
	mino.DataPoint{1, 1.0},
	mino.DataPoint{21, 1.5},
	mino.DataPoint{5, 2},
	mino.DataPoint{1, 2.5},
	mino.DataPoint{8, 3},
	mino.DataPoint{13, 3.5},
}
var fibonacci = []float64{1, 21, 5, 1, 2, 8, 3, 13}
var empty = []float64{}
var one = []float64{1}
var epsilon = 0.001
