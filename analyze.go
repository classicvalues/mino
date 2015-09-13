package mino

type Analyzer struct {
	data Collection
}

func Analyze(data Collection) *Analyzer {
	return &Analyzer{
		data: data,
	}
}

// Runs a transform on the analyzer's contained data.
func (a *Analyzer) Transform(t Transform) (interface{}, error) {
	defer a.data.Reset()
	return t.Transform(a, a.data)
}
