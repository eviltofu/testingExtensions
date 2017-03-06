package testingExtensions

import "testing"

// Testing interface to test unit
type Testing interface {
	AssertTrue(value bool)
	AssertFalse(value bool)
	AssertBoolEqual(expected, value bool)
}

// TestUnit is a test unit
type TestUnit struct {
	T *testing.T
}

// AssertTrue tests if a boolean is true
func (t TestUnit) AssertTrue(value bool) {
	t.AssertBoolEqual(true, value)
}

// AssertFalse tests if a boolean is false
func (t TestUnit) AssertFalse(value bool) {
	t.AssertBoolEqual(false, value)
}

// AssertBoolEqual tests if a boolean is equal
func (t TestUnit) AssertBoolEqual(expected, value bool) {
	if value != expected {
		t.T.Errorf("Expected %t but got %t", expected, value)
	}
}
