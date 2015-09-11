package mino

type Analyzer struct {
	data []float64
}

func Analyze(data []float64) *Analyzer {
	return &Analyzer{
		data: data,
	}
}

// Runs a transform on the analyzer's contained data.
func (a *Analyzer) Transform(t Transform) (interface{}, error) {
	return t.Transform(a, a.data)
}
