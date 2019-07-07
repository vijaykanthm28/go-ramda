package goramda

import (
	"reflect"
)

var (
	defaultBool       bool
	defaultByte       byte
	defaultComplex128 complex128
	defaultComplex64  complex64
	defaultError      error
	defaultFloat32    float32
	defaultFloat64    float64
	defaultInt        int
	defaultInt16      int16
	defaultInt32      int32
	defaultInt64      int64
	defaultInt8       int8
	defaultRune       rune
	defaultString     string
	defaultUint       uint
	defaultUint16     uint16
	defaultUint32     uint32
	defaultUint64     uint64
	defaultUint8      uint8
	defaultUintptr    uintptr
	defaultStruct     struct{}
	defaultInterface  interface{}
)

var comparableKind = []reflect.Kind{
	reflect.String,
	reflect.Int,
	reflect.Int8,
	reflect.Int16,
	reflect.Int32,
	reflect.Int64,
	reflect.Uint,
	reflect.Uint8,
	reflect.Uint16,
	reflect.Uint32,
	reflect.Uint64,
	reflect.Float32,
	reflect.Float64,
	reflect.Uintptr,
}

var (
	defaultKindValues = map[reflect.Kind]interface{}{
		reflect.String:     defaultString,
		reflect.Bool:       defaultBool,
		reflect.Int:        defaultInt,
		reflect.Int8:       defaultInt8,
		reflect.Int16:      defaultInt16,
		reflect.Int32:      defaultInt32,
		reflect.Int64:      defaultInt64,
		reflect.Uint:       defaultUint,
		reflect.Uint8:      defaultUint8,
		reflect.Uint16:     defaultUint16,
		reflect.Uint32:     defaultUint32,
		reflect.Uint64:     defaultUint64,
		reflect.Uintptr:    defaultUintptr,
		reflect.Float32:    defaultFloat32,
		reflect.Float64:    defaultFloat64,
		reflect.Complex64:  defaultComplex64,
		reflect.Complex128: defaultComplex128,
	}
)

/* reflect
Invalid Kind = iota
Bool
Int
Int8
Int16
Int32
Int64
Uint
Uint8
Uint16
Uint32
Uint64
Uintptr
Float32
Float64
Complex64
Complex128
Array
Chan
Func
Interface
Map
Ptr
Slice
String
Struct
UnsafePointer
*/
