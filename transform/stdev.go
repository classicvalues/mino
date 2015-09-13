package transform

import (
	"math"
	"math/rand"

	"github.com/WatchBeam/mino"
)

const (
	log2pi = 1.83787706640934548356065947281123527972279494727556682563430308096553139185452079538948659727190839524
)

// Calculates the standard deviation for data.
type Stddev struct {
	// Whether to use Bessel's correction (n - 1). See:
	// https://en.wikipedia.org/wiki/Bessel%27s_correction
	Bessel bool
}

var _ mino.Transform = Stddev{}

// Performs a standard deviation, and returns StddevResults. If the data
// set's length is less than 2, then an error will be returned along
// with results of a standard deviation and average of zero.
// These  results may be inspected, however some
// methods (such as LogProb) will fail.
//
// Based on the algorithm from RosettaCode, which can be found at:
// http://rosettacode.org/wiki/Standard_deviation#Go, and
// adapted for weighted standard deviation, the equation
// for which which equation which was found at:
// http://www.itl.nist.gov/div898/software/dataplot/refman2/ch2/weightsd.pdf
func (s Stddev) Transform(analyzer *mino.Analyzer, data mino.Collection) (interface{}, error) {
	l := data.Len()
	// Special case: not enough data in set
	if l == 0 {
		return StddevResults{0, 0}, InsufficientDataError
	} else if l == 1 {
		return StddevResults{data.Next().Value, 0}, InsufficientDataError
	}

	var average, variance float64
	n := float64(0)
	for data.HasMore() {
		x := data.Next()
		n += x.Weight
		a := average + (x.Value-average)/(n/x.Weight)
		variance, average = variance+(x.Value-average)*(x.Value-a)*x.Weight, a
	}

	besselLen := data.Len()
	if s.Bessel {
		besselLen -= 1
	}

	return StddevResults{
		Average: average,
		Deviation: math.Sqrt((float64(data.Len()) * variance) /
			(float64(besselLen) * data.Weight())),
	}, nil
}

type StddevResults struct {
	Average   float64
	Deviation float64
}

// Returns a random number normally distributed within the results.
func (s StddevResults) Random() float64 {
	return rand.NormFloat64()*s.Deviation + s.Average
}

// Returns the logarithmic probability of x.
func (s StddevResults) LogProb(x float64) float64 {
	return -math.Log(s.Deviation) - 0.5*log2pi - math.Pow(((x-s.Average)/s.Deviation), 2)/2.0
}
