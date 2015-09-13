package test

import (
	"testing"

	"github.com/WatchBeam/mino"
	"github.com/WatchBeam/mino/transform"
	"github.com/stretchr/testify/assert"
)

func TestTwoQuartiles(t *testing.T) {
	result, err := mino.FromList(fibonacci).Transform(transform.Quartile{Sections: 2})
	assert.Nil(t, err)
	assert.Equal(t, []float64{4}, result.([]float64))
}

func TestThreeQuartiles(t *testing.T) {
	result, err := mino.FromList(fibonacci).Transform(transform.Quartile{Sections: 3})
	assert.Nil(t, err)
	assert.Equal(t, []float64{2, 8}, result.([]float64))
}

func TestFourQuartiles(t *testing.T) {
	result, err := mino.FromList(fibonacci).Transform(transform.Quartile{})
	assert.Nil(t, err)
	assert.Equal(t, []float64{1.25, 4, 11.75}, result.([]float64))
}

func TestQuartileEmpty(t *testing.T) {
	_, err := mino.FromList(empty).Transform(transform.Quartile{})
	assert.Equal(t, transform.InsufficientDataError, err)
}

func TestQuartileSameSectionAsNumbers(t *testing.T) {
	result, err := mino.FromList([]float64{1, 2}).Transform(transform.Quartile{Sections: 2})
	assert.Nil(t, err)
	assert.Equal(t, []float64{1.5}, result.([]float64))
}
