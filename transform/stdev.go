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

// Performs a standard deviation, and returns StdevResults. Based
// on the algorithim from RosettaCode, which can be found at:
// http://rosettacode.org/wiki/Standard_deviation#Go
func (s Stddev) Transform(analyzer *mino.Analyzer, data []float64) (interface{}, error) {
	// Special case: no data in set
	if len(data) == 0 {
		return StddevResults{0, 0}, nil
	}

	var a, q float64
	n := float64(0)
	for _, x := range data {
		n += 1
		a1 := a + (x-a)/n
		q, a = q+(x-a)*(x-a1), a1
	}

	if s.Bessel && n > 1 {
		n -= 1
	}

	return StddevResults{
		Average:   a,
		Deviation: math.Sqrt(q / n),
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
