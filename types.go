package mino

import (
	"sort"
)

// A Transform operates over a data source and calculates "some"
// value for it.
type Transform interface {
	Transform(analyzer *Analyzer, data Collection) (result interface{}, err error)
}

type DataPoint struct {
	Value  float64
	Weight float64
}

// Sorting function for data points. Returns true iff a < b
type Sorter func(a, b DataPoint) bool

// A collection represents a group of values that can be analyzed.
// Data values in the collection are immutable, but the list's sorting
// may change over time.
type Collection interface {
	// Returns the next data point in the collection, and increments the
	// internal "reader". If there are no more data points available, the
	// method may panic.
	Next() DataPoint
	// Decrements the internal reader by the given amount.
	Rewind(n int)
	// Returns whether there are more data points available.
	HasMore() bool
	// Resets the collection counter, so that calling Next() once again
	// returns the first item in the collection.
	Reset()
	// Returns the total weight of all DataPoints.
	Weight() float64
	// Returns the number of values contained in the collection.
	Len() int
	// Sorts the underlying data list.
	Sort(Sorter)
}

type dataSet struct {
	data   []DataPoint
	sortFn Sorter
}

var _ sort.Interface = dataSet{}

func (d dataSet) Len() int {
	return len(d.data)
}
func (d dataSet) Swap(i, j int) {
	d.data[i], d.data[j] = d.data[j], d.data[i]
}
func (d dataSet) Less(i, j int) bool {
	return d.sortFn(d.data[i], d.data[j])
}

// Basic implementation of the Collection interface.
type baseCollection struct {
	d      dataSet
	idx    int
	weight float64
}

var _ Collection = new(baseCollection)

func (b *baseCollection) calculateWeight() {
	b.weight = 0
	for _, pnt := range b.d.data {
		b.weight += pnt.Weight
	}
}

func (b *baseCollection) Next() DataPoint {
	pnt := b.d.data[b.idx]
	b.idx += 1

	return pnt
}

func (b *baseCollection) Rewind(n int) {
	b.idx -= n
}

func (b *baseCollection) Reset() {
	b.idx = 0
}

func (b *baseCollection) HasMore() bool {
	return b.idx < len(b.d.data)
}

func (b *baseCollection) Weight() float64 {
	return b.weight
}

func (b *baseCollection) Len() int {
	return len(b.d.data)
}

func (b *baseCollection) Sort(fn Sorter) {
	b.d.sortFn = fn
	sort.Sort(b.d)
}

// Returns a new collection formed from the slice of unweighted float values.
func FromList(list []float64) *Analyzer {
	points := make([]DataPoint, len(list))
	for i, value := range list {
		points[i] = DataPoint{Value: value, Weight: 1.0}
	}

	return FromPoints(points)
}

// Returns a collection formed from the slice of data points.
func FromPoints(list []DataPoint) *Analyzer {
	col := &baseCollection{d: dataSet{data: list}}
	col.calculateWeight()

	return Analyze(col)
}
