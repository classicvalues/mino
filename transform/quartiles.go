package transform

import (
	"math"
	"sort"

	"github.com/WatchBeam/mino"
)

const epsilon = 0.00000001

// Quartiles segments the data into 4 sections, by recursively
// taking the median of each section.
type Quartile struct {
	// Number of sections to cut into. Defaults to 4 sections.
	Sections int
}

// Slices the data into quartiles.
func (q Quartile) Transform(analyzer *mino.Analyzer, data []float64) (interface{}, error) {
	target := make([]float64, len(data))
	copy(target, data)
	sort.Float64s(target)

	sections := 4
	if q.Sections > 0 {
		sections = q.Sections
	}

	size := float64(len(target)+1) / float64(sections)
	results := make([]float64, sections-1)
	for i := 1; i < sections; i++ {
		idx, frac := math.Modf(float64(i)*size - 1)
		half := int(idx)
		if frac < epsilon {
			results[i-1] = target[half]
		} else {
			results[i-1] = target[half]*(1-frac) + target[half+1]*frac
		}
	}

	return results, nil
}
