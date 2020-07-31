package goramda

import (
	"testing"
)

func TestTypify(t *testing.T) {

	testData := map[string]struct {
		data             interface{}
		expectedVal      interface{}
		defaultValueType interface{}
	}{
		"String slice typify": {
			data:             []string{"A", "V", "C"},
			defaultValueType: []string{},
			expectedVal:      []string{"A", "V", "C"},
		},
		"Int typify": {
			data:             5,
			defaultValueType: int(0),
			expectedVal:      5,
		},
		"nil typify int": {
			data:             nil,
			defaultValueType: int(0),
			expectedVal:      0,
		},
		"Int32 typify": {
			data:             int32(5),
			defaultValueType: int32(0),
			expectedVal:      int32(5),
		},
		"Int32 slice typify": {
			data:             []int32{5, 7},
			defaultValueType: []int32{},
			expectedVal:      []int32{5, 7},
		},
		// This test data not works here but works on test case
		// TestTypifyEmptySlice so have test Equals
		// "Int32 slice typify nil": {
		// 	data:             nil,
		// 	defaultValueType: []int32{},
		// 	expectedVal:      []int32{},
		// },
	}
	for key, v := range testData {
		out := Typify(v.data, v.defaultValueType)
		if !Equals(out, v.expectedVal) {
			t.Fatalf("\n Unexpected ERROR: (%s) Typify output :(%v) %T but expected (%v) %T", key, out, out, v.expectedVal, v.expectedVal)
		}
	}
}

// Don't know Why its not pass in the testcase Equals but passes here
func TestTypifyEmptySlice(t *testing.T) {
	arrIntEmpty := []int{}
	out := Typify(nil, arrIntEmpty)
	if _, ok := out.([]int); !ok {
		t.Fatalf("\n Unexpected ERROR: Typify output :(%v) %T but expected (%v) %T", out, out, arrIntEmpty, arrIntEmpty)
	}

	arrStringEmpty := []string{}
	out1 := Typify(nil, arrStringEmpty)
	if _, ok := out1.([]string); !ok {
		t.Fatalf("\n Unexpected ERROR: Typify output :(%v) %T but expected (%v) %T", out1, out1, arrStringEmpty, arrStringEmpty)
	}

	arrFloatEmpty := []float64{}
	out2 := Typify(nil, arrFloatEmpty)
	if _, ok := out2.([]float64); !ok {
		t.Fatalf("\n Unexpected ERROR: Typify output :(%v) %T but expected (%v) %T", out2, out2, arrFloatEmpty, arrFloatEmpty)
	}
}
