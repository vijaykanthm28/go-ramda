package goramda

import (
	"fmt"
	"reflect"
	"testing"
)

type MainStruct struct {
	A   int
	B   string
	Sub Sub
}

type Sub struct {
	Sa       string
	SuperSub SuperSub
}

type SuperSub struct {
	AJ float64
}

var tt = MainStruct{
	A: 23,
	B: "skidoo",
	Sub: Sub{
		Sa: "Welcome",
		SuperSub: SuperSub{
			AJ: 5.2,
		},
	},
}

func TestIsEmpty(t *testing.T) {
	data := []struct {
		name           string
		element        interface{}
		expectedResult bool
	}{
		{
			name:           "empty int",
			element:        0,
			expectedResult: true,
		},
		{
			name:           "value int",
			element:        3,
			expectedResult: false,
		},
		{
			name:           "empty string",
			element:        "",
			expectedResult: true,
		},
		{
			name:           "value string",
			element:        "Welcome",
			expectedResult: false,
		},
		{
			name:           "default bool",
			element:        false,
			expectedResult: true,
		},
		{
			name:           "empty struct",
			element:        SuperSub{},
			expectedResult: true,
		},
		{
			name:           "value struct",
			element:        SuperSub{AJ: 5.2},
			expectedResult: false,
		},
		{
			name:           "empty struct main",
			element:        MainStruct{},
			expectedResult: true,
		},
		{
			name:           "empty pointer struct main",
			element:        &MainStruct{},
			expectedResult: true,
		},
		{
			name:           "value pointer struct main",
			element:        MainStruct{A: 23, B: "skidoo"},
			expectedResult: false,
		},
		{
			name:           "empty nil",
			element:        nil,
			expectedResult: true,
		},
	}
	fmt.Println("Total data :", reflect.DeepEqual(&MainStruct{A: 23, B: "skidoo"}, &MainStruct{}))
	for _, d := range data {
		result := IsEmpty(d.element)
		if result != d.expectedResult {
			fmt.Printf("[%s] data failed ", d.name)
			t.Fatalf("Unexpected result expected (%v) for element (%#v) got: %v", d.expectedResult, d.element, result)
		}
	}
}

func TestHasPath(t *testing.T) {
	path := []string{"Sub", "SuperSub", "AJ"}
	if !HasPath(path, &tt) {
		t.Fatal("Unexpected result has path expected true got: false")
	}
}

func TestPathOr(t *testing.T) {
	path := []string{"Subb"}
	expected := Sub{}
	result := PathOr(expected, path, &tt).(Sub)
	if expected != result {
		t.Fatalf("Unexpected result pathOr expected (%T) got (%T) ", expected, result)
	}
}

func TestFirstLevelFieldOnPathOr(t *testing.T) {
	path := []string{"A"}
	expected := 23
	result := PathOr(0, path, &tt).(int)
	if expected != result {
		t.Fatalf("Unexpected result pathOr expected (%T) got (%T) ", expected, result)
	}
}

func TestPropsSatisfies(t *testing.T) {
	validator := func(d interface{}) bool {
		val, ok := d.(Sub)
		if ok && val == tt.Sub {
			return true
		}
		return false
	}
	if !PropSatisfies(validator, "Sub", &tt) {
		t.Fatal("Unexpected result ")
	}
}
