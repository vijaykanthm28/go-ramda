package goramda

import (
	"fmt"
	"reflect"

	"math/rand"
	"testing"
)

func TestHead(t *testing.T) {
	data := map[string]interface{}{
		"Sting Slice":   []string{"A", "V", "C"},
		"Integer Slice": []int{5, 50, 500},
		"Float Slice":   []float64{0.11, 0.22, 0.33},
	}
	for key, v := range data {
		out := Head(v)
		if out == nil {
			t.Fatalf("\n Unexpected error %s Tail value :[%v] and type: %T", key, out, out)
		}
	}
}

func TestTail(t *testing.T) {
	data := map[string]struct {
		data        interface{}
		expectedVal interface{}
	}{
		"Sting Slice": {
			data:        []string{"A", "V", "C"},
			expectedVal: []string{"A", "V", "C"},
		},
		"Int Slice": {
			data:        []int{5, 50, 500, 5000},
			expectedVal: []int{5, 50, 500, 5000},
		},
		"Float Slice": {
			data:        []float64{0.11, 0.22, 0.23, 0.24, 0.33},
			expectedVal: []float64{0.11, 0.22, 0.23, 0.24, 0.33},
		},
		"Empty Slice": {
			data:        []int{},
			expectedVal: []int{},
		},
		"Sting": {
			data:        "2547893",
			expectedVal: "547893",
		},
		"EmptySting": {
			data:        "",
			expectedVal: "",
		},
	}

	for key, v := range data {
		out := Tail(v.data)
		if NotEquals(out, v.expectedVal) {
			t.Fatalf("\n Unexpected error %s Tail value :[%v] and type: %T", key, out, out)
		}
	}
}

func TestLast(t *testing.T) {
	data := map[string]struct {
		data        interface{}
		expectedVal interface{}
	}{
		"Sting Slice": {
			data:        []string{"A", "V", "C"},
			expectedVal: "C",
		},
		"Int Slice": {
			data:        []int{5, 50, 500, 5000},
			expectedVal: 5000,
		},
		"Float Slice": {
			data:        []float64{0.11, 0.22, 0.23, 0.24, 0.33},
			expectedVal: 0.33,
		},
		"Empty Slice": {
			data:        []int{},
			expectedVal: nil,
		},
		"Sting": {
			data:        "2547893",
			expectedVal: uint8('3'),
		},
		"EmptySting": {
			data:        "",
			expectedVal: nil,
		},
	}

	for key, v := range data {
		out := Last(v.data)
		if NotEquals(out, v.expectedVal) {
			t.Fatalf("\n Unexpected error on (%s) Tail value :[%v] and type: %T", key, out, out)
		}
	}
}

func TestDropTypeAssert(t *testing.T) {
	data := map[string]struct {
		data        interface{}
		expectedLen int
	}{

		"Sting Slice": {
			data:        []string{"A", "V", "C"},
			expectedLen: 1,
		},
		"Sting Slice 4": {
			data:        []string{"Aasd", "ddd", "Cccc", "dddd"},
			expectedLen: 2,
		},
	}

	for _, v := range data {
		out := Drop(2, v.data)
		outLen := reflect.ValueOf(out).Len()
		_, ok := out.([]string)
		if !ok || outLen != v.expectedLen {
			t.Fatalf("unexported result type assert (%v) length of array (%d) but expected: %d", ok, outLen, 1)
		}
	}
}

func TestDrop(t *testing.T) {
	data := map[string]interface{}{
		"Sting Slice":   []string{"A", "V", "C"},
		"Integer Slice": []int{5, 50, 500},
		"Float Slice":   []float64{0.11, 0.22, 0.33},
	}

	for _, v := range data {
		out := Drop(2, v)
		outLen := reflect.ValueOf(out).Len()
		if outLen != 1 {
			t.Fatalf("unexported length of array (%d) but expected: %d", outLen, 1)
		}
	}
}

func TestNth(t *testing.T) {
	data := map[string]interface{}{
		"Sting Slice":   []string{"A", "V", "C"},
		"Integer Slice": []int{5, 50, 500},
		"Float Slice":   []float64{0.11, 0.22, 0.33},
	}

	for key, v := range data {
		index := rand.Intn(len(data))
		out := Nth(index, v)
		fmt.Printf("\n %s IndexOf[%d] value :[%v] and type: %T", key, index, out, out)
	}
}

func TestFind(t *testing.T) {
	data := map[string]interface{}{
		"Sting Slice":   []string{"V", "I", "J", "A", "Y", "A", "K", "N", "T", "H"},
		"Integer Slice": []int{5, 50, 500, 2, 3, 5, 85, 96},
		"Float Slice":   []float64{0.11, 0.22, 0.33, 0.32, 2.2, 3.3, 5.5},
	}
	for key, v := range data {
		arrV := reflect.ValueOf(v)
		index := rand.Intn(arrV.Len() - 1)
		v1 := arrV.Index(index).Interface()
		findFunc := Find(func(d interface{}) bool {
			return Equals(d, v1)
		})
		out := findFunc(v)
		if v1 != out {
			t.Fatalf("Find value of (%v) unexpected result expected value [%v] got: %v", key, v1, out)
		}
	}
}

func TestFindIndex(t *testing.T) {
	data := map[string]interface{}{
		"Sting Slice":   []string{"V", "I", "J", "A", "Y", "A", "K", "N", "T", "H"},
		"Integer Slice": []int{5, 50, 500, 2, 3, 51, 85, 96},
		"Float Slice":   []float64{0.11, 0.22, 0.33, 0.32, 2.2, 3.3, 5.5},
	}
	for key, v := range data {
		arrV := reflect.ValueOf(v)
		index := rand.Intn(arrV.Len() - 1)
		v1 := arrV.Index(index).Interface()
		findFunc := FindIndex(func(d interface{}) bool {
			return Equals(d, v1)
		})
		out := findFunc(v)
		if out != index {
			t.Fatalf("findIndex of (%v) unexpected result expected index [%v] got: %v", key, index, out)
		}
	}
}
