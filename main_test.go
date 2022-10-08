package gconcat

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"testing"
)

// This function is named ExampleConcat()
// it with the Examples type.
func ExampleConcat() {
	// This example Concat function accepts all types
	// specific func (...interface{}).
	s := Concat("/api/v1/", 39383838, "/", 129494, "/product/", 2012)
	s = Concat(s, ":", 33.22, []string{" - ", "jeffotoni", "-", "2021"})
	fmt.Println(s)
	// Output: /api/v1/39383838/129494/product/2012:33.22 - jeffotoni-2021
}

// This function is named ExampleBuild()
// it with the Examples type.
func ExampleBuild() {
	// This example Build function accepts all types
	// specific func (...interface{}).
	s := Build("jeffotoni", " test example ", 2021)
	fmt.Println(s)
	// Output: jeffotoni test example 2021
}

// This function is named ExampleConcatStr()
// it with the Examples type.
func ExampleConcatStr() {
	// This example ConcatStr function accepts only type string
	// specific func (...string).
	s := ConcatStr("jeffotoni", "/", "2021")
	fmt.Println(s)
	// Output: jeffotoni/2021
}

// This function is named ExampleConcatStrInt()
// it with the Examples type.
func ExampleConcatStrInt() {
	// This example ConcatStrInt function accepts only type string and int
	// specific func (...interface{}).
	s := ConcatStrInt("jeffotoni", "/", 2021, "-", 1000, "/", 239393)
	fmt.Println(s)
	// Output: jeffotoni/2021-1000/239393
}

// This function is named ExampleIntToStringFast()
// it with the Examples type.
func ExampleConcatSliceInt() {
	// This example ConcatSliceInt function accepts only type []int
	// specific func ([]int).
	s := ConcatSliceInt([]int{3, 4, 5, 6, 7, 10, 45, 99})
	fmt.Println(s)
	// Output: 34567104599
}

// This function is named Int32ToStringFast()
// it with the Examples type.
func ExampleConcatSliceInt32() {
	// This example ConcatSliceInt32 function accepts only type []int
	// specific func ([]int32).
	s := ConcatSliceInt32([]int32{3, 4, 5, 6, 7, 10, 45, 99})
	fmt.Println(s)
	// Output: 34567104599
}

// This function is named Int32ToStringFast()
// it with the Examples type.
func ExampleConcatSliceInt64() {
	// This example ConcatSliceInt64 function accepts only type []int
	// specific func ([]int64).
	s := ConcatSliceInt64([]int64{3, 4, 5, 6, 7, 10, 45, 99})
	fmt.Println(s)
	// Output: 34567104599
}

// This function is named ExampleIntToConcatFunc()
// it with the Examples type.
func ExampleIntToConcatFunc() {
	// This example ExampleIntToConcatFunc function accepts only return type int
	// specific func (<anything>) return int
	s := ConcatFunc(func(a, b, c int) int {
		return a + b*c
	}(1, 2, 3))
	fmt.Println(s)
	// Output: 7
}

// This function is named ExampleAnyToConcatFunc()
// it with the Examples type.
func ExampleAnyToConcatFunc() {
	// This example ExampleAnyToConcatFunc function accepts only return type string
	// specific func (<anything>) return string
	f1 := func(a float64) float64 {
		return 1 * 2.2
	}(float64(55.55))

	f2 := func(s string) string {
		return s + "2021"
	}(" hello ")

	f3 := func(a int) int {
		return a * 2
	}(3)

	f4 := func(a []int) (t []int) {
		for _, v := range a {
			t = append(t, v*2)
		}
		return
	}([]int{4, 5, 6, 7, 8})

	f5 := func(a []int) (t []float64) {
		for _, v := range a {
			t = append(t, float64(v)*1.2)
		}
		return
	}([]int{4.0, 5.0, 6.0, 7.0, 8.0})

	s1 := Concat([]bool{true, false, true})
	s := ConcatFunc(f1, f2, f3, f4, f5)
	fmt.Println(s + " " + s1)
	// Output: 2.2 hello 202168101214164.867.1999999999999998.49.6 truefalsetrue
}

// go test -v -run ^TestConcatStr
func TestConcatStr(t *testing.T) {
	s := ConcatStr("jeffotoni", "concat", "go", "only")
	if s != "jeffotoniconcatgoonly" {
		t.Errorf("Error type string to ConcatStr")
	}

	s = ConcatStr()
	if s != "" {
		t.Errorf("Error type string to ConcatStr length zero")
	}
}

// go test -v -run ^TestConcatFuncMany
func TestConcatFuncMany(t *testing.T) {
	// example function of simple returns
	fn := func(i interface{}) interface{} {
		return i
	}

	var many1String, many1Int, many1Int8, many1Int16, many1Int32, many1Int64 interface{}
	var many1Uint, many1Uint8, many1Uint16, many1Uint32, many1Uint64 interface{}
	var many1Float32, many1Float64 interface{}
	var many1bool, many2bool interface{}

	many1bool = false
	many2bool = true
	many1String = "21"
	many1Int = 100
	many1Int8 = int8(100)
	many1Int16 = int16(100)
	many1Int32 = int32(100)
	many1Int64 = int64(100)
	many1Uint = uint(100)
	many1Uint8 = uint8(100)
	many1Uint16 = uint16(100)
	many1Uint32 = uint32(100)
	many1Uint64 = uint64(100)
	many1Float32 = float32(100.01)
	many1Float64 = float64(100.01)
	type args struct {
		str interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"test_ConcatFunc_", args{str: many1String}, "21"},
		{"test_ConcatFunc_", args{str: many1Int}, "100"},
		{"test_ConcatFunc_", args{str: many1Int8}, "100"},
		{"test_ConcatFunc_", args{str: many1Int16}, "100"},
		{"test_ConcatFunc_", args{str: many1Int32}, "100"},
		{"test_ConcatFunc_", args{str: many1Int64}, "100"},
		{"test_ConcatFunc_", args{str: many1Uint}, "100"},
		{"test_ConcatFunc_", args{str: many1Uint8}, "100"},
		{"test_ConcatFunc_", args{str: many1Uint16}, "100"},
		{"test_ConcatFunc_", args{str: many1Uint32}, "100"},
		{"test_ConcatFunc_", args{str: many1Uint64}, "100"},
		{"test_ConcatFunc_", args{str: many1Float32}, "100.010002"},
		{"test_ConcatFunc_", args{str: many1Float64}, "100.010000"},
		{"test_ConcatFunc_", args{str: many1bool}, "false"},
		{"test_ConcatFunc_", args{str: many2bool}, "true"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			tt1 := tt
			got := ConcatFunc(fn(tt1.args.str))
			switch {
			case got == tt1.want:
				break
			default:
				t.Errorf("ConcatFunc() = %v, want %v", got, tt.want)
			}
			print("\n", got)
		})
	}
}

// go test -v -run ^TestConcatFuncNil
func TestConcatFuncNil(t *testing.T) {
	typeNil := func() interface{} {
		return nil
	}
	s := ConcatFunc(typeNil())
	if s != "" {
		t.Errorf("TestConcatFuncNil() = %v, want %v", s, "")
	}
}

// go test -v -run ^TestConcatFuncError
func TestConcatFuncError(t *testing.T) {
	typeError := func() error {
		return errors.New("error")
	}
	s := ConcatFunc(typeError())
	if s != "error" {
		t.Errorf("TestConcatFuncError() = %v, want %v", s, "error")
	}
}

// go test -v -run ^TestConcatFuncSlice
func TestConcatFuncSlice(t *testing.T) {
	// example functions of simple returns
	sliceString := func() []string {
		return []string{"jeffotoni", "func"}
	}
	sliceInt := func() []int {
		return []int{20, 21}
	}

	sliceBool := func() []bool {
		return []bool{true, false}
	}

	s := ConcatFunc(sliceString())
	if s != "jeffotonifunc" {
		t.Errorf("ConcatFunc() = %v, want %v", s, "jeffotonifunc")
	}

	i := ConcatFunc(sliceInt())
	if i != "2021" {
		t.Errorf("ConcatFunc() = %v, want %v", i, "2021")
	}

	b := ConcatFunc(sliceBool())
	if b != "truefalse" {
		t.Errorf("ConcatFunc() = %v, want %v", b, "truefalse")
	}

}

// go test -v -run ^TestConcatStrInt
func TestConcatStrInt(t *testing.T) {
	s := ConcatStrInt("jeffotoni", 2021, "concat", 100, "go", 9800, "only")
	if s != "jeffotoni2021concat100go9800only" {
		t.Errorf("Error type string to ConcatStr")
	}

	s = ConcatStrInt()
	if s != "" {
		t.Errorf("Error type string to ConcatStr length zero")
	}
}

// go test -v -run ^TestStringType
func TestStringType(t *testing.T) {
	var st string = "765423-2021"
	s := ConcatString(st)
	if s != "765423-2021" {
		t.Errorf("Error type string")
	}
}

// go test -v -run ^TestStringIntType
func TestStringIntType(t *testing.T) {
	var st string = "765423-2021-"
	var it int = 765423
	s := ConcatStringInt(st, it)
	if s != "765423-2021-765423" {
		t.Errorf("Error type string and int")
	}
}

// go test -v -run ^TestIntType
func TestIntType(t *testing.T) {
	var i64 int64 = 3456789765423
	s := Concat(i64)
	if s != "3456789765423" {
		t.Errorf("Error type int64")
	}

	var i32 int32 = 189765423
	s = Concat(i32)
	if s != "189765423" {
		t.Errorf("Error type int32")
	}

	var it int = 765423
	s = Concat(it)
	if s != "765423" {
		t.Errorf("Error type int")
	}
}

// go test -v -run ^TestInterface
func TestInterface(t *testing.T) {
	var iv []interface{}
	s := Concat(iv)
	if s != "" {
		t.Errorf("Error type []intterface")
	}

	var ivv []interface{}
	ivv = append(ivv, []int{1, 2, 3})
	ivv = append(ivv, []string{"1", "2", "3"})
	s = Concat(ivv)
	var want string = "123123"
	if s != want {
		t.Errorf("Error want=%s got=%s", want, s)
	}
}

// go test -v -run ^TestZero
func TestZero(t *testing.T) {
	type typeZero struct{}
	var tn typeZero
	s := Concat(tn)
	if s != "<not exist type without suport>" {
		t.Errorf("Error type no exist")
	}

	s = Concat(nil)
	if s != "nil" {
		t.Errorf("Error nil")
	}
}

// go test -v -run ^TestWithFunc
func TestWithFunc(t *testing.T) {
	nothing := func() {}
	s := Concat(nothing)
	if s != "<not exist type without suport>" {
		t.Errorf("Error not exist type")
	}
}

// go test -v -run ^TestComplex
func TestComplex(t *testing.T) {
	var cplx64 complex64
	s := Concat(cplx64)
	if s != "not suport complex 64" {
		t.Errorf("Error complex64 not suport")
	}

	var cplx128 complex128
	s = Concat(cplx128)
	if s != "not suport complex 128" {
		t.Errorf("Error complex128 not suport")
	}
}

// go test -v -run ^TestIntTypes
func TestIntTypes(t *testing.T) {
	//[]int empty
	var a []int
	s := ConcatSliceInt(a)
	if s != "" {
		t.Errorf("Error ConcatSliceInt: return empty")
	}

	//[]int
	var ai []int = []int{23}
	si := ConcatSliceInt(ai)
	if si != "23" {
		t.Errorf("Error ConcatSliceInt: %v, want %v", si, "23")
	}

	//[]int8 empty
	var ei8 []int8
	se8 := ConcatSliceInt8(ei8)
	if se8 != "" {
		t.Errorf("Error ConcatSliceInt8: return empty")
	}

	//[]int8
	var i8 []int8 = []int8{int8(23)}
	si8 := ConcatSliceInt8(i8)
	if si8 != "23" {
		t.Errorf("Error ConcatSliceInt8: %v, want %v", si8, "23")
	}

	//[]int16 empty
	var ei16 []int16
	se16 := ConcatSliceInt16(ei16)
	if se16 != "" {
		t.Errorf("Error ConcatSliceInt16: return empty")
	}

	//[]int16
	var i16 []int16 = []int16{int16(23)}
	si16 := ConcatSliceInt16(i16)
	if si16 != "23" {
		t.Errorf("Error ConcatSliceInt16: %v, want %v", si16, "23")
	}

	//[]int32 empty
	var ei32 []int32
	se32 := ConcatSliceInt32(ei32)
	if se32 != "" {
		t.Errorf("Error ConcatSliceInt32: return empty")
	}

	//[]int32
	var i32 []int32 = []int32{int32(23)}
	s32 := ConcatSliceInt32(i32)
	if s32 != "23" {
		t.Errorf("Error ConcatSliceInt32: %v, want %v", s32, "23")
	}

	//[]int64 empty
	var ei64 []int64
	se64 := ConcatSliceInt64(ei64)
	if se64 != "" {
		t.Errorf("Error ConcatSliceInt64: return empty")
	}

	//[]int64 empty
	var i64 []int64 = []int64{int64(23)}
	s64 := ConcatSliceInt64(i64)
	if s64 != "23" {
		t.Errorf("Error ConcatSliceInt64: %v, want %v", s64, "23")
	}
}

// go test -v -run ^TestUIntTypes
func TestUIntTypes(t *testing.T) {
	//[]uint empty
	var a []uint
	s := ConcatSliceUint(a)
	if s != "" {
		t.Errorf("Error ConcatSliceUint: return empty")
	}

	//[]uint
	var ai []uint = []uint{23}
	si := ConcatSliceUint(ai)
	if si != "23" {
		t.Errorf("Error ConcatSliceUint: %v, want %v", si, "23")
	}

	//[]uint8 empty
	var eu8 []uint8
	se8 := ConcatSliceUint8(eu8)
	if se8 != "" {
		t.Errorf("Error ConcatSliceUint8: return empty")
	}

	//[]uint8
	var u8 []uint8 = []uint8{uint8(23)}
	su8 := ConcatSliceUint8(u8)
	if su8 != "23" {
		t.Errorf("Error ConcatSliceUint8: %v, want %v", su8, "23")
	}

	//[]uint16 empty
	var eu16 []uint16
	se16 := ConcatSliceUint16(eu16)
	if se16 != "" {
		t.Errorf("Error ConcatSliceUint16: return empty")
	}

	//[]uint16
	var u16 []uint16 = []uint16{uint16(23)}
	su16 := ConcatSliceUint16(u16)
	if su16 != "23" {
		t.Errorf("Error ConcatSliceUint16: %v, want %v", su16, "23")
	}

	//[]uint32 empty
	var eu32 []uint32
	se32 := ConcatSliceUint32(eu32)
	if se32 != "" {
		t.Errorf("Error ConcatSliceUint32: return empty")
	}

	//[]uint32
	var u32 []uint32 = []uint32{uint32(23)}
	su32 := ConcatSliceUint32(u32)
	if su32 != "23" {
		t.Errorf("Error ConcatSliceUint32: %v, want %v", su32, "23")
	}

	//[]uint64 empty
	var eu64 []uint64
	se64 := ConcatSliceUint64(eu64)
	if se64 != "" {
		t.Errorf("Error ConcatSliceUint64: return empty")
	}

	//[]uint64
	var u64 []uint64 = []uint64{uint64(23)}
	su64 := ConcatSliceUint64(u64)
	if su64 != "23" {
		t.Errorf("Error ConcatSliceUint64: %v, want %v", su64, "23")
	}
}

// go test -v -run ^TestConcatSliceBool
func TestConcatSliceBool(t *testing.T) {
	//[]bool empty
	var be []bool
	sbe := ConcatSliceBool(be)
	if sbe != "" {
		t.Errorf("Error ConcatSliceBool: return empty")
	}
	//[]bool
	var b []bool = []bool{true}
	sb := ConcatSliceBool(b)
	if sb != "true" {
		t.Errorf("Error ConcatSliceBool: %v, want %v", sb, "")
	}
}

// go test -v -run ^TestConcatFloats
func TestConcatFloats(t *testing.T) {
	//[]float32 empty
	var ef32 []float32
	sef32 := ConcatSliceFloat32(ef32)
	if sef32 != "" {
		t.Errorf("Error ConcatSliceFloat32: return empty")
	}
	//[]float32
	var f32 []float32 = []float32{float32(23)}
	sf32 := ConcatSliceFloat32(f32)
	if sf32 != "23.000000" {
		t.Errorf("Error ConcatSliceFloat32: %v, want %v", sf32, "23.000000")
	}
	//[]float64 empty
	var ef64 []float64
	sef64 := ConcatSliceFloat64(ef64)
	if sef64 != "" {
		t.Errorf("Error ConcatSliceFloat64: return empty")
	}
	//[]float64
	var f64 []float64 = []float64{float64(23)}
	sf64 := ConcatSliceFloat64(f64)
	if sf64 != "23.000000" {
		t.Errorf("Error ConcatSliceFloat64: %v, want %v", sf64, "23.000000")
	}
}

// go test -v -run ^TestConcatMany
func TestConcatMany(t *testing.T) {
	// int8
	var i8 int8 = int8(23)
	si8 := Concat(i8)
	if si8 != "23" {
		t.Errorf("Error TestConcatMany for int8: %v, want %v", si8, "23")
	}
	// uint8
	var ui uint = uint(23)
	su := Concat(ui)
	if su != "23" {
		t.Errorf("Error TestConcatMany for uint: %v, want %v", su, "23")
	}

	// uint8
	var ui8 uint8 = uint8(23)
	su8 := Concat(ui8)
	if su8 != "23" {
		t.Errorf("Error TestConcatMany for uint8: %v, want %v", su8, "23")
	}
	// int16
	var i16 int16 = int16(23)
	si16 := Concat(i16)
	if si16 != "23" {
		t.Errorf("Error TestConcatMany for int16: %v, want %v", si16, "23")
	}
	// uint16
	var ui16 uint16 = uint16(23)
	su16 := Concat(ui16)
	if su16 != "23" {
		t.Errorf("Error TestConcatMany for uint16: %v, want %v", su16, "23")
	}
	// uint32
	var ui32 uint32 = uint32(23)
	su32 := Concat(ui32)
	if su32 != "23" {
		t.Errorf("Error TestConcatMany for uint32: %v, want %v", su32, "23")
	}
	// uint64
	var ui64 uint64 = uint64(23)
	su64 := Concat(ui64)
	if su64 != "23" {
		t.Errorf("Error TestConcatMany for uint64: %v, want %v", su64, "23")
	}

	// float32
	var f32 float32 = float32(23)
	sf32 := Concat(f32)
	if sf32 != "23" {
		t.Errorf("Error TestConcatMany for float32: %v, want %v", sf32, "23.000000")
	}
	// float64
	var f64 float64 = float64(23)
	sf64 := Concat(f64)
	if sf64 != "23" {
		t.Errorf("Error TestConcatMany for float64: %v, want %v", su16, "23.000000")
	}

}

// go test -v -run ^Test_Concat_OneStr
func Test_Concat_OneStr(t *testing.T) {
	var many1String, many1Int, many1Int8, many1Int16, many1Int32, many1Int64 interface{}
	var many1Uint, many1Uint8, many1Uint16, many1Uint32, many1Uint64 interface{}
	var many1Float32, many1Float64 interface{}
	var many1bool, many2bool interface{}

	many1bool = false
	many2bool = true
	many1String = "21"
	many1Int = 100
	many1Int8 = int8(100)
	many1Int16 = int16(100)
	many1Int32 = int32(100)
	many1Int64 = int64(100)
	many1Uint = uint(100)
	many1Uint8 = uint8(100)
	many1Uint16 = uint16(100)
	many1Uint32 = uint32(100)
	many1Uint64 = uint64(100)
	//many1Uintptr = uintptr(100)

	many1Float32 = float32(100.01)
	many1Float64 = float64(100.01)
	type args struct {
		str interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"test_ConcatStr_", args{str: many1String}, "21"},
		{"test_ConcatStr_", args{str: many1Int}, "100"},
		{"test_ConcatStr_", args{str: many1Int8}, "100"},
		{"test_ConcatStr_", args{str: many1Int16}, "100"},
		{"test_ConcatStr_", args{str: many1Int32}, "100"},
		{"test_ConcatStr_", args{str: many1Int64}, "100"},
		{"test_ConcatStr_", args{str: many1Uint}, "100"},
		{"test_ConcatStr_", args{str: many1Uint8}, "100"},
		{"test_ConcatStr_", args{str: many1Uint16}, "100"},
		{"test_ConcatStr_", args{str: many1Uint32}, "100"},
		{"test_ConcatStr_", args{str: many1Uint64}, "100"},
		//{"test_ConcatStr_12", args{str: many1Uintptr}, "d"}, //endereço do ponteiro
		{"test_ConcatStr_", args{str: many1Float32}, "100.010002"},
		{"test_ConcatStr_", args{str: many1Float64}, "100.010000"},
		{"test_ConcatStr_", args{str: many1bool}, "false"},
		{"test_ConcatStr_", args{str: many2bool}, "true"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			tt1 := tt
			got := Concat(tt1.args.str)
			switch {
			case got == tt1.want:
				break
			default:
				t.Errorf("Concat() = %v, want %v", got, tt.want)
			}
			print("\n", got)
		})
	}
}

// go test -v -run ^Test_Concat_Many
func Test_Concat_Many(t *testing.T) {

	var many1String, many1Int, many1Int8, many1Int16, many1Int32, many1Int64 []interface{}
	var many1Uint, many1Uint8, many1Uint16, many1Uint32, many1Uint64 []interface{}
	var many1Float32, many1Float64 []interface{}
	var many1bool []interface{}
	many1bool = append(many1bool, []bool{true, false})
	many1String = append(many1String, []string{"21", "joao", "beta"})
	many1Int = append(many1Int, []int{100, 98, 23})
	many1Int8 = append(many1Int8, []int8{100, 98, 23})
	many1Int16 = append(many1Int16, []int16{100, 98, 23})
	many1Int32 = append(many1Int32, []int32{100, 98, 23})
	many1Int64 = append(many1Int64, []int64{100, 98, 23})
	many1Uint = append(many1Uint, []uint{100, 98, 23})
	many1Uint8 = append(many1Uint8, []uint8{100, 98, 23})
	many1Uint16 = append(many1Uint16, []uint16{100, 98, 23})
	many1Uint32 = append(many1Uint32, []uint32{100, 98, 23})
	many1Uint64 = append(many1Uint64, []uint64{100, 98, 23})
	//many1Uintptr = append(many1Uintptr, []uintptr{100, 1, 3})

	many1Float32 = append(many1Float32, []float32{100.01, 100.01})
	many1Float64 = append(many1Float64, []float64{100.01, 100.01})
	type args struct {
		strs []interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		//{"test_concat_12", args{strs: many1Uintptr}, "d"}, //endereço do ponteiro
		{"test_concat_", args{strs: many1bool}, "truefalse"},
		{"test_concat_", args{strs: many1String}, "21joaobeta"},
		{"test_concat_", args{strs: many1Int}, "1009823"},
		{"test_concat_", args{strs: many1Int8}, "1009823"},
		{"test_concat_", args{strs: many1Int16}, "1009823"},
		{"test_concat_", args{strs: many1Int32}, "1009823"},
		{"test_concat_", args{strs: many1Int64}, "1009823"},
		{"test_concat_", args{strs: many1Uint}, "1009823"},
		{"test_concat_", args{strs: many1Uint8}, "1009823"},
		{"test_concat_", args{strs: many1Uint16}, "1009823"},
		{"test_concat_", args{strs: many1Uint32}, "1009823"},
		{"test_concat_", args{strs: many1Uint64}, "1009823"},
		{"test_concat_", args{strs: many1Float32}, "100.010002100.010002"},
		{"test_concat_", args{strs: many1Float64}, "100.010000100.010000"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//t.Parallel()
			tt1 := tt
			got := Concat(tt1.args.strs...)
			switch {
			case got == tt1.want:
				break
			default:
				t.Errorf("Concat() = %v, want %v", got, tt.want)
			}
			print("\n", got)
		})
	}
}

var longStr string = `qwertyuiopqwertyuiopqwertyuio
qwertyuiopqwertyuiopqwertyuiopqwertyuiopqwertyuiopqwertyuiopqwertyuiop`

func BenchmarkConcatVector(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = Concat(
			[]int{1, 2, 3, 4, 5, 56, 6, 7, 7, 778, 8, 88, 8},
			//[]float32{1.1, 2.2, 3.4, 4.4, 5.0, 56.9, 6.0, 7.8},
			[]float64{1.1, 2.2, 3.4, 4.4, 5.0, 56.9, 6.0, 7.8},
			[]string{"jeff", "otoni", "lima", " ", " vamos testar concat Go "},
			//[]string{"jeff", "otoni", "lima", " ", " vamos testar concat Go "},
			// []string{"jeff", "otoni", "lima", " ", " vamos testar concat Go "},
			// []string{"jeff", "otoni", "lima", " ", " vamos testar concat Go "},
			// []string{"jeff", "otoni", "lima", " ", " vamos testar concat Go "},
			// []string{"jeff", "otoni", "lima", " ", " vamos testar concat Go "},
			// []string{"jeff", "otoni", "lima", " ", " vamos testar concat Go "},
			// []string{"jeff", "otoni", "lima", " ", " vamos testar concat Go "},
			// []string{"jeff", "otoni", "lima", " ", " vamos testar concat Go "},
		)
	}
}

func BenchmarkConcatFuncString(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = ConcatFunc(func(a, b, c int) string { return "jeffstring" }(1, 2, 3))
	}
}

func BenchmarkConcatFuncInt(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = ConcatFunc(func(a, b, c int) int { return 1345 }(1, 2, 3))
	}
}

func BenchmarkConcatFuncStringVector(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = ConcatFunc(func(a, b, c int) []string {
			var vect []string
			vect = append(vect, "jeff")
			return vect
		}(1, 2, 3))
	}
}

func BenchmarkConcatFuncFuncAny(b *testing.B) {
	for n := 0; n < b.N; n++ {
		f1 := func(a int) int {
			return a * 3
		}(n)

		f2 := func(s string) string {
			return s + "2021"
		}(" hello ")
		_ = ConcatFunc(f1, f2)
	}
}

func BenchmarkConcatStr(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = ConcatStr("string_jeffotoni here:", longStr, longStr, longStr, longStr, " go is life ", "2022")
	}
}

func BenchmarkCopyStr(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = CopyStr("string_jeffotoni here:", longStr, longStr, longStr, longStr, " go is life ", "2022")
	}
}

func BenchmarkConcatStrCopy(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = ConcatStrCopy("string_jeffotoni here:", longStr, longStr, longStr, longStr, " go is life ", "2022")
	}
}

func BenchmarkConcatStrInt(b *testing.B) {
	for n := 0; n < b.N; n++ {
		ConcatStrInt("string_jeffotoni", 2021, longStr, longStr, 99, longStr, 100, longStr)
	}
}

func BenchmarkConcat(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Concat("string_jeffotoni", longStr, longStr, longStr, longStr)
	}
}

func BenchmarkConcatSliceIntString(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Concat(2021, "jeffotoni", "Go", "is love", " ", 1233, "jeff", 99, "somente", "string heree........")
	}
}

func BenchmarkLongJoin(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = strings.Join([]string{"string_jeffotoni%s", longStr, longStr, longStr, longStr}, "")
	}
}

func BenchmarkLongSprintf(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = fmt.Sprintf("string_jeffotoni%s%s%s%s", longStr, longStr, longStr, longStr)
	}
}

func BenchmarkBuilder(b *testing.B) {
	var bb strings.Builder
	for n := 0; n < b.N; n++ {
		bb.WriteString("string_jeffotoni")
		bb.WriteString(longStr)
		bb.WriteString(longStr)
		bb.WriteString(longStr)
		bb.WriteString(longStr)
	}
}

func BenchmarkMarshal(b *testing.B) {
	type X struct {
		Num   []int
		Str1  []string
		Str2  []string
		Str3  []string
		Str4  []string
		Str5  []string
		Str6  []string
		Str7  []string
		Str8  []string
		Str9  []string
		Str10 []string
	}

	var obj = X{
		Num: []int{1, 2, 3, 4, 5, 56, 6, 7, 7, 778, 8, 88, 8,
			8, 8, 8, 8, 8, 9, 9, 123, 4, 4, 5, 6, 7, 77,
			8, 8, 99, 9, 93, 3, 3, 3, 3, 45, 5, 6, 6, 7},
		Str1:  []string{"jeff", "otoni", "lima", " ", " vamos testar concat Go "},
		Str2:  []string{"jeff", "otoni", "lima", " ", " vamos testar concat Go "},
		Str3:  []string{"jeff", "otoni", "lima", " ", " vamos testar concat Go "},
		Str4:  []string{"jeff", "otoni", "lima", " ", " vamos testar concat Go "},
		Str5:  []string{"jeff", "otoni", "lima", " ", " vamos testar concat Go "},
		Str6:  []string{"jeff", "otoni", "lima", " ", " vamos testar concat Go "},
		Str7:  []string{"jeff", "otoni", "lima", " ", " vamos testar concat Go "},
		Str8:  []string{"jeff", "otoni", "lima", " ", " vamos testar concat Go "},
		Str9:  []string{"jeff", "otoni", "lima", " ", " vamos testar concat Go "},
		Str10: []string{"jeff", "somente", "string heree........"},
	}

	for n := 0; n < b.N; n++ {
		_, err := json.Marshal(&obj)
		if err != nil {
			panic(err)
		}
	}
}

// go test -v -failfast -run ^TestConcatFunc$
func TestConcatFunc(t *testing.T) {
	type args struct {
		f []interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test",
			args: args{
				f: []interface{}{
					func(a, b, c int) int {
						return a + b*c
					}(1, 2, 3),
				},
			},
			want: "7",
		},
		{
			name: "test with function as result",
			args: args{
				f: []interface{}{
					func(a, b, c int) func() int {
						return func() int { x := 0; return x }
					}(1, 2, 3),
				},
			},
			want: "0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConcatFunc(tt.args.f...); got != tt.want {
				t.Errorf("ConcatFunc() = %v, want %v", got, tt.want)
			}
		})
	}
}

var concatFuncResp string

func BenchmarkConcatFunc(b *testing.B) {

	for i := 0; i < b.N; i++ {
		concatFuncResp = ConcatFunc(func(a, b, c int) int {
			return a + b*c
		}(1, 2, 3), []int{1, 2, 3, 4})
	}

}

// go test -v -failfast -run ^TestConcatFuncOLD$
func TestConcatFuncOLD(t *testing.T) {
	type args struct {
		f []interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test",
			args: args{
				f: []interface{}{
					func(a, b, c int) int {
						return a + b*c
					}(1, 2, 3),
				},
			},
			want: "7",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConcatFuncOLD(tt.args.f...); got != tt.want {
				t.Errorf("ConcatFuncOLD() = %v, want %v", got, tt.want)
			}
		})
	}
}

var concatFuncRespOLD string

func BenchmarkConcatFuncOLD(b *testing.B) {

	for i := 0; i < b.N; i++ {
		concatFuncRespOLD = ConcatFuncOLD(func(a, b, c int) int {
			return a + b*c
		}(1, 2, 3), []int{1, 2, 3, 4})
	}

}
