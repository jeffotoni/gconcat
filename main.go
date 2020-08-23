package gconcat

import (
	"strconv"
	"strings"
)

//Concat
func Build(strs ...interface{}) string {
	var sb strings.Builder
	for _, str := range strs {
		sb.WriteString(buildStr(str))
	}
	return sb.String()
}

//buildStr monta a string
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
		return BoolToStringFast(str.([]bool))

	case int:
		return strconv.Itoa(str.(int))
	case []int:
		return IntToStringFast(str.([]int))
	case uint:
		return strconv.FormatUint(uint64(str.(uint)), 10)
	case []uint:
		return IntUToStringFast(str.([]uint))
	case int8:
		return strconv.Itoa(int(str.(int8)))
	//provavelmente funciona para byte, pois é um aliais para uint8 // não testado para isto  @ancogamer
	case uint8:
		return strconv.FormatUint(uint64(str.(uint8)), 10)
	case []int8:
		return Int8ToStringFast(str.([]int8))

	case []uint8:
		return IntU8ToStringFast(str.([]uint8))
	case int16:
		return strconv.Itoa(int(str.(int16)))
	case uint16:
		return strconv.FormatUint(uint64(str.(uint16)), 10)
	case []int16:
		return Int16ToStringFast(str.([]int16))
	case []uint16:
		return IntU16ToStringFast(str.([]uint16))

	//probably work for rune too, since rune is a alias for int32 //Not tested @ancogamer
	case int32:
		return strconv.FormatInt(int64(str.(int32)), 10)
	case uint32:
		return strconv.FormatUint(uint64(str.(uint32)), 10)
	case []int32:
		return Int32ToStringFast(str.([]int32))
	case []uint32:
		return IntU32ToStringFast(str.([]uint32))
	case int64:
		return strconv.FormatInt(int64(str.(int64)), 10)
	case uint64:
		return strconv.FormatUint(uint64(str.(uint64)), 10)
	case []int64:
		return Int64ToStringFast(str.([]int64))
	case []uint64:
		return IntU64ToStringFast(str.([]uint64))
	case float64:
		return strconv.FormatFloat(str.(float64), 'f', 6, 64)
	case []float64:
		return Float64TostringFast(str.([]float64))
	case float32:
		return strconv.FormatFloat(float64(str.(float32)), 'f', 6, 64)
	case []float32:
		return Float32TostringFast(str.([]float32))

	// case uintptr:
	// 	return string(str.(uintptr))
	// case []uintptr:
	// 	concat := ""
	// 	for _, val := range str.([]uintptr) {
	// 		concat = Build(concat, val)
	// 	}
	// 	return concat

	case complex64:
		return "not suport complex 64, maybe in the future"
	case complex128:
		return "not suport complex 128, maybe in the future"
	default:
		println("type not found")
		break
	}
	return ""
}
func BoolToStringFast(a []bool) string {
	if len(a) == 0 {
		return ""
	}
	b := make([]string, len(a))
	for i, v := range a {
		b[i] = strconv.FormatBool(v)
	}
	return strings.Join(b, "")
}
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

func Int8ToStringFast(a []int8) string {
	if len(a) == 0 {
		return ""
	}
	b := make([]string, len(a))
	for i, v := range a {
		b[i] = strconv.Itoa(int(v))
	}
	return strings.Join(b, "")
}

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
func Int16ToStringFast(a []int16) string {
	if len(a) == 0 {
		return ""
	}
	b := make([]string, len(a))
	for i, v := range a {
		b[i] = strconv.FormatInt(int64(v), 10)
	}
	return strings.Join(b, "")
}

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

func IntUToStringFast(a []uint) string {
	if len(a) == 0 {
		return ""
	}
	b := make([]string, len(a))
	for i, v := range a {
		b[i] = strconv.FormatUint(uint64(v), 10)
	}
	return strings.Join(b, "")
}
func IntU8ToStringFast(a []uint8) string {
	if len(a) == 0 {
		return ""
	}
	b := make([]string, len(a))
	for i, v := range a {
		b[i] = strconv.FormatUint(uint64(v), 10)
	}
	return strings.Join(b, "")
}
func IntU16ToStringFast(a []uint16) string {
	if len(a) == 0 {
		return ""
	}
	b := make([]string, len(a))
	for i, v := range a {
		b[i] = strconv.FormatUint(uint64(v), 10)
	}
	return strings.Join(b, "")
}

func IntU32ToStringFast(a []uint32) string {
	if len(a) == 0 {
		return ""
	}
	b := make([]string, len(a))
	for i, v := range a {
		b[i] = strconv.FormatUint(uint64(v), 10)
	}
	return strings.Join(b, "")
}
func IntU64ToStringFast(a []uint64) string {
	if len(a) == 0 {
		return ""
	}
	b := make([]string, len(a))
	for i, v := range a {
		b[i] = strconv.FormatUint(v, 10)
	}
	return strings.Join(b, "")
}

func Float32TostringFast(a []float32) string {
	if len(a) == 0 {
		return ""
	}
	b := make([]string, len(a))
	for i, v := range a {
		b[i] = strconv.FormatFloat(float64(v), 'f', 6, 64)
	}
	return strings.Join(b, "")
}
func Float64TostringFast(a []float64) string {
	if len(a) == 0 {
		return ""
	}
	b := make([]string, len(a))
	for i, v := range a {
		b[i] = strconv.FormatFloat(v, 'f', 6, 64)
	}
	return strings.Join(b, "")
}
