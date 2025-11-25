package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCase1(t *testing.T) {
	assert.Equal(t, 3, climbStairs(3))
}

func TestCase2(t *testing.T) {
	assert.Equal(t, 5, climbStairs(4))
}
