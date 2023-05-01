package wallet

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOXENToDecimal(t *testing.T) {
	assert.Equal(t, "0.034000200000", OXENToDecimal(34000200000))
	assert.Equal(t, "15.000000000000", OXENToDecimal(15e12))
}

func TestOXENToFloat64(t *testing.T) {
	assert.Equal(t, float64(0.02), OXENToFloat64(20000000000))
	assert.Equal(t, float64(3.14), OXENToFloat64(314e10))
}
