package test

import (
	"testing"

	"github.com/WatchBeam/mino"
	"github.com/WatchBeam/mino/transform"
	"github.com/stretchr/testify/assert"
)

func TestQuartiles(t *testing.T) {
	result, err := mino.Analyze(fibonacci).Transform(transform.Quartile{Sections: 2})
	assert.Nil(t, err)
	assert.Equal(t, []float64{4}, result.([]float64))

	result, err = mino.Analyze(fibonacci).Transform(transform.Quartile{Sections: 3})
	assert.Nil(t, err)
	assert.Equal(t, []float64{2, 8}, result.([]float64))

	result, err = mino.Analyze(fibonacci).Transform(transform.Quartile{})
	assert.Nil(t, err)
	assert.Equal(t, []float64{1.25, 4, 11.75}, result.([]float64))
}
