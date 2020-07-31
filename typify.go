package goramda

func String(v interface{}) string {
	if value, ok := v.(string); ok {
		return value
	}
	return ""
}

func Integer64(v interface{}) int64 {
	if value, ok := v.(int64); ok {
		return value
	}
	return int64(0)
}

func Integer32(v interface{}) int32 {
	if value, ok := v.(int32); ok {
		return value
	}
	return int32(0)
}

func Integer(v interface{}) int {
	if value, ok := v.(int); ok {
		return value
	}
	return int(0)
}

func IntSlice(v interface{}) []int {
	if value, ok := v.([]int); ok {
		return value
	}
	return []int{}
}

func StringSlice(v interface{}) []string {
	if value, ok := v.([]string); ok {
		return value
	}
	return []string{}
}

/* Typify converts v into default type t if it is nil
   Typify(nil, 5) => returns 0
      a := ClassA{name: "john"}
   Typify(nil, a) => returns empty ClassA object
   Typify(a, ClassA{}) => returns a

   This function is to avoid condition checkings or panic errors of ramda method
   that return nil values and replace nil into default value
   that provided of default_type value
*/
func Typify(v interface{}, default_type interface{}) interface{} {
	defaultType := getDefaultValueOf(default_type)
	defaultValue := getDefaultValueOf(v)
	if IsNil(v) || !Equals(defaultType, defaultValue) {
		return defaultType
	}
	return v
}
