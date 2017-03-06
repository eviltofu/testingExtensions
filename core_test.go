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
