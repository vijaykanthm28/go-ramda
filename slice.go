package goramda

import (
	// "go/types"
	"reflect"
)

// TODO: following have to implement
/*
Adjust
All
Any
Aperture
Concat
Contains
Drop
DropLast
DropLastWhile
DropRepeats
DropWhile
EndsWith
Filter
Find

Update


*/
// Returns the first element of the given list or string. In some libraries this function is named first.
// TODO: support string first char
func Head(d interface{}) interface{} {
	return indexOf(0, d)
}

// TODO: support string last char
func Tail(d interface{}) interface{} {
	arrV, ok := isSlice(d)
	if ok {
		return indexOf(arrV.Len()-1, d)
	}
	return nil
}

func indexOf(index int, d interface{}) interface{} {
	arrV, ok := isSlice(d)
	if ok && arrV.Len() > index {
		return arrV.Index(index).Interface()
	}
	return nil
}

func isSlice(d interface{}) (reflect.Value, bool) {
	arrV := reflect.ValueOf(d)
	ok := arrV.Kind() == reflect.Slice
	return arrV, ok
}

func IndexOf(index int, d interface{}) interface{} {
	return indexOf(index, d)
}

func Contains(s, elem interface{}) bool {
	arrV := reflect.ValueOf(s)
	if arrV.Kind() == reflect.Slice {
		for i := 0; i < arrV.Len(); i++ {
			if arrV.Index(i).Interface() == elem {
				return true
			}
		}
	}
	return false
}

// Applies a function to the value at the given index of an array, returning a new
// copy of the array with the element at the given index replaced with the result of
// the function application.

// func Adjust() interface{} {
// 	arrV := reflect.ValueOf(s)
// 	if arrV.Kind() == reflect.Slice {
// 		for i := 0; i < arrV.Len(); i++ {
// 			// XXX - panics if slice element points to an unexported struct field
// 			// see https://golang.org/pkg/reflect/#Value.Interface
// 			if arrV.Index(i).Interface() == elem {
// 				return true
// 			}
// 		}
// 	}
// }
