//Module responsible for handling the verification of
// use cases regarding the various types of language

package gconcat

import (
	"fmt"
	"strconv"
)

//Formating of struct base.
type Formating struct{}

//NewFormat will return point of Formating struct.
func NewFormat() *Formating {
	return &Formating{}
}

//Format will be responsible for handling the various types and returning their values.
func (f *Formating) Format(tmp string, val interface{}) string {
	tmp = formatString(tmp, val)
	tmp = formatInts(tmp, val)
	tmp = formatUints(tmp, val)
	tmp = formatFloats(tmp, val)
	tmp = formatBools(tmp, val)
	return tmp
}

//format cases of strings
func formatString(tmp string, val interface{}) string {

	switch v := val.(type) {
	case string:
		tmp = fmt.Sprint(tmp, v)
	case []string:
		for _, vs := range v {
			tmp = fmt.Sprint(tmp, vs)
		}
	default:
		tmp = fmt.Sprint(tmp, "")
	}
	return tmp
}

//format cases of ints
func formatInts(tmp string, val interface{}) string {
	switch v := val.(type) {
	case int:
		tmp = fmt.Sprint(tmp, strconv.Itoa(v))
	case int8:
		var i8 = v
		i64 := int64(i8)
		tmp = fmt.Sprint(tmp, strconv.FormatInt(i64, 10))
	case int16:
		var i16 = v
		i64 := int64(i16)
		tmp = fmt.Sprint(tmp, strconv.FormatInt(i64, 10))
	case int32:
		var i32 = v
		i64 := int64(i32)
		tmp = fmt.Sprint(tmp, strconv.FormatInt(i64, 10))
	case int64:
		i64 := v
		tmp = fmt.Sprint(tmp, strconv.FormatInt(i64, 10))
	case []int:
		for _, vi := range v {
			tmp = fmt.Sprint(tmp, strconv.Itoa(vi))
		}
	case []int8:
		for _, vi := range v {
			var si8 = vi
			s64 := int64(si8)
			tmp = fmt.Sprint(tmp, strconv.FormatInt(s64, 10))
		}
	case []int16:
		for _, vi := range v {
			var si16 = vi
			s64 := int64(si16)
			tmp = fmt.Sprint(tmp, strconv.FormatInt(s64, 10))
		}
	case []int32:
		for _, vi := range v {
			var si32 = vi
			s64 := int64(si32)
			tmp = fmt.Sprint(tmp, strconv.FormatInt(s64, 10))
		}
	case []int64:
		for _, vi := range v {
			tmp = fmt.Sprint(tmp, strconv.FormatInt(vi, 10))
		}

	default:
		tmp = fmt.Sprint(tmp, "")
	}
	return tmp
}

//format cases uint
func formatUints(tmp string, val interface{}) string {
	switch v := val.(type) {
	case uint:
		var ui = v
		i := int(ui)
		tmp = fmt.Sprint(tmp, strconv.Itoa(i))
	case uint8:
		var ui8 = v
		i64 := int64(ui8)
		tmp = fmt.Sprint(tmp, strconv.FormatInt(i64, 10))
	case uint16:
		var ui16 = v
		i64 := int64(ui16)
		tmp = fmt.Sprint(tmp, strconv.FormatInt(i64, 10))
	case uint32:
		var ui32 = v
		i64 := int64(ui32)
		tmp = fmt.Sprint(tmp, strconv.FormatInt(i64, 10))
	case uint64:
		var ui64 = v
		i64 := int64(ui64)
		tmp = fmt.Sprint(tmp, strconv.FormatInt(i64, 10))
	case []uint:
		for _, vi := range v {
			var ui = vi
			i := int(ui)
			tmp = fmt.Sprint(tmp, strconv.Itoa(i))
		}
	case []uint8:
		for _, vi := range v {
			var ui8 = vi
			i64 := int64(ui8)
			tmp = fmt.Sprint(tmp, strconv.FormatInt(i64, 10))
		}
	case []uint16:
		for _, vi := range v {
			var ui16 = vi
			i64 := int64(ui16)
			tmp = fmt.Sprint(tmp, strconv.FormatInt(i64, 10))
		}
	case []uint32:
		for _, vi := range v {
			var ui32 = vi
			i64 := int64(ui32)
			tmp = fmt.Sprint(tmp, strconv.FormatInt(i64, 10))
		}
	case []uint64:
		for _, vi := range v {
			var ui64 = vi
			i64 := int64(ui64)
			tmp = fmt.Sprint(tmp, strconv.FormatInt(i64, 10))
		}
	default:
		tmp = fmt.Sprint(tmp, "")
	}

	return tmp
}

//format cases float
func formatFloats(tmp string, val interface{}) string {
	switch v := val.(type) {
	case float32:
		var f32 = v
		f64 := float64(f32)
		tmp = fmt.Sprint(tmp, strconv.FormatFloat(f64, 'f', -1, 64))
	case []float32:
		for _, vf32 := range v {
			var f32 = vf32
			f64 := float64(f32)
			tmp = fmt.Sprint(tmp, strconv.FormatFloat(f64, 'f', -1, 64))
		}
	case float64:
		tmp = fmt.Sprint(tmp, strconv.FormatFloat(v, 'f', -1, 64))
	case []float64:
		for _, vf64 := range v {
			tmp = fmt.Sprint(tmp, strconv.FormatFloat(vf64, 'f', -1, 64))
		}
	default:
		tmp = fmt.Sprint(tmp, "")
	}
	return tmp
}

//format cases bool
func formatBools(tmp string, val interface{}) string {
	switch v := val.(type) {
	case bool:
		tmp = fmt.Sprint(tmp, strconv.FormatBool(v))
	case []bool:
		for _, vb := range v {
			tmp = fmt.Sprint(tmp, strconv.FormatBool(vb))
		}
	default:
		tmp = fmt.Sprint(tmp, "")
	}
	return tmp
}
