package goramda

import (
	"reflect"
	"strconv"
)

// TODO: following have to implement
/*
And
Or
Not
IfElse
HasPath
PathOr
PathEq
PathSatisfies
PropSatisfies

May implement

until
when
*/

func And(expressions ...bool) bool {
	return !Includes(expressions, false)
}

func Or(expressions ...bool) bool {
	return !Includes(expressions, true)
}

func Not(expression bool) bool {
	return !expression
}

func IfElse(cond bool, trueBlock interface{}, falseBlock interface{}) interface{} {
	if cond == true {
		return trueBlock
	}
	return falseBlock
}

func Path(path []string, data interface{}) interface{} {
	var result interface{}
	if IsEmpty(data) {
		return nil
	}
	f := DeepFields(0, path, data, &result)
	if Equals(path, f) {
		return result
	}
	return nil
}

func PathOr(defaultResult interface{}, path []string, data interface{}) interface{} {
	var result interface{}
	if IsEmpty(data) {
		return defaultResult
	}
	f := DeepFields(0, path, data, &result)
	if Equals(path, f) {
		return result
	}
	return defaultResult
}

func IsInteger(s string) bool {
	_, err := strconv.Atoi(s)
	if err != nil {
		return false
	}
	return true
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
	case reflect.Slice:
		sindex, err := strconv.Atoi(path[index])
		if err != nil {
			return
		}
		if len(path) <= index+1 && sindex >= 0 {
			if ifv.IsValid() {
				field = append(field, path[index])
			}
			setResult()
			val := reflect.ValueOf(result)
			if result != nil {
				val.Elem().Set(ifv.Index(sindex))
			}
			return
		}

		field = append(field, path[index])
		subField := DeepFields(index+1, path, ifv.Index(sindex).Interface(), result)
		if len(subField) > 0 {
			field = append(field, subField...)
		}
	case reflect.Ptr:
		subField := DeepFields(index, path,
			ifv.Elem().Interface(),
			result)
		if len(subField) > 0 {
			field = append(field, subField...)
		}
	case reflect.Map:
		for _, e := range ifv.MapKeys() {
			if Equals(e.Interface(), path[index]) {
				mv := ifv.MapIndex(e)
				switch mv.Kind() {
				case reflect.Int, reflect.Int32, reflect.Int64, reflect.String, reflect.Float64:
					val := reflect.ValueOf(result)
					if result != nil {
						val.Elem().Set(mv)
						if mv.IsValid() {
							field = append(field, path[index])
						}
					}
				case reflect.Struct:
					if !mv.IsValid() {
						return
					}
					field = append(field, path[index])
					subField := DeepFields(index+1, path, mv.Interface(), result)
					if len(subField) > 0 {
						field = append(field, subField...)
					}
				case reflect.Ptr:
					field = append(field, path[index])
					subField := DeepFields(index+1, path, mv.Interface(), result)
					if len(subField) > 0 {
						field = append(field, subField...)
					}
				default:
					return
				}
			}

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

func PathSatisfies(fn DeciderFun, path []string, data interface{}) bool {
	var result interface{}
	f := DeepFields(0, path, data, &result)
	if Equals(path, f) {
		return fn(result)
	}
	return false
}

func PropSatisfies(fn DeciderFun, path string, d interface{}) bool {
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
	if d == nil {
		return true
	}
	switch reflect.TypeOf(d).Kind() {
	case reflect.Ptr, reflect.Map, reflect.Array, reflect.Chan, reflect.Slice:
		return reflect.ValueOf(d).IsNil()
	}
	return false
}

// IsEmpty will return true if it has its default value
func IsEmpty(d interface{}) bool {
	if IsNil(d) {
		return true
	}
	defaultValue := getDefaultValueOf(d)
	return Equals(d, defaultValue)
}
