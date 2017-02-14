package testingExtensions

import "testing"

// Testing interface to test unit
type Testing interface {
	AssertTrue(value bool) Testing
	AssertFalse(value bool) Testing
	AssertBoolEqual(expected, value bool) Testing
	AssertIntEqual(expected, value int) Testing
	AssertFloatEqual(expected, value float64) Testing
	AssertStringEqual(expected, value string) Testing
}

// TestUnit is a test unit
type TestUnit struct {
	T *testing.T
}

// AssertTrue tests if a boolean is true
func (t TestUnit) AssertTrue(value bool) Testing {
	return t.AssertBoolEqual(true, value)
}

// AssertFalse tests if a boolean is false
func (t TestUnit) AssertFalse(value bool) Testing {
	return t.AssertBoolEqual(false, value)
}

// AssertBoolEqual tests if a boolean is equal
func (t TestUnit) AssertBoolEqual(expected, value bool) Testing {
	if value != expected {
		t.T.Errorf("Expected %t but got %t", expected, value)
	}
	return t
}

// AssertIntEqual tests if ints are equal
func (t TestUnit) AssertIntEqual(expected, value int) Testing {
	if value != expected {
		t.T.Errorf("Expected %d but got %d", expected, value)
	}
	return t
}

// AssertStringEqual tests if strings are equal
func (t TestUnit) AssertStringEqual(expected, value string) Testing {
	if value != expected {
		t.T.Errorf("Expected %s but got %s", expected, value)
	}
	return t
}

// AssertFloatEqual tests if float64 are equal
func (t TestUnit) AssertFloatEqual(expected, value float64) Testing {
	if value != expected {
		t.T.Errorf("Expected %f but got %f", expected, value)
	}
	return t
}
