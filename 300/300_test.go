package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCase1(t *testing.T) {
	assert.Equal(t, 4, lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}))
}

func TestCase2(t *testing.T) {
	assert.Equal(t, 4, lengthOfLIS([]int{0, 1, 0, 3, 2, 3}))
}
