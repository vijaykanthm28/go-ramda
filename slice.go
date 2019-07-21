package goramda

import (
	// "go/types"
	"reflect"
	"fmt"
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

var typeError func (expected, given interface{}) error = func (e, g interface{}) error {
	return fmt.Errorf("Type mismatch expected %T got %T", e,g)
}

// Returns the first element of the given list or string. In some libraries this function is named first.
func Head(d interface{}) interface{} {
	arrV := reflect.ValueOf(d)
	switch arrV.Kind() {
	case reflect.String, reflect.Slice:
		return indexOf(0, arrV)
	case reflect.Ptr:
		return nil
	default:
		return getDefaultValueOf(d)
	}
}

func Tail(d interface{}) interface{} {
	arrV := reflect.ValueOf(d)
	switch arrV.Kind() {
	case reflect.String, reflect.Slice:
		return indexOf(arrV.Len()-1, arrV)
	case reflect.Ptr:
		return nil
	default:
		return getDefaultValueOf(d)
	}
}

func indexOf(index int, arrV reflect.Value) interface{} {
	if isIndexable(index, arrV) {
    return arrV.Index(index).Interface()
	}
	return nil
}

func isIndexable(index int, arrV reflect.Value) bool {
	ok := arrV.Kind() == reflect.Slice && arrV.Kind() == reflect.String &&
	arrV.Len() > index && index >= 0
	return ok
}

func IndexOf(index int, d interface{}) interface{} {
	arrV := reflect.ValueOf(d)
	return indexOf(index, arrV)
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
