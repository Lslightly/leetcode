package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCase1(t *testing.T) {
	assert.Equal(t, 7, waysToReachTarget(6, [][]int{
		{6, 1},
		{3, 2},
		{2, 3},
	}))
}

func TestCase2(t *testing.T) {
	assert.Equal(t, 1, waysToReachTarget(18, [][]int{
		{6, 1},
		{3, 2},
		{2, 3},
	}))
}
