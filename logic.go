package goramda

import (
	"reflect"
)

// TODO: following have to implement
/*
And
Or
Not
pathSatisfies // have to do minnor change (first or is func )
propSatisfies
IfElse
hasPath
pathOr
pathEq

May implement
until
when
*/

func IfElse(cond bool, trueBlock interface{}, falseBlock interface{}) interface{} {
	if cond == true {
		return trueBlock
	}
	return falseBlock
}

func And(expressions ...bool) bool {
	return !Contains(expressions, false)
}

func Or(expressions ...bool) bool {
	return !Contains(expressions, true)
}

func DeepFields(index int, path []string, iface interface{}, result interface{}) (field []string) {
	setResult := func() {
		val := reflect.ValueOf(result)
		if result != nil {
			val.Elem().Set(reflect.ValueOf(iface))
		}
	}

	if len(path) <= index {
		setResult()
		return
	}
	ifv := reflect.ValueOf(iface)
	switch ifv.Kind() {
	case reflect.Struct:
		v := ifv.FieldByName(path[index])
		if !v.IsValid() {
			return
		}
		field = append(field, path[index])
		subField := DeepFields(index+1, path, v.Interface(), result)
		if len(subField) > 0 {
			field = append(field, subField...)
		}
	case reflect.Ptr:
		subField := DeepFields(index, path, ifv.Elem().Interface(), result)
		if len(subField) > 0 {
			field = append(field, subField...)
		}
	default:
		v := ifv.FieldByName(path[index])
		if v.IsValid() {
			field = append(field, path[index])
		}
		setResult()
	}

	return
}

// Pointer elements fields find
// t := reflect.TypeOf(item)
// for i := 0; i < t.NumField(); i++ {
//     ft := t.Field(i).Type
//     if ft.Kind() == reflect.Ptr {
//         ft = ft.Elem()
//     }
//     fmt.Println(ft.Kind())
// }

// c := C{"foo", "bar", "baz"}

// s := reflect.ValueOf(dd).Elem()
// typeOfT := s.Type()
//
// for i := 0; i < s.NumField(); i++ {
// 	f := s.Field(i)
// 	fmt.Printf("%d: %s %s %s = %v\n", i,
// 		typeOfT.Field(i).Name, f.Type(), f.Kind(), f.Interface())
// 	if f.Kind() == reflect.Struct {
// 		s1 := reflect.ValueOf(f.Interface()).Elem()
// 		fmt.Println("Sub struct Fields", s1.NumField())
// 	}
// }

func PathOr(defaultResult interface{}, path []string, data interface{}) interface{} {
	var result interface{}
	f := DeepFields(0, path, data, &result)
	if Equals(path, f) {
		return result
	}
	return defaultResult
}

func HasPath(path []string, dd interface{}) bool {
	f := DeepFields(0, path, dd, nil)
	return Equals(path, f)
}

// TODO: have to change as function check
func PathSatisfies(path string, d interface{}) bool {
	return HasPath([]string{path}, d)
}

// TODO: have to change as function check
func PropSatisfies(path string, d interface{}) bool {
	return HasPath([]string{path}, d)
}

func IsEmpty(d interface{}) bool {
	return false
}
