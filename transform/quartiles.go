package transform

import (
	"github.com/WatchBeam/mino"
)

// Quartiles segments the data into 4 sections, by recursively
// taking the median of each section.
type Quartile struct {
	// Number of sections to cut into. Defaults to 4 sections.
	Sections int
}

// Slices the data into sections, returning a slice of float64s.
// Returns an error if the data set is empty.
func (q Quartile) Transform(analyzer *mino.Analyzer, data mino.Collection) (interface{}, error) {
	sections := q.Sections
	if sections == 0 {
		sections = 4
	}
	results := make([]float64, sections-1)
	dividend := (1 + data.Weight()) / float64(sections)

	// todo(connor4312): my algorithm below behaves oddly when the data
	// size is less than the number of requested sections. Whoever discovers
	// a solution to it is awarded three cookies.
	if data.Len() < sections {
		return results, InsufficientDataError
	}

	data.Sort(func(a, b mino.DataPoint) bool {
		return a.Value < b.Value
	})

	var weight float64
	var prev mino.DataPoint

SECTION:
	for i := 1; i < sections; i++ {
		target := float64(i) * dividend
		for data.HasMore() {
			current := data.Next()
			last := prev
			prev = current

			if weight+current.Weight > target {
				frac := (target - weight) / current.Weight
				results[i-1] = last.Value*(1-frac) + current.Value*frac
				data.Rewind(1)
				continue SECTION
			}

			weight += current.Weight
		}

		results[i-1] = prev.Value
	}

	return results, nil
}
