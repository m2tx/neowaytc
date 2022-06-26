package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseString(t *testing.T) {
	s := Reverse("0123456789")
	assert.Equal(t, "9876543210", s)
}

func TestAllDigitsEqualsTrue(t *testing.T) {
	assert.True(t, AllDigitsEquals("0000000000000000000000"))
}

func TestAllDigitsEqualsFalse(t *testing.T) {
	assert.False(t, AllDigitsEquals("01"))
}

func TestToInt(t *testing.T) {
	assert.Equal(t, ToInt("1"[0]), 1)
}

func TestStringToInt(t *testing.T) {
	assert.Equal(t, StringToInt("100"), 100)
}
