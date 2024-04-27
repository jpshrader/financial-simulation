package assets

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_To2f(t *testing.T) {
	f1 := To2f(float64(1.2345))
	f2 := To2f(float64(2.3456))

	assert.Equal(t, float64(1.23), f1)
	assert.Equal(t, float64(2.35), f2)
}