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
	"log"
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
			tmp = fmt.Sprintf("%v%v", tmp, "")
		case error:
			log.Fatalf("Error AddFunc return:  %v", v.Error())
		default:
			switch reflect.TypeOf(v).Kind() {
			case reflect.Slice:
				s := reflect.ValueOf(v)
				for i := 0; i < s.Len(); i++ {
					tmp = fmt.Sprintf("%v%v", tmp, s.Index(i))
				}
			default:
				tmp = fmt.Sprintf("%v%v", tmp, v)

			}

		}

	}
	return tmp

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
		tmp := ""
		for _, val := range str.([]bool) {
			tmp = Build(tmp, val)
		}
		concat = tmp
	case int:
		concat = strconv.Itoa(int(str.(int)))
	case []int:
		concat = IntToStringFast(str.([]int))
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
	case uint8:
		concat = strconv.FormatUint(uint64(str.(uint8)), 10)
	case []int8:
		tmp := ""
		for _, val := range str.([]int8) {
			tmp = Build(tmp, val)
		}
		concat = tmp
	case []uint8:
		tmp := ""
		for _, val := range str.([]uint8) {
			tmp = Build(tmp, val)
		}
		concat = tmp
	case int16:
		concat = strconv.Itoa(int(str.(int16)))
	case uint16:
		concat = strconv.FormatUint(uint64(str.(uint16)), 10)
	case []int16:
		tmp := ""
		for _, val := range str.([]int16) {
			tmp = Build(tmp, val)
		}
		concat = tmp
	case []uint16:
		tmp := ""
		for _, val := range str.([]uint16) {
			tmp = Build(tmp, val)
		}
		concat = tmp
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
		concat = Int32ToStringFast(str.([]int32))
	case []uint32:
		tmp := ""
		for _, val := range str.([]uint32) {
			tmp = Build(tmp, val)
		}
		concat = tmp
	case int64:
		concat = strconv.FormatInt(int64(str.(int64)), 10)
	case uint64:
		concat = strconv.FormatUint(uint64(str.(uint64)), 10)
	case []int64:
		concat = Int64ToStringFast(str.([]int64))
	case []uint64:
		tmp := ""
		for _, val := range str.([]uint64) {
			tmp = Build(tmp, val)
		}
		concat = tmp
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
		concat = strconv.FormatFloat(str.(float64), 'f', 6, 64)
	case []float64:
		tmp := ""
		for _, val := range str.([]float64) {
			tmp = Build(tmp, val)
		}
		concat = tmp
	case float32:
		concat = strconv.FormatFloat(float64(str.(float32)), 'f', 6, 64)
	case []float32:
		tmp := ""
		for _, val := range str.([]float32) {
			tmp = Build(tmp, val)
		}
		concat = tmp
	case uint:
		concat = strconv.FormatUint(uint64(str.(uint)), 10)
	case []uint:
		tmp := ""
		for _, val := range str.([]uint) {
			tmp = Build(tmp, val)
		}
		concat = tmp
	default:
		concat = "not exist type without suport"
	}
	return
}

// IntToStringFast Function that converts []int to string optimally
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

// Int32ToStringFast Function that converts []int32 to string optimally
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

// Int64ToStringFast Function that converts []int64 to string optimally
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
	//sv := make([]string, len(strs))
	var sb strings.Builder
	for _, val := range strs {
		switch val.(type) {
		case string:
			sb.WriteString(string(val.(string)))
			//sv = append(sv, string(val.(string)))
		case int:
			sb.WriteString(strconv.Itoa(int(val.(int))))
			//sv = append(sv, strconv.Itoa(int(val.(int))))
		}
	}
	concat = sb.String()
	//concat = strings.Join(sv, "")
	return
}
