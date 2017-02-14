package testingExtensions_test

import (
	"testing"

	"github.com/eviltofu/testingExtensions"
)

func TestBoolAssertTruePass(t *testing.T) {
	tT := new(testing.T)
	testUnit := testingExtensions.TestUnit{T: tT}
	testUnit.AssertTrue(true)
	expectPass(tT, t)
}

func TestBoolAssertTrueFail(t *testing.T) {
	tT := new(testing.T)
	testUnit := testingExtensions.TestUnit{T: tT}
	testUnit.AssertTrue(false)
	expectFail(tT, t)
}

func TestBoolAssertFalsePass(t *testing.T) {
	tT := new(testing.T)
	testUnit := testingExtensions.TestUnit{T: tT}
	testUnit.AssertFalse(false)
	expectPass(tT, t)
}

func TestBoolAssertFalseFail(t *testing.T) {
	tT := new(testing.T)
	testUnit := testingExtensions.TestUnit{T: tT}
	testUnit.AssertFalse(true)
	expectFail(tT, t)
}

func TestStringAssertEqual(t *testing.T) {
	type data struct {
		first, second string
		pass          bool
	}
	testData := []data{
		data{first: "abc", second: "abc", pass: true},
		data{first: "1234567890", second: "1234567890", pass: true},
		data{first: "abc", second: "ABC", pass: false},
		data{first: "1234567890", second: "0123456789", pass: false},
		data{first: "1234567890 ", second: "1234567890", pass: false},
		data{first: " a ", second: "a", pass: false},
		data{first: " a ", second: " A ", pass: false},
		data{first: "aa", second: "aaa", pass: false},
		data{first: " ", second: "  ", pass: false}}
	for _, value := range testData {
		tT := new(testing.T)
		testUnit := testingExtensions.TestUnit{T: tT}
		testUnit.AssertStringEqual(value.first, value.second)
		switch value.pass {
		case true:
			expectPass(tT, t)
		case false:
			expectFail(tT, t)
		}
	}
}

func TestIntAssertEqual(t *testing.T) {
	type data struct {
		first, second int
		pass          bool
	}
	testData := []data{
		data{first: 0, second: 0, pass: true},
		data{first: 0, second: 1, pass: false}}
	for _, value := range testData {
		tT := new(testing.T)
		testUnit := testingExtensions.TestUnit{T: tT}
		testUnit.AssertIntEqual(value.first, value.second)
		switch value.pass {
		case true:
			expectPass(tT, t)
		case false:
			expectFail(tT, t)
		}
	}
}

func TestFloatAssertEqual(t *testing.T) {
	type data struct {
		first, second float64
		pass          bool
	}
	testData := []data{
		data{first: 0, second: 0, pass: true},
		data{first: 0, second: 1, pass: false}}
	for _, value := range testData {
		tT := new(testing.T)
		testUnit := testingExtensions.TestUnit{T: tT}
		testUnit.AssertFloatEqual(value.first, value.second)
		switch value.pass {
		case true:
			expectPass(tT, t)
		case false:
			expectFail(tT, t)
		}
	}
}

func TestBoolAssertEqual(t *testing.T) {
	type data struct {
		first, second bool
		pass          bool
	}
	testData := []data{
		data{first: true, second: true, pass: true},
		data{first: true, second: false, pass: false},
		data{first: false, second: true, pass: false},
		data{first: false, second: false, pass: true}}
	for _, value := range testData {
		tT := new(testing.T)
		testUnit := testingExtensions.TestUnit{T: tT}
		testUnit.AssertBoolEqual(value.first, value.second)
		switch value.pass {
		case true:
			expectPass(tT, t)
		case false:
			expectFail(tT, t)
		}
	}
}

// Private functions

func expectPass(tT *testing.T, t *testing.T) {
	if tT.Failed() != false {
		t.Error("Expected pass but got fail")
	}
}

func expectFail(tT *testing.T, t *testing.T) {
	if tT.Failed() != true {
		t.Error("Expected fail but got pass")
	}
}
