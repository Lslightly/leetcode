package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCase(t *testing.T) {
	assert.Len(t, subsets([]int{1, 2, 3}), 8)
	assert.Len(t, subsets([]int{0}), 2)
}
