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
	data := map[string]interface{}{
		"Sting Slice":   []string{"A", "V", "C"},
		"Integer Slice": []int{5, 50, 500},
		"Float Slice":   []float64{0.11, 0.22, 0.33},
		"Empty Slice":   []int{},
		"String":        "00000",
		"Empty String":  "",
	}

	for key, v := range data {
		out := Tail(v)
		if out == nil {
			t.Fatalf("\n Unexpected error %s Tail value :[%v] and type: %T", key, out, out)
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

func TestIndexOf(t *testing.T) {
	data := map[string]interface{}{
		"Sting Slice":   []string{"A", "V", "C"},
		"Integer Slice": []int{5, 50, 500},
		"Float Slice":   []float64{0.11, 0.22, 0.33},
	}

	for key, v := range data {
		index := rand.Intn(len(data))
		out := IndexOf(index, v)
		fmt.Printf("\n %s IndexOf[%d] value :[%v] and type: %T", key, index, out, out)
	}
}
