package goramda

import (
	// "go/types"
	"fmt"
	"reflect"
)

// TODO: following have to implement
/*
Head
Tail
Nth
IndexOf
LastIndexOf
Find
FindLast
FindIndex
FindLastIndex
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

Update


*/

var typeError func(expected, given interface{}) error = func(e, g interface{}) error {
	return fmt.Errorf("Type mismatch expected %T got %T", e, g)
}

// Returns the first element of the given list or string. In some libraries this function is named first.
func Head(d interface{}) interface{} {
	arrV := reflect.ValueOf(d)
	switch arrV.Kind() {
	case reflect.String, reflect.Slice:
		return nth(0, arrV)
	case reflect.Ptr:
		return nil
	default:
		return getDefaultValueOf(d)
	}
}

func finder(fn func(d interface{}) bool, list interface{}) (interface{}, int) {
	if isSlice(list) {
		arrV := reflect.ValueOf(list)
		for i := 0; i < arrV.Len(); i++ {
			v := arrV.Index(i).Interface()
			if fn(v) {
				return v, i
			}
		}
	}
	return nil, -1
}

func find(fn func(d interface{}) bool, list interface{}) interface{} {
	v, _ := finder(fn, list)
	return v
}

func findIndex(fn func(d interface{}) bool, list interface{}) int {
	_, index := finder(fn, list)
	return index
}

func Find(fn func(d interface{}) bool) func(list interface{}) interface{} {
	return func(list interface{}) interface{} {
		return find(fn, list)
	}
}

func FindIndex(fn func(d interface{}) bool) func(list interface{}) int {
	return func(list interface{}) int {
		return findIndex(fn, list)
	}
}

func isSlice(d interface{}) bool {
	return reflect.ValueOf(d).Kind() == reflect.Slice
}

func Tail(d interface{}) interface{} {
	arrV := reflect.ValueOf(d)
	switch arrV.Kind() {
	case reflect.String, reflect.Slice:
		return nth(arrV.Len()-1, arrV)
	case reflect.Ptr:
		return nil
	default:
		return getDefaultValueOf(d)
	}
}

func nth(index int, arrV reflect.Value) interface{} {
	if isIndexable(index, arrV) {
		return arrV.Index(index).Interface()
	}
	return nil
}

func isIndexable(index int, arrV reflect.Value) bool {
	ok := (arrV.Kind() == reflect.Slice || arrV.Kind() == reflect.String) &&
		arrV.Len() > index && index >= 0
	return ok
}

func IndexOf(val interface{}, d interface{}) int {
	return contains(val, d)
}

func LastIndexOf(val interface{}, d interface{}) int {
	return lastContains(val, d)
}

func Nth(index int, d interface{}) interface{} {
	arrV := reflect.ValueOf(d)
	return nth(index, arrV)
}

func Contains(s, elem interface{}) bool {
	return contains(s, elem) >= 0
}

func contains(s, elem interface{}) int {
	arrV := reflect.ValueOf(s)
	if arrV.Kind() == reflect.Slice {
		for i := 0; i < arrV.Len(); i++ {
			if arrV.Index(i).Interface() == elem {
				return i
			}
		}
	}
	return -1
}

func lastContains(s, elem interface{}) int {
	arrV := reflect.ValueOf(s)
	if arrV.Kind() == reflect.Slice {
		for i := arrV.Len() - 1; i >= 0; i-- {
			if arrV.Index(i).Interface() == elem {
				return i
			}
		}
	}
	return -1
}

func drop(start, end int, s interface{}) interface{} {
	n := reflect.ValueOf(s)
	rest := reflect.MakeSlice(reflect.TypeOf(s), 0, 0)
	for i := start; i < end; i++ {
		rest = reflect.Append(rest, n.Index(i))
	}
	r := rest.Interface()
	return r
}

func Drop(count int, d interface{}) interface{} {
	v := reflect.ValueOf(d)
	if v.Kind() != reflect.Slice {
		return d
	}
	return drop(count, v.Len(), d)
}

func DropLast(count int, d interface{}) interface{} {
	v := reflect.ValueOf(d)
	if v.Kind() != reflect.Slice {
		return d
	}
	return drop(0, v.Len()-(1+count), d)
}

func append(arrV interface{}, s ...interface{}) interface{} {
	n := reflect.ValueOf(arrV)
	n2 := reflect.ValueOf(s)
	newV := reflect.MakeSlice(reflect.TypeOf(s), 0, 0)
	for i := 0; i < n.Len(); i++ {
		newV = reflect.Append(newV, n.Index(i))
	}
	for i := 0; i < add.Len(); i++ {
		newV = reflect.Append(newV, n2.Index(i))
	}
	return newV.Interface()
}

// func All(condFunc func(v, iter interface{}) bool, list interface{}) bool {
// 	condFunc
// }

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
