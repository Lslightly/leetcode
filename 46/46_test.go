package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCase1(t *testing.T) {
	assert.Len(t, permute([]int{1, 2, 3}), 6)
}
