package goramda

import (
	// "reflect"
	// "fmt"
	// "go/types"
	// "math/rand"
	"fmt"
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

func TestHasPath(t *testing.T) {
	path := []string{"Sub", "SuperSub", "AJ"}
	if !HasPath(path, &tt) {
		t.Fatal("Unexpected result")
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
	if !PropSatisfies("Sub", &tt) {
		t.Fatal("Unexpected result ")
	}
}
