package test

import (
	"testing"

	"github.com/WatchBeam/mino"
	"github.com/WatchBeam/mino/transform"
	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	result, err := mino.FromList(fibonacci).Transform(transform.Sum{})
	assert.Nil(t, err)
	assert.Equal(t, float64(54), result.(float64))
}

func TestSumEmpty(t *testing.T) {
	result, err := mino.FromList(empty).Transform(transform.Sum{})
	assert.Nil(t, err)
	assert.Equal(t, float64(0), result.(float64))
}

func TestSumWeighted(t *testing.T) {
	result, err := mino.FromPoints(weightedFib).Transform(transform.Sum{})
	assert.Nil(t, err)
	assert.InEpsilon(t, 114.5, result.(float64), epsilon)
}

func TestAverage(t *testing.T) {
	result, err := mino.FromList(fibonacci).Transform(transform.Average{})
	assert.Nil(t, err)
	assert.Equal(t, float64(6.75), result.(float64))
}

func TestAverageEmpty(t *testing.T) {
	result, err := mino.FromList(empty).Transform(transform.Average{})
	assert.Nil(t, err)
	assert.Equal(t, float64(0), result.(float64))
}

func TestAverageWeighted(t *testing.T) {
	result, err := mino.FromPoints(weightedFib).Transform(transform.Average{})
	assert.Nil(t, err)
	assert.InEpsilon(t, 8.4814, result.(float64), epsilon)
}
