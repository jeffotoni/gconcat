package gconcat

import (
	"testing"
)

// go test -v -run ^TestNewFormat
func TestNewFormat(t *testing.T) {
	f := NewFormat()
	if f == nil {
		t.Errorf("Error of create new format")
	}
}

// go test -v -run ^TestFormatMany
func TestFormatMany(t *testing.T) {
	f := NewFormat()
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
		// TODO: Add test cases.
		{"test_Format_", args{str: many1String}, "21"},
		{"test_Format_", args{str: many1Int}, "100"},
		{"test_Format_", args{str: many1Int8}, "100"},
		{"test_Format_", args{str: many1Int16}, "100"},
		{"test_Format_", args{str: many1Int32}, "100"},
		{"test_Format_", args{str: many1Int64}, "100"},
		{"test_Format_", args{str: many1Uint}, "100"},
		{"test_Format_", args{str: many1Uint8}, "100"},
		{"test_Format_", args{str: many1Uint16}, "100"},
		{"test_Format_", args{str: many1Uint32}, "100"},
		{"test_Format_", args{str: many1Uint64}, "100"},
		//{"test_ConcatStr_12", args{str: many1Uintptr}, "d"}, //endereço do ponteiro
		{"test_Format_", args{str: many1Float32}, "100.010002"},
		{"test_Format_", args{str: many1Float64}, "100.010000"},
		{"test_Format_", args{str: many1bool}, "false"},
		{"test_Format_", args{str: many2bool}, "true"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			tt1 := tt
			got := f.Format("", tt1.args.str)
			switch {
			case got == tt1.want:
				break
			default:
				t.Errorf("Format() = %v, want %v", got, tt.want)
			}
			print("\n", got)
		})
	}
}

//go test -v -run ^TestFormatString
func TestFormatString(t *testing.T) {
	sliceString := []string{"jeffotoni", "ancogamer"}
	simpleString := "gconcat"

	tmp := formatString("", sliceString)
	if tmp != "jeffotoniancogamer" {
		t.Errorf("TestFormatString(): %v, want %v", tmp, "jeffotoniancogamer")
	}

	tmp = formatString("", simpleString)
	if tmp != "gconcat" {
		t.Errorf("TestFormatString(): %v, want %v", tmp, "gconcat")
	}
}

//go test -v -run ^TestFormatInts
func TestFormatInts(t *testing.T) {

	tmp := formatInts("", 78)
	if tmp != "78" {
		t.Errorf("TestFormatInts() = %v, want %v", tmp, "78")
	}
	tmp = formatInts("", int8(78))
	if tmp != "78" {
		t.Errorf("TestFormatInts() = %v, want %v", tmp, "78")
	}
	tmp = formatInts("", int16(78))
	if tmp != "78" {
		t.Errorf("TestFormatInts() = %v, want %v", tmp, "78")
	}
	tmp = formatInts("", int32(78))
	if tmp != "78" {
		t.Errorf("TestFormatInts() = %v, want %v", tmp, "78")
	}
	tmp = formatInts("", int64(78))
	if tmp != "78" {
		t.Errorf("TestFormatInts() = %v, want %v", tmp, "78")
	}
	//slices
	tmp = formatInts("", []int{78})
	if tmp != "78" {
		t.Errorf("TestFormatInts() = %v, want %v", tmp, "78")
	}
	tmp = formatInts("", []int8{int8(78)})
	if tmp != "78" {
		t.Errorf("TestFormatInts() = %v, want %v", tmp, "78")
	}
	tmp = formatInts("", []int16{int16(78)})
	if tmp != "78" {
		t.Errorf("TestFormatInts() = %v, want %v", tmp, "78")
	}
	tmp = formatInts("", []int32{int32(78)})
	if tmp != "78" {
		t.Errorf("TestFormatInts() = %v, want %v", tmp, "78")
	}
	tmp = formatInts("", []int64{int64(78)})
	if tmp != "78" {
		t.Errorf("TestFormatInts() = %v, want %v", tmp, "78")
	}
}

//go test -v -run ^TestFormatUints
func TestFormatUints(t *testing.T) {
	tmp := formatUints("", uint(78))
	if tmp != "78" {
		t.Errorf("TestFormatUints() = %v, want %v", tmp, "78")
	}
	tmp = formatUints("", uint8(78))
	if tmp != "78" {
		t.Errorf("TestFormatUints() = %v, want %v", tmp, "78")
	}
	tmp = formatUints("", uint16(78))
	if tmp != "78" {
		t.Errorf("TestFormatUints() = %v, want %v", tmp, "78")
	}
	tmp = formatUints("", uint32(78))
	if tmp != "78" {
		t.Errorf("TestFormatUints() = %v, want %v", tmp, "78")
	}
	tmp = formatUints("", uint64(78))
	if tmp != "78" {
		t.Errorf("TestFormatUints() = %v, want %v", tmp, "78")
	}
	//slices
	tmp = formatUints("", []uint{uint(78)})
	if tmp != "78" {
		t.Errorf("TestFormatUints() = %v, want %v", tmp, "78")
	}
	tmp = formatUints("", []uint8{uint8(78)})
	if tmp != "78" {
		t.Errorf("TestFormatUints() = %v, want %v", tmp, "78")
	}
	tmp = formatUints("", []uint16{uint16(78)})
	if tmp != "78" {
		t.Errorf("TestFormatUints() = %v, want %v", tmp, "78")
	}
	tmp = formatUints("", []uint32{uint32(78)})
	if tmp != "78" {
		t.Errorf("TestFormatUints() = %v, want %v", tmp, "78")
	}
	tmp = formatUints("", []uint64{uint64(78)})
	if tmp != "78" {
		t.Errorf("TestFormatUints() = %v, want %v", tmp, "78")
	}
}

//go test -v -run ^TestFormatFloats
func TestFormatFloats(t *testing.T) {
	tmp := formatFloats("", float32(4.5))
	if tmp != "4.5" {
		t.Errorf("TestFormatFloats() = %v, want %v", tmp, "4.5")
	}
	tmp = formatFloats("", float64(4.5))
	if tmp != "4.5" {
		t.Errorf("TestFormatFloats() = %v, want %v", tmp, "4.5")
	}
	tmp = formatFloats("", []float32{float32(4.5)})
	if tmp != "4.5" {
		t.Errorf("TestFormatFloats() = %v, want %v", tmp, "4.5")
	}
	tmp = formatFloats("", []float64{float64(4.5)})
	if tmp != "4.5" {
		t.Errorf("TestFormatFloats() = %v, want %v", tmp, "4.5")
	}
}
