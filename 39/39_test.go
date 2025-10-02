package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCase(t *testing.T) {
	assert.Len(t, combinationSum([]int{2, 3, 6, 7}, 7), 2)
	assert.Len(t, combinationSum([]int{2, 3, 5}, 8), 3)
}
