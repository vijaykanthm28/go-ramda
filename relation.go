package goramda

import (
	// "fmt"
	"reflect"
)

// Equals returns true if its arguments are equivalent, false otherwise. Handles
// cyclical data structures.
/*
  R.equals(1, 1); //=> true
  R.equals(1, '1'); //=> false
  R.equals([1, 2, 3], [1, 2, 3]); //=> true
*/

func NotEquals(a, b interface{}) bool {
	return !Equals(a, b)
}

func Equals(a, b interface{}) bool {
	return reflect.DeepEqual(a, b)
}

/*

// HasElem checks if a given slice of elements contains the provided single element value.
// if element available return true
func Gt(a, b interface{}) (bool, error) {
	aV := reflect.ValueOf(a)
	bV := reflect.ValueOf(b)
	if aV.Kind() != bV.Kind() {
		return false, fmt.Errorf("Type mismatch (%T,  %T)", a, b)
	}
	switch v := ttest.(type) {
	case int:
		fmt.Println(v > ee.(int))
	default:
		fmt.Println(ee)

	}
	return false, fmt.Errorf("invalid operation: aV > bV (operator > not defined on %T)", a)
}
*/
