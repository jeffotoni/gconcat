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
// The result will be  Concat(int, []int, []int32, []string, string, bool, reflect.Value(string, bool, int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64))
func Concat(str ...interface{}) string {
	return Build(str...)
}

// ConcatFuncOLD is deprecated.
func ConcatFuncOLD(f ...interface{}) string {
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

// ConcatFunc will be responsible for concatenating the function's return into a string.
// nil return parameters are ignored.
// parameters with valid errors are still triggered and must be corrected.
func ConcatFunc(f ...interface{}) string {
	var tmpS strings.Builder
	for _, val := range f {
		switch v := val.(type) {
		case nil, func(), struct{}:
		case error:
			tmpS.WriteString(v.Error())
		default: // TODO REMOVE WHEN buildStr1 to buildStr4 be capable of not use reflect package
			tmpS.WriteString(buildStr5(reflect.ValueOf(v)))
		}
	}
	return tmpS.String()
}

// Build Function responsible for concatenating, and accepting different types
// The result will be BuildConcat(int, []int, []int32, []string, string, bool, reflect.Value(string, bool, int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64))
func Build(strs ...interface{}) string {
	// testing various ways to improve
	// performance is using Builder
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
		concat = ConcatSliceBool(str.([]bool))
	case int:
		concat = strconv.Itoa(int(str.(int)))
	case []int:
		concat = ConcatSliceInt(str.([]int))
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
		concat = ConcatSliceInt8(str.([]int8))
	case uint8:
		concat = strconv.FormatUint(uint64(str.(uint8)), 10)
	case []uint8:
		concat = ConcatSliceUint8(str.([]uint8))
	case int16:
		concat = strconv.Itoa(int(str.(int16)))
	case []int16:
		concat = ConcatSliceInt16(str.([]int16))
	case uint16:
		concat = strconv.FormatUint(uint64(str.(uint16)), 10)
	case []uint16:
		concat = ConcatSliceUint16(str.([]uint16))
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
		concat = ConcatSliceInt32(str.([]int32))
	case []uint32:
		concat = ConcatSliceUint32(str.([]uint32))
	case int64:
		concat = strconv.FormatInt(int64(str.(int64)), 10)
	case uint64:
		concat = strconv.FormatUint(uint64(str.(uint64)), 10)
	case []int64:
		concat = ConcatSliceInt64(str.([]int64))
	case []uint64:
		concat = ConcatSliceUint64(str.([]uint64))
	default:
		concat = buildStr4(str)
	}
	return
}

// buildStr4 Function responsible abstracting and case some types
// buildStr4(float64, []float64, float32, []float32, uint, []uint)
func buildStr4(str interface{}) (concat string) {
	switch str.(type) {
	case float64:
		concat = strconv.FormatFloat(str.(float64), 'f', -1, 64)
	case []float64:
		concat = ConcatSliceFloat64(str.([]float64))
	case float32:
		concat = strconv.FormatFloat(float64(str.(float32)), 'f', -1, 64)
	case []float32:
		concat = ConcatSliceFloat32(str.([]float32))
	case uint:
		concat = strconv.FormatUint(uint64(str.(uint)), 10)
	case []uint:
		concat = ConcatSliceUint(str.([]uint))
	default:
		concat = buildStr5(str)
	}
	return
}

// buildStr5 Function responsible abstracting cases where interface is reflect.Value
// buildStr5( string, struct EXPORTED FIELDS VALUES AND EXPORTED METHODS, digest result of func() function with paramters are not supported,
// bool, slice,array, int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64)
//
// For the struct methods, if the struct has exported method, it will first concat the values inside the struct. ex Struct{A:10} with method TurnTo20( return 20), will be concat 1020
func buildStr5(str interface{}) (concat string) {
	switch str.(type) {
	case reflect.Value:
		v := str.(reflect.Value)
		k := v.Kind()
		fmt.Println(k)
		if k == reflect.Bool {
			if v.Interface().(bool) {
				concat = "true"
				return
			}
			concat = "false"
			return
		}
		if k == reflect.Int {
			concat = strconv.FormatInt(int64(v.Interface().(int)), 10)
			return
		}
		if k == reflect.Int8 {
			concat = strconv.FormatInt(int64(v.Interface().(int8)), 10)
			return
		}
		if k == reflect.Int16 {
			concat = strconv.FormatInt(int64(v.Interface().(int16)), 10)
			return
		}
		if k == reflect.Int32 {
			concat = strconv.FormatInt(int64(v.Interface().(int32)), 10)
			return
		}
		if k == reflect.Int64 {
			concat = strconv.FormatInt(v.Interface().(int64), 10)
			return
		}
		if k == reflect.Uint {
			concat = strconv.FormatUint(uint64(v.Interface().(uint)), 10)
			return
		}
		if k == reflect.Uint8 {
			concat = strconv.FormatUint(uint64(v.Interface().(uint8)), 10)
			return
		}
		if k == reflect.Uint16 {
			concat = strconv.FormatUint(uint64(v.Interface().(uint16)), 10)
			return
		}
		if k == reflect.Uint32 {
			concat = strconv.FormatUint(uint64(v.Interface().(uint32)), 10)
			return
		}
		if k == reflect.Uint64 {
			concat = strconv.FormatUint(v.Interface().(uint64), 10)
			return
		}
		if k == reflect.Float32 {
			concat = strconv.FormatFloat(float64(v.Interface().(float32)), 'f', -1, 64)
			return
		}
		if k == reflect.Float64 {
			concat = strconv.FormatFloat(v.Interface().(float64), 'f', -1, 64)
			return
		}
		if k == reflect.String {
			concat = v.Interface().(string)
			return
		}
		if k == reflect.Slice {
			var tmp strings.Builder
			for i := 0; i < v.Len(); i++ {
				tmp.WriteString(Build(v.Index(i)))
			}
			concat = tmp.String()
			return
		}
		if k == reflect.Array {
			var tmp strings.Builder
			for i := 0; i < v.Len(); i++ {
				tmp.WriteString(Build(v.Index(i)))
			}
			concat = tmp.String()
			return
		}
		if k == reflect.Func { // calling back buildStr5 to run over reflect types
			var tmp strings.Builder
			if v.IsNil() {
				return
			}
			s := v.Call([]reflect.Value{})
			for i := 0; i < len(s); i++ {
				tmp.WriteString(buildStr5(s[i]))
			}
			concat = tmp.String()
			return
		}
		if k == reflect.Struct {
			var tmp strings.Builder
			exF := reflect.VisibleFields(v.Type())
			for i := 0; i < len(exF); i++ {
				if exF[i].IsExported() {
					tmp.WriteString(buildStr5(v.FieldByName(exF[i].Name)))
				}
			}

			concat = tmp.String()
			return
		}

		// TODO REMOVE when new handling of reflect type arrives
		concat = fmt.Sprint(v)
		return
		/* TODO

		if k == reflect.Interface {
			fmt.Println("to aqui")
			v.Interface()
		}

		if k == reflect.Map {
		}
		if k == reflect.Pointer {
		}

		*/
	default: // case the reflect type isn't already reflect.Value, we reflect the type and call the function again.
		concat = buildStr5(reflect.ValueOf(str))
		if concat == "" {
			break
		}
		return
	}
	concat = "<not exist type without suport>"
	return
}

// ConcatString Function that converts only string
// ConcatString(strs ...string)
func ConcatString(strs ...string) string {
	return ConcatStr(strs...)
}

// ConcatStringInt Function that converts only string
// ConcatStringInt(strs ...interface{})
func ConcatStringInt(strs ...interface{}) string {
	return ConcatStrInt(strs...)
}

// ConcatStr Function that converts only string
// ConcatStr(strs ...string)
func ConcatStr(strs ...string) string {
	if len(strs) == 0 {
		return ""
	}
	return strings.Join(strs, "")
}

func ConcatStrCopy(strs ...string) string {
	if len(strs) == 0 {
		return ""
	}

	i := 0
	for _, val := range strs {
		i += len(val)
	}

	bs := make([]byte, i, i)
	bl := 0
	for _, val := range strs {
		bl += copy(bs[bl:], val)
	}
	return string(bs)
}
func CopyStr(strs ...string) string {
	return ConcatStrCopy(strs...)
}

// ConcatStrInt Function that converts only string and Int
// ConcatStrInt(strs ...interface{})
func ConcatStrInt(strs ...interface{}) string {
	if len(strs) == 0 {
		return ""
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
	return sb.String()
}

// ConcatSliceBool Function that converts []bool to string optimally
// ConcatSliceBool(a []bool)
func ConcatSliceBool(a []bool) string {
	if len(a) == 0 {
		return ""
	}
	b := make([]string, len(a))
	for i, v := range a {
		b[i] = strconv.FormatBool(v)
	}
	return strings.Join(b, "")
}

// ConcatSliceInt Function that converts []int to string optimally
// ConcatSliceInt(a []int)
func ConcatSliceInt(a []int) string {
	if len(a) == 0 {
		return ""
	}
	b := make([]string, len(a))
	for i, v := range a {
		b[i] = strconv.Itoa(v)
	}
	return strings.Join(b, "")
}

// ConcatSliceInt8 Function that converts []int8 to string optimally
// ConcatSliceInt8(a []int8)
func ConcatSliceInt8(a []int8) string {
	if len(a) == 0 {
		return ""
	}
	b := make([]string, len(a))
	for i, v := range a {
		b[i] = strconv.Itoa(int(v))
	}
	return strings.Join(b, "")
}

// ConcatSliceInt16 Function that converts []int16 to string optimally
// ConcatSliceInt16(a []int16)
func ConcatSliceInt16(a []int16) string {
	if len(a) == 0 {
		return ""
	}
	b := make([]string, len(a))
	for i, v := range a {
		b[i] = strconv.Itoa(int(v))
	}
	return strings.Join(b, "")
}

// ConcatSliceInt32 Function that converts []int32 to string optimally
// ConcatSliceInt32(a []int32)
func ConcatSliceInt32(a []int32) string {
	if len(a) == 0 {
		return ""
	}
	b := make([]string, len(a))
	for i, v := range a {
		b[i] = strconv.FormatInt(int64(v), 10)
	}
	return strings.Join(b, "")
}

// ConcatSliceInt64 Function that converts []int64 to string optimally
// ConcatSliceInt64(a []int64)
func ConcatSliceInt64(a []int64) string {
	if len(a) == 0 {
		return ""
	}
	b := make([]string, len(a))
	for i, v := range a {
		b[i] = strconv.FormatInt(int64(v), 10)
	}
	return strings.Join(b, "")
}

// ConcatSliceUint Function that converts []uint to string optimally
// ConcatSliceUint(a []uint)
func ConcatSliceUint(a []uint) string {
	if len(a) == 0 {
		return ""
	}
	b := make([]string, len(a))
	for i, v := range a {
		b[i] = strconv.FormatUint(uint64(v), 10)
	}
	return strings.Join(b, "")
}

// ConcatSliceUint8 Function that converts []uint8 to string optimally
// ConcatSliceUint8(a []uint8)
func ConcatSliceUint8(a []uint8) string {
	if len(a) == 0 {
		return ""
	}
	b := make([]string, len(a))
	for i, v := range a {
		b[i] = strconv.FormatUint(uint64(v), 10)
	}
	return strings.Join(b, "")
}

// ConcatSliceUint16 Function that converts []uint16 to string optimally
// ConcatSliceUint16(a []uint16)
func ConcatSliceUint16(a []uint16) string {
	if len(a) == 0 {
		return ""
	}
	b := make([]string, len(a))
	for i, v := range a {
		b[i] = strconv.FormatUint(uint64(v), 10)
	}
	return strings.Join(b, "")
}

// ConcatSliceUint32 Function that converts []uint32 to string optimally
// ConcatSliceUint32(a []uint32)
func ConcatSliceUint32(a []uint32) string {
	if len(a) == 0 {
		return ""
	}
	b := make([]string, len(a))
	for i, v := range a {
		b[i] = strconv.FormatUint(uint64(v), 10)
	}
	return strings.Join(b, "")
}

// ConcatSliceUint64 Function that converts []uint64 to string optimally
// ConcatSliceUint64(a []uint64)
func ConcatSliceUint64(a []uint64) string {
	if len(a) == 0 {
		return ""
	}
	b := make([]string, len(a))
	for i, v := range a {
		b[i] = strconv.FormatUint(uint64(v), 10)
	}
	return strings.Join(b, "")
}

// ConcatSliceFloat32 Function that converts []float32 to string optimally
// ConcatSliceFloat32(a []float32)
func ConcatSliceFloat32(a []float32) string {
	if len(a) == 0 {
		return ""
	}
	b := make([]string, len(a))
	for i, v := range a {
		b[i] = strconv.FormatFloat(float64(v), 'f', 6, 64)
	}
	return strings.Join(b, "")
}

// ConcatSliceFloat64 Function that converts []float64 to string optimally
// ConcatSliceFloat64(a []float64)
func ConcatSliceFloat64(a []float64) string {
	if len(a) == 0 {
		return ""
	}
	b := make([]string, len(a))
	for i, v := range a {
		b[i] = strconv.FormatFloat(v, 'f', 6, 64)
	}
	return strings.Join(b, "")
}
