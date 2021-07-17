// Package gconcat ccepts different types of data.
// Very flexible and easy to use. In our day-to-day lives, we often encounter
// the need for concatenations throughout our project. This lib was born to
// make our daily lives a little more flexible when dealing with concatenations.
//
// Notice that this doc is written in godoc itself as package documentation.
// The defined types are just for making the table of contents at the
// head of the page; they have no meanings as types.
//
// If you have any suggestion or comment, please feel free to open an issue on
// this tutorial's GitHub page!
//
// By @jeffotoni
package gconcat

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

// Concat Function responsible for abstracting the build function and delivering
// The result will be  Concat(int, []int, []int32, []string, string)
func Concat(str ...interface{}) string {
	return Build(str...)
}

//ConcatFunc will be responsible for concatenating the function's return into a string.
//nil return parameters are ignored.
//parameters with valid errors are still triggered and must be corrected.
func ConcatFunc(f ...interface{}) string {
	var tmp string

	for _, val := range f {
		switch v := val.(type) {
		case nil, func(), struct{}:
			tmp = fmt.Sprint(tmp, "")
		case error:
			tmp = fmt.Sprint(tmp, v)
		default:
			switch reflect.TypeOf(v).Kind() {
			case reflect.Slice:
				s := reflect.ValueOf(v)
				for i := 0; i < s.Len(); i++ {
					tmp = fmt.Sprint(tmp, s.Index(i))
				}
			default:
				tmp = fmt.Sprint(tmp, v)
			}
		}
	}
	return string(tmp)
}

// Build Function responsible for concatenating, and accepting different types
// The result will be Build(int, []int, []int32, []string, string)
func Build(strs ...interface{}) string {
	// testing various ways to improve
	// performance is using Builder
	//
	// sv := make([]string, len(strs))
	// for _, str := range strs {
	// 	sv = append(sv, buildStr(str))
	// }
	// return strings.Join(sv, "")

	var sb strings.Builder
	for _, str := range strs {
		sb.WriteString(buildStr(str))
	}
	return sb.String()
}

// buildStr Function responsible abstracting and case some types
// buildStr(nil, string, []string, []interface{}, bool, []bool, int, []int)
func buildStr(str interface{}) (concat string) {
	switch str.(type) {
	case nil:
		concat = "nil"
	case string:
		concat = string(str.(string))
	case []string:
		concat = strings.Join(str.([]string), "")
	case []interface{}:
		tmp := ""
		for _, val := range str.([]interface{}) {
			tmp = Build(tmp, val)
		}
		concat = tmp
	case bool:
		concat = strconv.FormatBool(str.(bool))
	case []bool:
		concat = ConcatBool(str.([]bool))
	case int:
		concat = strconv.Itoa(int(str.(int)))
	case []int:
		concat = ConcatInt(str.([]int))
	case complex64:
		concat = "not suport complex 64"
	case complex128:
		concat = "not suport complex 128"
	default:
		concat = buildStr2(str)
	}
	return
}

// buildStr2 Function responsible abstracting and case some types
// buildStr2(int8, uint8, []int8, []uint8, int16, uint16, []int16, []uint16)
func buildStr2(str interface{}) (concat string) {
	switch str.(type) {
	case int8:
		concat = strconv.Itoa(int(str.(int8)))
	case []int8:
		concat = ConcatInt8(str.([]int8))
	case uint8:
		concat = strconv.FormatUint(uint64(str.(uint8)), 10)
	case []uint8:
		concat = ConcatUint8(str.([]uint8))
	case int16:
		concat = strconv.Itoa(int(str.(int16)))
	case []int16:
		concat = ConcatInt16(str.([]int16))
	case uint16:
		concat = strconv.FormatUint(uint64(str.(uint16)), 10)
	case []uint16:
		concat = ConcatUint16(str.([]uint16))
	default:
		concat = buildStr3(str)
	}
	return
}

// buildStr3 Function responsible abstracting and case some types
// buildStr3(int32, uint32, []int32, []uint32, int64, uint64, []int64, []uint64)
func buildStr3(str interface{}) (concat string) {
	switch str.(type) {
	case int32:
		concat = strconv.FormatInt(int64(str.(int32)), 10)
	case uint32:
		concat = strconv.FormatUint(uint64(str.(uint32)), 10)
	case []int32:
		concat = ConcatInt32(str.([]int32))
	case []uint32:
		concat = ConcatUint32(str.([]uint32))
	case int64:
		concat = strconv.FormatInt(int64(str.(int64)), 10)
	case uint64:
		concat = strconv.FormatUint(uint64(str.(uint64)), 10)
	case []int64:
		concat = ConcatInt64(str.([]int64))
	case []uint64:
		concat = ConcatUint64(str.([]uint64))
	default:
		concat = buildStr4(str)
	}
	return
}

// buildStr4 Function responsible abstracting and case some types
// buildStr4(float64, []float64, float32, []float32, uint, []uint)
func buildStr4(str interface{}) (concat string) {
	switch x := str.(type) {
	case float64:
		concat = strconv.FormatFloat(str.(float64), 'f', 6, 64)
	case []float64:
		concat = ConcatFloat64(str.([]float64))
	case float32:
		concat = strconv.FormatFloat(float64(str.(float32)), 'f', 6, 64)
	case []float32:
		concat = ConcatFloat32(str.([]float32))
	case uint:
		concat = strconv.FormatUint(uint64(str.(uint)), 10)
	case []uint:
		concat = ConcatUint(str.([]uint))
	default:
		fmt.Println(x)
		concat = "not exist type without suport"
	}
	return
}

// ConcatStr Function that converts only string
func ConcatStr(strs ...string) (concat string) {
	if len(strs) == 0 {
		return
	}
	concat = strings.Join(strs, "")
	return
}

// ConcatStrInt Function that converts only string and Int
func ConcatStrInt(strs ...interface{}) (concat string) {
	if len(strs) == 0 {
		return
	}
	var sb strings.Builder
	for _, val := range strs {
		switch val.(type) {
		case string:
			sb.WriteString(string(val.(string)))
		case int:
			sb.WriteString(strconv.Itoa(int(val.(int))))
		}
	}
	concat = sb.String()
	return
}

// ConcatBool Function that converts []bool to string optimally
func ConcatBool(a []bool) string {
	if len(a) == 0 {
		return ""
	}
	b := make([]string, len(a))
	for i, v := range a {
		b[i] = strconv.FormatBool(v)
	}
	return strings.Join(b, "")
}

// ConcatInt Function that converts []int to string optimally
func ConcatInt(a []int) string {
	if len(a) == 0 {
		return ""
	}
	b := make([]string, len(a))
	for i, v := range a {
		b[i] = strconv.Itoa(v)
	}
	return strings.Join(b, "")
}

// ConcatIntFast Function that converts []int to string optimally
func ConcatIntFast(a []int) (tmp string) {
	for _, v := range a {
		tmp = fmt.Sprint(tmp, v)
	}
	return
}

// ConcatInt8 Function that converts []int8 to string optimally
func ConcatInt8(a []int8) string {
	if len(a) == 0 {
		return ""
	}
	b := make([]string, len(a))
	for i, v := range a {
		b[i] = strconv.Itoa(int(v))
	}
	return strings.Join(b, "")
}

// ConcatInt16 Function that converts []int16 to string optimally
func ConcatInt16(a []int16) string {
	if len(a) == 0 {
		return ""
	}
	b := make([]string, len(a))
	for i, v := range a {
		b[i] = strconv.Itoa(int(v))
	}
	return strings.Join(b, "")
}

// ConcatInt32 Function that converts []int32 to string optimally
func ConcatInt32(a []int32) string {
	if len(a) == 0 {
		return ""
	}
	b := make([]string, len(a))
	for i, v := range a {
		b[i] = strconv.FormatInt(int64(v), 10)
	}
	return strings.Join(b, "")
}

// ConcatInt64 Function that converts []int64 to string optimally
func ConcatInt64(a []int64) string {
	if len(a) == 0 {
		return ""
	}
	b := make([]string, len(a))
	for i, v := range a {
		b[i] = strconv.FormatInt(int64(v), 10)
	}
	return strings.Join(b, "")
}

// ConcatUint Function that converts []uint to string optimally
func ConcatUint(a []uint) string {
	if len(a) == 0 {
		return ""
	}
	b := make([]string, len(a))
	for i, v := range a {
		b[i] = strconv.FormatUint(uint64(v), 10)
	}
	return strings.Join(b, "")
}

// ConcatUint8 Function that converts []uint8 to string optimally
func ConcatUint8(a []uint8) string {
	if len(a) == 0 {
		return ""
	}
	b := make([]string, len(a))
	for i, v := range a {
		b[i] = strconv.FormatUint(uint64(v), 10)
	}
	return strings.Join(b, "")
}

// ConcatUint16 Function that converts []uint16 to string optimally
func ConcatUint16(a []uint16) string {
	if len(a) == 0 {
		return ""
	}
	b := make([]string, len(a))
	for i, v := range a {
		b[i] = strconv.FormatUint(uint64(v), 10)
	}
	return strings.Join(b, "")
}

// ConcatUint32 Function that converts []uint32 to string optimally
func ConcatUint32(a []uint32) string {
	if len(a) == 0 {
		return ""
	}
	b := make([]string, len(a))
	for i, v := range a {
		b[i] = strconv.FormatUint(uint64(v), 10)
	}
	return strings.Join(b, "")
}

// ConcatUint64 Function that converts []uint64 to string optimally
func ConcatUint64(a []uint64) string {
	if len(a) == 0 {
		return ""
	}
	b := make([]string, len(a))
	for i, v := range a {
		b[i] = strconv.FormatUint(uint64(v), 10)
	}
	return strings.Join(b, "")
}

// ConcatFloat32 Function that converts []float32 to string optimally
func ConcatFloat32(a []float32) string {
	if len(a) == 0 {
		return ""
	}
	b := make([]string, len(a))
	for i, v := range a {
		b[i] = strconv.FormatFloat(float64(v), 'f', 6, 64)
	}
	return strings.Join(b, "")
}

// ConcatFloat64 Function that converts []float64 to string optimally
func ConcatFloat64(a []float64) string {
	if len(a) == 0 {
		return ""
	}
	b := make([]string, len(a))
	for i, v := range a {
		b[i] = strconv.FormatFloat(v, 'f', 6, 64)
	}
	return strings.Join(b, "")
}
