package goramda

import (
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

// func Equals(a ...interface{}) func() interface{} {
// 	if len(a) == 1 {
// 		return func() interface{} {
// 			return Head(a)
// 		}
// 	}
//
// 	return func() interface{} {
// 		head := Head(a)
// 		for i := 1; i < len(a); i++ {
// 			if a[i] != head {
// 				return false
// 			}
// 		}
// 		return true
// 	}
// }

// // HasElem checks if a given slice of elements contains the provided single element value.
// // if element available return true
// func Gt(a, b interface{}) bool {
// 	aV := reflect.ValueOf(a)
// 	bV := reflect.ValueOf(b)
// 	if aV.Kind() == bV.Kind() {
// 		for i := 0; i < arrV.Len(); i++ {
// 			// XXX - panics if slice element points to an unexported struct field
// 			// see https://golang.org/pkg/reflect/#Value.Interface
// 			if aV > bV {
// 				return true
// 			}
// 		}
// 	}
// 	return false
// }
