package goramda

import (
	"reflect"
)

// TODO: following have to implement
/*
And
Or
Not
pathSatisfies
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

func PathOr(defaultResult interface{}, path []string, data interface{}) interface{} {
	var result interface{}
	f := DeepFields(0, path, data, &result)
	if Equals(path, f) {
		return result
	}
	return defaultResult
}

type DeciderFun func(interface{}) bool

func PathEq(path []string, elm interface{}) DeciderFun {
	return func(data interface{}) bool {
		eq := func(d interface{}) bool { return Equals(d, elm) }
		return PathSatisfies(eq, path, data)
	}
}

func HasPath(path []string, dd interface{}) bool {
	f := DeepFields(0, path, dd, nil)
	return Equals(path, f)
}

func Path(path []string, data interface{}) interface{} {
	var result interface{}
	f := DeepFields(0, path, data, &result)
	if Equals(path, f) {
		return result
	}
	return nil
}

func PathSatisfies(fn DeciderFun, path []string, data interface{}) bool {
	var result interface{}
	f := DeepFields(0, path, data, &result)
	if Equals(path, f) {
		return fn(result)
	}
	return false
}

func PropSatisfies(fn func(d interface{}) bool, path string, d interface{}) bool {
	return PathSatisfies(fn, []string{path}, d)
}

func getDefaultValueOf(d interface{}) interface{} {
	if IsNil(d) {
		return nil
	}
	ifv := reflect.ValueOf(d)
	var defaultValue interface{}
	switch ifv.Kind() {
	case reflect.Ptr:
		elmType := reflect.TypeOf(d).Elem()
		elmPtr2 := reflect.New(elmType)
		defaultValue = elmPtr2.Interface()
	default:
		ptrOfElm := reflect.New(reflect.TypeOf(d))
		defaultValue = ptrOfElm.Elem().Interface()
	}
	return defaultValue
}

func IsNil(d interface{}) bool {
	return d == nil
}

// IsEmpty will return true if it has its default value
func IsEmpty(d interface{}) bool {
	if IsNil(d) {
		return true
	}
	defaultValue := getDefaultValueOf(d)
	return Equals(d, defaultValue)
}
