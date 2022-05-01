package goramda

import (
	// "fmt"
	"reflect"
)

/*
Have to implement

difference
Union
*/

// Equals returns true if its arguments are equivalent, false otherwise. Handles
// cyclical data structures.
/*
  R.equals(1, 1); //=> true
  R.equals([1, 2, 3], [1, 2, 3]); //=> true
*/

func Union(s1, s2 interface{}) interface{} {
	arrV1 := reflect.ValueOf(s1)
	arrV2 := reflect.ValueOf(s2)
	if arrV2.Kind() == arrV2.Kind() && arrV2.Kind() == reflect.Slice {
		for i := 0; i < arrV2.Len(); i++ {
			if contains(s1, arrV2.Index(i).Interface()) < 0 {
				arrV1 = reflect.Append(arrV1, arrV2.Index(i))
			}
		}
	}
	return arrV1.Interface()
}

func Difference(slice1, slice2 interface{}) interface{} {
	arrV1 := reflect.ValueOf(slice1)
	arrV2 := reflect.ValueOf(slice2)
	if arrV2.Kind() != arrV1.Kind() || arrV1.Kind() != reflect.Slice {
		return nil
	}
	// diff := getDefaultValueOf(slice2)
	diff := reflect.MakeSlice(reflect.TypeOf(slice2), 0, 0)
	// 	var sliceType reflect.Type
	// sliceType = reflect.SliceOf(length, t)
	// return reflect.Zero(sliceType)
	for i := 0; i < arrV1.Len(); i++ {
		if contains(slice2, arrV1.Index(i).Interface()) < 0 {
			diff = reflect.Append(diff, arrV1.Index(i))
		}
	}

	return diff.Interface()
}

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
