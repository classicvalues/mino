package test

import (
	"testing"

	"github.com/WatchBeam/mino"
	"github.com/WatchBeam/mino/transform"
	"github.com/stretchr/testify/assert"
)

func TestStddevNormal(t *testing.T) {
	result, err := mino.Analyze(fibonacci).Transform(transform.Stddev{})
	assert.Nil(t, err)

	r := result.(transform.StddevResults)
	assert.InEpsilon(t, 6.6097, r.Deviation, epsilon)
	assert.InEpsilon(t, 6.7500, r.Average, epsilon)
}

func TestStddevNormalEmpty(t *testing.T) {
	result, err := mino.Analyze(empty).Transform(transform.Stddev{})
	r := result.(transform.StddevResults)
	assert.Nil(t, err)
	assert.InEpsilon(t, 0, r.Deviation, epsilon)
	assert.InEpsilon(t, 0, r.Average, epsilon)
}

func TestStddevNormalOne(t *testing.T) {
	result, err := mino.Analyze(one).Transform(transform.Stddev{})
	r := result.(transform.StddevResults)
	assert.Nil(t, err)
	assert.InEpsilon(t, 0, r.Deviation, epsilon)
	assert.InEpsilon(t, 1, r.Average, epsilon)
}

func TestStddevBessel(t *testing.T) {
	result, err := mino.Analyze(fibonacci).Transform(transform.Stddev{Bessel: true})
	assert.Nil(t, err)

	r := result.(transform.StddevResults)
	assert.InEpsilon(t, 7.0660, r.Deviation, epsilon)
	assert.InEpsilon(t, 6.7500, r.Average, epsilon)
}

func TestStddevBesselEmpty(t *testing.T) {
	result, err := mino.Analyze(empty).Transform(transform.Stddev{Bessel: true})
	r := result.(transform.StddevResults)
	assert.Nil(t, err)
	assert.InEpsilon(t, 0, r.Deviation, epsilon)
	assert.InEpsilon(t, 0, r.Average, epsilon)
}

func TestStddevBesselOne(t *testing.T) {
	result, err := mino.Analyze(one).Transform(transform.Stddev{Bessel: true})
	r := result.(transform.StddevResults)
	assert.Nil(t, err)
	assert.InEpsilon(t, 0, r.Deviation, epsilon)
	assert.InEpsilon(t, 1, r.Average, epsilon)
}
