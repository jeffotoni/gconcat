package gconcat

import (
	"strconv"
	"strings"
)

// Function responsible for abstracting the build function and delivering a Concat
func Concat(str ...interface{}) string {
	return Build(str...)
}

// Function responsible for concatenating, and accepting different types
func Build(strs ...interface{}) string {
	var sb strings.Builder
	for _, str := range strs {
		sb.WriteString(buildStr(str))
	}
	return sb.String()
}

func buildStr(str interface{}) string {
	switch str.(type) {
	case nil:
		return "nil"
	case string:
		return string(str.(string))

	case []string:
		return strings.Join(str.([]string), "")

	case []interface{}:
		concat := ""
		for _, val := range str.([]interface{}) {
			concat = Build(concat, val)
		}
		return concat

	case bool:
		return strconv.FormatBool(str.(bool))
	case []bool:
		concat := ""
		for _, val := range str.([]bool) {
			concat = Build(concat, val)
		}
		return concat
	//
	case int:
		return strconv.Itoa(int(str.(int)))

	case []int:
		return IntToStringFast(str.([]int))
	case uint:
		return strconv.FormatUint(uint64(str.(uint)), 10)
	case []uint:
		concat := ""
		for _, val := range str.([]uint) {
			concat = Build(concat, val)
		}
		return concat
	case int8:
		return strconv.Itoa(int(str.(int8)))
	case uint8:
		return strconv.FormatUint(uint64(str.(uint8)), 10)
	case []int8:
		concat := ""
		for _, val := range str.([]int8) {
			concat = Build(concat, val)
		}
		return concat

	case []uint8:
		concat := ""
		for _, val := range str.([]uint8) {
			concat = Build(concat, val)
		}
		return concat

	case int16:
		return strconv.Itoa(int(str.(int16)))
	case uint16:
		return strconv.FormatUint(uint64(str.(uint16)), 10)
	case []int16:
		concat := ""
		for _, val := range str.([]int16) {
			concat = Build(concat, val)
		}
		return concat

	case []uint16:
		concat := ""
		for _, val := range str.([]uint16) {
			concat = Build(concat, val)
		}
		return concat
	case int32:
		return strconv.FormatInt(int64(str.(int32)), 10)
	case uint32:
		return strconv.FormatUint(uint64(str.(uint32)), 10)
	case []int32:
		return Int32ToStringFast(str.([]int32))
	case []uint32:
		concat := ""
		for _, val := range str.([]uint32) {
			concat = Build(concat, val)
		}
		return concat
	case int64:
		return strconv.FormatInt(int64(str.(int64)), 10)
	case uint64:
		return strconv.FormatUint(uint64(str.(uint64)), 10)
	case []int64:
		return Int64ToStringFast(str.([]int64))
	case []uint64:
		concat := ""
		for _, val := range str.([]uint64) {
			concat = Build(concat, val)
		}
		return concat

	case float64:
		return strconv.FormatFloat(str.(float64), 'f', 6, 64)
	case []float64:
		concat := ""
		for _, val := range str.([]float64) {
			concat = Build(concat, val)
		}
		return concat
	case float32:
		return strconv.FormatFloat(float64(str.(float32)), 'f', 6, 64)
	case []float32:
		concat := ""
		for _, val := range str.([]float32) {
			concat = Build(concat, val)
		}
		return concat
	case complex64:
		return "not suport complex 64"
	case complex128:
		return "not suport complex 128"
	default:
		return "not exist type without suport"
		break
	}

	return ""
}

// Function that converts []int to string optimally
func IntToStringFast(a []int) string {
	if len(a) == 0 {
		return ""
	}
	b := make([]string, len(a))
	for i, v := range a {
		b[i] = strconv.Itoa(v)
	}
	return strings.Join(b, "")
}

// Function that converts []int32 to string optimally
func Int32ToStringFast(a []int32) string {
	if len(a) == 0 {
		return ""
	}
	b := make([]string, len(a))
	for i, v := range a {
		b[i] = strconv.FormatInt(int64(v), 10)
	}
	return strings.Join(b, "")
}

// Function that converts []int64 to string optimally
func Int64ToStringFast(a []int64) string {
	if len(a) == 0 {
		return ""
	}
	b := make([]string, len(a))
	for i, v := range a {
		b[i] = strconv.FormatInt(int64(v), 10)
	}
	return strings.Join(b, "")
}
