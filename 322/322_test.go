package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCase1(t *testing.T) {
	assert.Equal(t, 3, coinChange([]int{1, 2, 5}, 11))
}

func TestCase2(t *testing.T) {
	assert.Equal(t, -1, coinChange([]int{2}, 3))
}
