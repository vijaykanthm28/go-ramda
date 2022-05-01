package goramda

import (
	"fmt"
	"reflect"
	"testing"
)


type MainStruct struct {
	A   int
	B   string
	Sub *Sub
}

type Sub struct {
	Sa       string
	SuperSub SuperSub
	Number   int
}

type SuperSub struct {
	AJ           float64
	A            string
	SliceSet     []DataSet
	Map          map[string]string
	MapStruct    map[string]DataSet
	MapPtrStruct map[string]*DataSet
}

type DataSet struct {
	Name  string
	Value string
}

var tt = MainStruct{
	A: 23,
	B: "skidoo",
	Sub: &Sub{
		Sa: "Welcome",
		SuperSub: SuperSub{
			AJ: 5.2,
			Map: map[string]string{
				"A": "Testing",
			},
			SliceSet: []DataSet{
				{
					Name:  "FirstName",
					Value: "Vijay",
				},
				{
					Name:  "LastName",
					Value: "Kanth",
				},
				{
					Name:  "MiddleName",
					Value: "P",
				},
			},
			MapStruct: map[string]DataSet{
				"B": {
					Name:  "FirstName",
					Value: "Vijay",
				},
				"C": {
					Name:  "LastName",
					Value: "Kanth",
				},
			},
			MapPtrStruct: map[string]*DataSet{
				"B": {
					Name:  "FirstName",
					Value: "Vijay",
				},
				"C": {
					Name:  "LastName",
					Value: "Kanth",
				},
			},
		},
	},
}


func TestIsEmpty(t *testing.T) {
	data := []struct {
		name           string
		element        interface{}
		expectedResult bool
	}{
		{
			name:           "empty int",
			element:        0,
			expectedResult: true,
		},
		{
			name:           "value int",
			element:        3,
			expectedResult: false,
		},
		{
			name:           "empty string",
			element:        "",
			expectedResult: true,
		},
		{
			name:           "value string",
			element:        "Welcome",
			expectedResult: false,
		},
		{
			name:           "default bool",
			element:        false,
			expectedResult: true,
		},
		{
			name:           "empty struct",
			element:        SuperSub{},
			expectedResult: true,
		},
		{
			name:           "value struct",
			element:        SuperSub{AJ: 5.2},
			expectedResult: false,
		},
		{
			name:           "empty struct main",
			element:        MainStruct{},
			expectedResult: true,
		},
		{
			name:           "empty pointer struct main",
			element:        &MainStruct{},
			expectedResult: true,
		},
		{
			name:           "value pointer struct main",
			element:        MainStruct{A: 23, B: "skidoo"},
			expectedResult: false,
		},
		{
			name:           "empty nil",
			element:        nil,
			expectedResult: true,
		},
	}
	fmt.Println("Total data :", reflect.DeepEqual(&MainStruct{A: 23, B: "skidoo"}, &MainStruct{}))
	for _, d := range data {
		result := IsEmpty(d.element)
		if result != d.expectedResult {
			fmt.Printf("[%s] data failed ", d.name)
			t.Fatalf("Unexpected result expected (%v) for element (%#v) got: %v", d.expectedResult, d.element, result)
		}
	}
}

func TestHasPath(t *testing.T) {
	path := []string{"Sub", "SuperSub", "AJ"}
	if !HasPath(path, &tt) {
		t.Fatal("Unexpected result has path expected true got: false")
	}
}

func TestPathOr(t *testing.T) {
	path := []string{"Subb"}
	expected := &Sub{}
	result := PathOr(expected, path, &tt).(*Sub)
	if expected != result {
		t.Fatalf("Unexpected result pathOr expected (%T) got (%T) ", expected, result)
	}
	expected2 := tt.Sub
	result2 := PathOr(expected2, []string{"Sub"}, &tt).(*Sub)
	if expected2 != result2 {
		t.Fatalf("Unexpected result pathOr expected (%T) got (%T) ", expected2, result2)
	}
}

func TestFirstLevelFieldOnPathOr(t *testing.T) {
	path := []string{"A"}
	expected := 23
	result := PathOr(0, path, &tt).(int)
	if expected != result {
		t.Fatalf("Unexpected result pathOr expected (%T) got (%T) ", expected, result)
	}
}

func TestPathOrSlice(t *testing.T) {
	path := []string{"1"}
	data := []string{
		"Keyboard",
		"Mouse",
		"Touch Pen",
		"Monitor",
	}
	expected := data[1]
	result := PathOr("", path, data).(string)
	if NotEquals(expected, result) {
		t.Fatalf(fmt.Sprintf("Unexpected result pathOr expected (%T/%v) got (%T/%v) ", expected, expected, result, result))
	}
}

func TestPathOrStructSlice(t *testing.T) {
	path := []string{"Sub", "SuperSub", "SliceSet"}
	expected := tt.Sub.SuperSub.SliceSet
	result := PathOr([]DataSet{}, path, &tt).([]DataSet)
	if NotEquals(expected, result) {
		t.Fatalf(fmt.Sprintf("Unexpected result pathOr expected (%T/%v) got (%T/%v) ", expected, expected, result, result))
	}
	path1 := []string{"Sub", "SuperSub", "SliceSet", "1"}
	expected1 := tt.Sub.SuperSub.SliceSet[1]
	result1 := PathOr(DataSet{}, path1, &tt).(DataSet)
	if NotEquals(expected1, result1) {
		t.Fatalf(fmt.Sprintf("Unexpected result pathOr expected (%T/%v) got (%T/%v) ", expected1, expected1, result1, result1))
	}
	path2 := []string{"Sub", "SuperSub", "SliceSet", "2", "Name"}
	expected2 := tt.Sub.SuperSub.SliceSet[2].Name
	result2 := PathOr("default", path2, &tt).(string)
	if NotEquals(expected2, result2) {
		t.Fatalf(fmt.Sprintf("Unexpected result pathOr expected (%T/%v) got (%T/%v) ", expected2, expected2, result2, result2))
	}
}

func TestPathOrMap(t *testing.T) {
	path := []string{"Key"}
	data := []interface{}{
		map[string]string{
			"Key": "Value",
		},
		map[string]int{
			"Key": 10,
		},
		map[string]int32{
			"Key": 10,
		},
		map[string]int64{
			"Key": 10,
		},
		map[string]float64{
			"Key": 10.25,
		},
		map[string]string{},
	}
	expected := []interface{}{
		"Value",
		10,
		int32(10),
		int64(10),
		10.25,
		"default",
	}
	for index, v := range data {
		result := PathOr("default", path, v)
		if expected[index] != result {
			t.Fatalf(fmt.Sprintf("Unexpected result pathOr expected (%T/%v) got (%T/%v) ", expected[index], expected[index], result, result))
		}
	}

}

func TestPathOrMapInstruct(t *testing.T) {
	path := []string{"Sub", "SuperSub", "Map", "A"}
	expected := tt.Sub.SuperSub.Map["A"]
	result, ok := PathOr("default", path, &tt).(string)
	if !ok || expected != result {
		t.Fatalf(fmt.Sprintf("Unexpected result pathOr expected (%T/%v) got (%T/%v) ", expected, expected, result, result))
	}
}

func TestPathOrMapStruct(t *testing.T) {
	path := []string{"Sub", "SuperSub", "MapStruct", "C", "Name"}
	path2 := []string{"Sub", "SuperSub", "MapStruct", "C", "Value"}
	expected := tt.Sub.SuperSub.MapStruct["C"].Name
	expected2 := tt.Sub.SuperSub.MapStruct["C"].Value
	result, ok := PathOr("default", path, &tt).(string)
	if !ok || expected != result {
		t.Fatalf(fmt.Sprintf("Unexpected result pathOr expected (%T/%v) got (%T/%v) ", expected, expected, result, result))
	}
	result2, ok := PathOr(expected2, path2, &tt).(string)
	if !ok || expected2 != result2 {
		t.Fatalf(fmt.Sprintf("Unexpected result pathOr expected (%T/%v) got (%T/%v) ", expected, expected, result, result))
	}
	fmt.Println("Expected:", expected, expected2)
	fmt.Println("Result:", result, result2)
}

func TestPathOrMapPtrStruct(t *testing.T) {
	path := []string{"Sub", "SuperSub", "MapPtrStruct", "C", "Name"}
	path2 := []string{"Sub", "SuperSub", "MapPtrStruct", "C", "Value"}
	expected := tt.Sub.SuperSub.MapPtrStruct["C"].Name
	expected2 := tt.Sub.SuperSub.MapPtrStruct["C"].Value
	result, ok := PathOr("default", path, &tt).(string)
	if !ok || expected != result {
		t.Fatalf(fmt.Sprintf("Unexpected result pathOr expected (%T/%v) got (%T/%v) ", expected, expected, result, result))
	}
	result2, ok := PathOr(expected2, path2, &tt).(string)
	if !ok || expected2 != result2 {
		t.Fatalf(fmt.Sprintf("Unexpected result pathOr expected (%T/%v) got (%T/%v) ", expected, expected, result, result))
	}
	fmt.Println("Expected:", expected, expected2)
	fmt.Println("Result:", result, result2)
}

func TestPathOrWithEmpty(t *testing.T) {
	path := []string{"Sub", "SuperSub", "MapStruct", "C", "Name"}
	path2 := []string{"Sub", "SuperSub", "MapStruct", "C", "Value"}
	result, ok := PathOr("default", path, nil).(string)
	expected := "default"
	if !ok || expected != result {
		t.Fatalf(fmt.Sprintf("[1]-Unexpected result pathOr expected (%T/%v) got (%T/%v) ", expected, expected, result, result))
	}
	result1, ok := PathOr("default", path, new(MainStruct)).(string)
	if !ok || expected != result1 {
		t.Fatalf(fmt.Sprintf("[2]-Unexpected result pathOr expected (%T/%v) got (%T/%v) ", expected, expected, result, result))
	}
	result2, ok := PathOr("default", path2, new(MainStruct)).(string)
	if !ok || expected != result2 {
		t.Fatalf(fmt.Sprintf("[3]-Unexpected result pathOr expected (%T/%v) got (%T/%v) ", expected, expected, result, result))
	}
	fmt.Println("Expected:", expected, expected)
	fmt.Println("Result:", result, result2)
}


func TestPropsSatisfies(t *testing.T) {
	validator := func(d interface{}) bool {
		val, ok := d.(*Sub)
		if ok && val == tt.Sub {
			return true
		}
		return false
	}
	if !PropSatisfies(validator, "Sub", &tt) {
		t.Fatal("Unexpected result ")
	}
}

func TestGetDefaultNewData(t *testing.T) {
	s := &Sub{
		Sa: "Welcome",
	}
	v := getDefaultValueOf(s)
	d, ok := v.(*Sub)
	if !ok {
		t.Fatalf("Unexpected type (%T) expected type (%T)", d, s)
	}

	if d.Sa == s.Sa {
		t.Fatalf("Unexpected object expected empty object got (%v)", d)
	}
}
