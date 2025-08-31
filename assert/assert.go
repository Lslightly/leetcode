package assert

import (
	"runtime/debug"
	"testing"
)

func Equal[T comparable](t *testing.T, expected, actual T, msg string) {
	if expected != actual {
		t.Fatalf("Not Equal: expected=%v, actual=%v, msg=%s", expected, actual, msg)
		debug.PrintStack()
	}
}

func NotEqual[T comparable](t *testing.T, notExpected, actual T, msg string) {
	if notExpected == actual {
		t.Fatalf("Equal: notExpected=%v, actual=%v, msg=%s", notExpected, actual, msg)
		debug.PrintStack()
	}
}

func Sastisfy(t *testing.T, condition bool, msg string) {
	if !condition {
		t.Fatalf("Not Sastisfy: msg=%s", msg)
		debug.PrintStack()
	}
}
