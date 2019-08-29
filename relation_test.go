package goramda

import (
	"fmt"
	"reflect"
	"testing"
)

func TestUnion(t *testing.T) {
	data := []struct {
		name           string
		firstElement   interface{}
		secondElement  interface{}
		expectedResult interface{}
		hasError       bool
	}{
		{
			name:           "string array",
			firstElement:   []string{"Hello", "welcome", "to", "go", "ramda"},
			secondElement:  []string{"Hi", "welcome", "to", "Go", "Ramda"},
			expectedResult: []string{"Hello", "welcome", "to", "go", "ramda", "Hi", "Go", "Ramda"},
		},
		{
			name:           "Integer array",
			firstElement:   []int{1, 2, 3, 5, 7},
			secondElement:  []int{6, 7, 8, 1, 2, 2, 4},
			expectedResult: []int{1, 2, 3, 5, 7, 6, 8, 4},
		},
	}
	for _, d := range data {
		result := Union(d.firstElement, d.secondElement)
		fmt.Printf("\n union : %v :%v", result, d.expectedResult)
		if !reflect.DeepEqual(result, d.expectedResult) {
			t.Fatalf("Unexpected result expected result (%v) but got : (%v)", result, d.expectedResult)
		}
	}
}

func TestDifference(t *testing.T) {
	data := []struct {
		name           string
		firstElement   interface{}
		secondElement  interface{}
		expectedResult interface{}
		hasError       bool
	}{
		{
			name:           "string array",
			firstElement:   []string{"Hello", "welcome", "to", "go", "ramda"},
			secondElement:  []string{"Hi", "welcome", "to", "Go", "Ramda"},
			expectedResult: []string{"Hello", "go", "ramda"},
		},
		{
			name:           "Integer array",
			firstElement:   []int{1, 2, 3, 5, 7},
			secondElement:  []int{6, 7, 8, 1, 2, 2, 4},
			expectedResult: []int{3, 5},
		},
	}
	for _, d := range data {
		result := Difference(d.firstElement, d.secondElement)
		if !reflect.DeepEqual(result, d.expectedResult) {
			t.Fatalf("Unexpected result expected result (%v) but got : (%v)", d.expectedResult, result)
		}
	}
}
