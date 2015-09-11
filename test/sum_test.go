package test

import (
	"testing"

	"github.com/WatchBeam/mino"
	"github.com/WatchBeam/mino/transform"
	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	result, err := mino.Analyze(fibonacci).Transform(transform.Sum{})
	assert.Nil(t, err)
	assert.Equal(t, float64(54), result.(float64))

	result, err = mino.Analyze(empty).Transform(transform.Sum{})
	assert.Nil(t, err)
	assert.Equal(t, float64(0), result.(float64))
}

func TestAvergage(t *testing.T) {
	result, err := mino.Analyze(fibonacci).Transform(transform.Average{})
	assert.Nil(t, err)
	assert.Equal(t, float64(6.75), result.(float64))

	result, err = mino.Analyze(empty).Transform(transform.Average{})
	assert.Nil(t, err)
	assert.Equal(t, float64(0), result.(float64))
}
