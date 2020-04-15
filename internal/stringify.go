package internal

import (
	"reflect"
	"strconv"
)

func Stringify(value interface{}, nonStringPostFix string) string {
	return stringify(reflect.ValueOf(value), nonStringPostFix)
}

//stringify calls recursively until we get the type and value
func stringify(valueOf reflect.Value, nonStringPostFix string) string {
	switch valueOf.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(valueOf.Int(), 10) + nonStringPostFix
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return strconv.FormatUint(valueOf.Uint(), 10) + nonStringPostFix
	case reflect.Float32, reflect.Float64:
		return strconv.FormatFloat(valueOf.Float(), 'f', -1, 64) + nonStringPostFix
	case reflect.String:
		return valueOf.String()
	case reflect.Ptr:
		i := reflect.Indirect(valueOf)
		return stringify(i, nonStringPostFix)

	default:
		return ""
	}
}
