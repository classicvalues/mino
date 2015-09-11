package mino

// A Transform operates over a data source and calculates "some"
// value for it.
type Transform interface {
	Transform(analyzer *Analyzer, data []float64) (result interface{}, err error)
}
