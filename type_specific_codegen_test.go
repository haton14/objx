package objx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInter(t *testing.T) {

	m := map[string]interface{}{"value": interface{}("something"), "nothing": nil}
	assert.Equal(t, interface{}("something"), New(m).Get("value").Inter())
	assert.Equal(t, interface{}("something"), New(m).Get("value").MustInter())
	assert.Equal(t, interface{}(nil), New(m).Get("nothing").Inter())
	assert.Equal(t, interface{}("something"), New(m).Get("nothing").Inter("something"))

	assert.Panics(t, func() {
		New(m).Get("age").MustInter()
	})

}

func TestInterSlice(t *testing.T) {

	m := map[string]interface{}{"value": []interface{}{interface{}("something")}, "nothing": nil}
	assert.Equal(t, interface{}("something"), New(m).Get("value").InterSlice()[0])
	assert.Equal(t, interface{}("something"), New(m).Get("value").MustInterSlice()[0])
	assert.Equal(t, []interface{}(nil), New(m).Get("nothing").InterSlice())
	assert.Equal(t, interface{}("something"), New(m).Get("nothing").InterSlice([]interface{}{interface{}("something")})[0])

	assert.Panics(t, func() {
		New(m).Get("nothing").MustInterSlice()
	})

}

func TestBool(t *testing.T) {

	m := map[string]interface{}{"value": bool(true), "nothing": nil}
	assert.Equal(t, bool(true), New(m).Get("value").Bool())
	assert.Equal(t, bool(true), New(m).Get("value").MustBool())
	assert.Equal(t, bool(false), New(m).Get("nothing").Bool())
	assert.Equal(t, bool(true), New(m).Get("nothing").Bool(true))

	assert.Panics(t, func() {
		New(m).Get("age").MustBool()
	})

}

func TestBoolSlice(t *testing.T) {

	m := map[string]interface{}{"value": []bool{bool(true)}, "nothing": nil}
	assert.Equal(t, bool(true), New(m).Get("value").BoolSlice()[0])
	assert.Equal(t, bool(true), New(m).Get("value").MustBoolSlice()[0])
	assert.Equal(t, []bool(nil), New(m).Get("nothing").BoolSlice())
	assert.Equal(t, bool(true), New(m).Get("nothing").BoolSlice([]bool{bool(true)})[0])

	assert.Panics(t, func() {
		New(m).Get("nothing").MustBoolSlice()
	})

}

func TestIsBool(t *testing.T) {

	var o *O

	o = New(bool(true))
	assert.True(t, o.IsBool())

	o = New([]bool{bool(true)})
	assert.True(t, o.IsBoolSlice())

}

func TestStr(t *testing.T) {

	m := map[string]interface{}{"value": string("hello"), "nothing": nil}
	assert.Equal(t, string("hello"), New(m).Get("value").Str())
	assert.Equal(t, string("hello"), New(m).Get("value").MustStr())
	assert.Equal(t, string(""), New(m).Get("nothing").Str())
	assert.Equal(t, string("hello"), New(m).Get("nothing").Str("hello"))

	assert.Panics(t, func() {
		New(m).Get("age").MustStr()
	})

}

func TestStrSlice(t *testing.T) {

	m := map[string]interface{}{"value": []string{string("hello")}, "nothing": nil}
	assert.Equal(t, string("hello"), New(m).Get("value").StrSlice()[0])
	assert.Equal(t, string("hello"), New(m).Get("value").MustStrSlice()[0])
	assert.Equal(t, []string(nil), New(m).Get("nothing").StrSlice())
	assert.Equal(t, string("hello"), New(m).Get("nothing").StrSlice([]string{string("hello")})[0])

	assert.Panics(t, func() {
		New(m).Get("nothing").MustStrSlice()
	})

}

func TestIsStr(t *testing.T) {

	var o *O

	o = New(string("hello"))
	assert.True(t, o.IsStr())

	o = New([]string{string("hello")})
	assert.True(t, o.IsStrSlice())

}

func TestInt(t *testing.T) {

	m := map[string]interface{}{"value": int(1), "nothing": nil}
	assert.Equal(t, int(1), New(m).Get("value").Int())
	assert.Equal(t, int(1), New(m).Get("value").MustInt())
	assert.Equal(t, int(0), New(m).Get("nothing").Int())
	assert.Equal(t, int(1), New(m).Get("nothing").Int(1))

	assert.Panics(t, func() {
		New(m).Get("age").MustInt()
	})

}

func TestIntSlice(t *testing.T) {

	m := map[string]interface{}{"value": []int{int(1)}, "nothing": nil}
	assert.Equal(t, int(1), New(m).Get("value").IntSlice()[0])
	assert.Equal(t, int(1), New(m).Get("value").MustIntSlice()[0])
	assert.Equal(t, []int(nil), New(m).Get("nothing").IntSlice())
	assert.Equal(t, int(1), New(m).Get("nothing").IntSlice([]int{int(1)})[0])

	assert.Panics(t, func() {
		New(m).Get("nothing").MustIntSlice()
	})

}

func TestIsInt(t *testing.T) {

	var o *O

	o = New(int(1))
	assert.True(t, o.IsInt())

	o = New([]int{int(1)})
	assert.True(t, o.IsIntSlice())

}

func TestInt8(t *testing.T) {

	m := map[string]interface{}{"value": int8(1), "nothing": nil}
	assert.Equal(t, int8(1), New(m).Get("value").Int8())
	assert.Equal(t, int8(1), New(m).Get("value").MustInt8())
	assert.Equal(t, int8(0), New(m).Get("nothing").Int8())
	assert.Equal(t, int8(1), New(m).Get("nothing").Int8(1))

	assert.Panics(t, func() {
		New(m).Get("age").MustInt8()
	})

}

func TestInt8Slice(t *testing.T) {

	m := map[string]interface{}{"value": []int8{int8(1)}, "nothing": nil}
	assert.Equal(t, int8(1), New(m).Get("value").Int8Slice()[0])
	assert.Equal(t, int8(1), New(m).Get("value").MustInt8Slice()[0])
	assert.Equal(t, []int8(nil), New(m).Get("nothing").Int8Slice())
	assert.Equal(t, int8(1), New(m).Get("nothing").Int8Slice([]int8{int8(1)})[0])

	assert.Panics(t, func() {
		New(m).Get("nothing").MustInt8Slice()
	})

}

func TestIsInt8(t *testing.T) {

	var o *O

	o = New(int8(1))
	assert.True(t, o.IsInt8())

	o = New([]int8{int8(1)})
	assert.True(t, o.IsInt8Slice())

}

func TestInt16(t *testing.T) {

	m := map[string]interface{}{"value": int16(1), "nothing": nil}
	assert.Equal(t, int16(1), New(m).Get("value").Int16())
	assert.Equal(t, int16(1), New(m).Get("value").MustInt16())
	assert.Equal(t, int16(0), New(m).Get("nothing").Int16())
	assert.Equal(t, int16(1), New(m).Get("nothing").Int16(1))

	assert.Panics(t, func() {
		New(m).Get("age").MustInt16()
	})

}

func TestInt16Slice(t *testing.T) {

	m := map[string]interface{}{"value": []int16{int16(1)}, "nothing": nil}
	assert.Equal(t, int16(1), New(m).Get("value").Int16Slice()[0])
	assert.Equal(t, int16(1), New(m).Get("value").MustInt16Slice()[0])
	assert.Equal(t, []int16(nil), New(m).Get("nothing").Int16Slice())
	assert.Equal(t, int16(1), New(m).Get("nothing").Int16Slice([]int16{int16(1)})[0])

	assert.Panics(t, func() {
		New(m).Get("nothing").MustInt16Slice()
	})

}

func TestIsInt16(t *testing.T) {

	var o *O

	o = New(int16(1))
	assert.True(t, o.IsInt16())

	o = New([]int16{int16(1)})
	assert.True(t, o.IsInt16Slice())

}

func TestInt32(t *testing.T) {

	m := map[string]interface{}{"value": int32(1), "nothing": nil}
	assert.Equal(t, int32(1), New(m).Get("value").Int32())
	assert.Equal(t, int32(1), New(m).Get("value").MustInt32())
	assert.Equal(t, int32(0), New(m).Get("nothing").Int32())
	assert.Equal(t, int32(1), New(m).Get("nothing").Int32(1))

	assert.Panics(t, func() {
		New(m).Get("age").MustInt32()
	})

}

func TestInt32Slice(t *testing.T) {

	m := map[string]interface{}{"value": []int32{int32(1)}, "nothing": nil}
	assert.Equal(t, int32(1), New(m).Get("value").Int32Slice()[0])
	assert.Equal(t, int32(1), New(m).Get("value").MustInt32Slice()[0])
	assert.Equal(t, []int32(nil), New(m).Get("nothing").Int32Slice())
	assert.Equal(t, int32(1), New(m).Get("nothing").Int32Slice([]int32{int32(1)})[0])

	assert.Panics(t, func() {
		New(m).Get("nothing").MustInt32Slice()
	})

}

func TestIsInt32(t *testing.T) {

	var o *O

	o = New(int32(1))
	assert.True(t, o.IsInt32())

	o = New([]int32{int32(1)})
	assert.True(t, o.IsInt32Slice())

}

func TestInt64(t *testing.T) {

	m := map[string]interface{}{"value": int64(1), "nothing": nil}
	assert.Equal(t, int64(1), New(m).Get("value").Int64())
	assert.Equal(t, int64(1), New(m).Get("value").MustInt64())
	assert.Equal(t, int64(0), New(m).Get("nothing").Int64())
	assert.Equal(t, int64(1), New(m).Get("nothing").Int64(1))

	assert.Panics(t, func() {
		New(m).Get("age").MustInt64()
	})

}

func TestInt64Slice(t *testing.T) {

	m := map[string]interface{}{"value": []int64{int64(1)}, "nothing": nil}
	assert.Equal(t, int64(1), New(m).Get("value").Int64Slice()[0])
	assert.Equal(t, int64(1), New(m).Get("value").MustInt64Slice()[0])
	assert.Equal(t, []int64(nil), New(m).Get("nothing").Int64Slice())
	assert.Equal(t, int64(1), New(m).Get("nothing").Int64Slice([]int64{int64(1)})[0])

	assert.Panics(t, func() {
		New(m).Get("nothing").MustInt64Slice()
	})

}

func TestIsInt64(t *testing.T) {

	var o *O

	o = New(int64(1))
	assert.True(t, o.IsInt64())

	o = New([]int64{int64(1)})
	assert.True(t, o.IsInt64Slice())

}

func TestUint(t *testing.T) {

	m := map[string]interface{}{"value": uint(1), "nothing": nil}
	assert.Equal(t, uint(1), New(m).Get("value").Uint())
	assert.Equal(t, uint(1), New(m).Get("value").MustUint())
	assert.Equal(t, uint(0), New(m).Get("nothing").Uint())
	assert.Equal(t, uint(1), New(m).Get("nothing").Uint(1))

	assert.Panics(t, func() {
		New(m).Get("age").MustUint()
	})

}

func TestUintSlice(t *testing.T) {

	m := map[string]interface{}{"value": []uint{uint(1)}, "nothing": nil}
	assert.Equal(t, uint(1), New(m).Get("value").UintSlice()[0])
	assert.Equal(t, uint(1), New(m).Get("value").MustUintSlice()[0])
	assert.Equal(t, []uint(nil), New(m).Get("nothing").UintSlice())
	assert.Equal(t, uint(1), New(m).Get("nothing").UintSlice([]uint{uint(1)})[0])

	assert.Panics(t, func() {
		New(m).Get("nothing").MustUintSlice()
	})

}

func TestIsUint(t *testing.T) {

	var o *O

	o = New(uint(1))
	assert.True(t, o.IsUint())

	o = New([]uint{uint(1)})
	assert.True(t, o.IsUintSlice())

}

func TestUint8(t *testing.T) {

	m := map[string]interface{}{"value": uint8(1), "nothing": nil}
	assert.Equal(t, uint8(1), New(m).Get("value").Uint8())
	assert.Equal(t, uint8(1), New(m).Get("value").MustUint8())
	assert.Equal(t, uint8(0), New(m).Get("nothing").Uint8())
	assert.Equal(t, uint8(1), New(m).Get("nothing").Uint8(1))

	assert.Panics(t, func() {
		New(m).Get("age").MustUint8()
	})

}

func TestUint8Slice(t *testing.T) {

	m := map[string]interface{}{"value": []uint8{uint8(1)}, "nothing": nil}
	assert.Equal(t, uint8(1), New(m).Get("value").Uint8Slice()[0])
	assert.Equal(t, uint8(1), New(m).Get("value").MustUint8Slice()[0])
	assert.Equal(t, []uint8(nil), New(m).Get("nothing").Uint8Slice())
	assert.Equal(t, uint8(1), New(m).Get("nothing").Uint8Slice([]uint8{uint8(1)})[0])

	assert.Panics(t, func() {
		New(m).Get("nothing").MustUint8Slice()
	})

}

func TestIsUint8(t *testing.T) {

	var o *O

	o = New(uint8(1))
	assert.True(t, o.IsUint8())

	o = New([]uint8{uint8(1)})
	assert.True(t, o.IsUint8Slice())

}

func TestUint16(t *testing.T) {

	m := map[string]interface{}{"value": uint16(1), "nothing": nil}
	assert.Equal(t, uint16(1), New(m).Get("value").Uint16())
	assert.Equal(t, uint16(1), New(m).Get("value").MustUint16())
	assert.Equal(t, uint16(0), New(m).Get("nothing").Uint16())
	assert.Equal(t, uint16(1), New(m).Get("nothing").Uint16(1))

	assert.Panics(t, func() {
		New(m).Get("age").MustUint16()
	})

}

func TestUint16Slice(t *testing.T) {

	m := map[string]interface{}{"value": []uint16{uint16(1)}, "nothing": nil}
	assert.Equal(t, uint16(1), New(m).Get("value").Uint16Slice()[0])
	assert.Equal(t, uint16(1), New(m).Get("value").MustUint16Slice()[0])
	assert.Equal(t, []uint16(nil), New(m).Get("nothing").Uint16Slice())
	assert.Equal(t, uint16(1), New(m).Get("nothing").Uint16Slice([]uint16{uint16(1)})[0])

	assert.Panics(t, func() {
		New(m).Get("nothing").MustUint16Slice()
	})

}

func TestIsUint16(t *testing.T) {

	var o *O

	o = New(uint16(1))
	assert.True(t, o.IsUint16())

	o = New([]uint16{uint16(1)})
	assert.True(t, o.IsUint16Slice())

}

func TestUint32(t *testing.T) {

	m := map[string]interface{}{"value": uint32(1), "nothing": nil}
	assert.Equal(t, uint32(1), New(m).Get("value").Uint32())
	assert.Equal(t, uint32(1), New(m).Get("value").MustUint32())
	assert.Equal(t, uint32(0), New(m).Get("nothing").Uint32())
	assert.Equal(t, uint32(1), New(m).Get("nothing").Uint32(1))

	assert.Panics(t, func() {
		New(m).Get("age").MustUint32()
	})

}

func TestUint32Slice(t *testing.T) {

	m := map[string]interface{}{"value": []uint32{uint32(1)}, "nothing": nil}
	assert.Equal(t, uint32(1), New(m).Get("value").Uint32Slice()[0])
	assert.Equal(t, uint32(1), New(m).Get("value").MustUint32Slice()[0])
	assert.Equal(t, []uint32(nil), New(m).Get("nothing").Uint32Slice())
	assert.Equal(t, uint32(1), New(m).Get("nothing").Uint32Slice([]uint32{uint32(1)})[0])

	assert.Panics(t, func() {
		New(m).Get("nothing").MustUint32Slice()
	})

}

func TestIsUint32(t *testing.T) {

	var o *O

	o = New(uint32(1))
	assert.True(t, o.IsUint32())

	o = New([]uint32{uint32(1)})
	assert.True(t, o.IsUint32Slice())

}

func TestUint64(t *testing.T) {

	m := map[string]interface{}{"value": uint64(1), "nothing": nil}
	assert.Equal(t, uint64(1), New(m).Get("value").Uint64())
	assert.Equal(t, uint64(1), New(m).Get("value").MustUint64())
	assert.Equal(t, uint64(0), New(m).Get("nothing").Uint64())
	assert.Equal(t, uint64(1), New(m).Get("nothing").Uint64(1))

	assert.Panics(t, func() {
		New(m).Get("age").MustUint64()
	})

}

func TestUint64Slice(t *testing.T) {

	m := map[string]interface{}{"value": []uint64{uint64(1)}, "nothing": nil}
	assert.Equal(t, uint64(1), New(m).Get("value").Uint64Slice()[0])
	assert.Equal(t, uint64(1), New(m).Get("value").MustUint64Slice()[0])
	assert.Equal(t, []uint64(nil), New(m).Get("nothing").Uint64Slice())
	assert.Equal(t, uint64(1), New(m).Get("nothing").Uint64Slice([]uint64{uint64(1)})[0])

	assert.Panics(t, func() {
		New(m).Get("nothing").MustUint64Slice()
	})

}

func TestIsUint64(t *testing.T) {

	var o *O

	o = New(uint64(1))
	assert.True(t, o.IsUint64())

	o = New([]uint64{uint64(1)})
	assert.True(t, o.IsUint64Slice())

}

func TestUintptr(t *testing.T) {

	m := map[string]interface{}{"value": uintptr(1), "nothing": nil}
	assert.Equal(t, uintptr(1), New(m).Get("value").Uintptr())
	assert.Equal(t, uintptr(1), New(m).Get("value").MustUintptr())
	assert.Equal(t, uintptr(0), New(m).Get("nothing").Uintptr())
	assert.Equal(t, uintptr(1), New(m).Get("nothing").Uintptr(1))

	assert.Panics(t, func() {
		New(m).Get("age").MustUintptr()
	})

}

func TestUintptrSlice(t *testing.T) {

	m := map[string]interface{}{"value": []uintptr{uintptr(1)}, "nothing": nil}
	assert.Equal(t, uintptr(1), New(m).Get("value").UintptrSlice()[0])
	assert.Equal(t, uintptr(1), New(m).Get("value").MustUintptrSlice()[0])
	assert.Equal(t, []uintptr(nil), New(m).Get("nothing").UintptrSlice())
	assert.Equal(t, uintptr(1), New(m).Get("nothing").UintptrSlice([]uintptr{uintptr(1)})[0])

	assert.Panics(t, func() {
		New(m).Get("nothing").MustUintptrSlice()
	})

}

func TestIsUintptr(t *testing.T) {

	var o *O

	o = New(uintptr(1))
	assert.True(t, o.IsUintptr())

	o = New([]uintptr{uintptr(1)})
	assert.True(t, o.IsUintptrSlice())

}

func TestFloat32(t *testing.T) {

	m := map[string]interface{}{"value": float32(1), "nothing": nil}
	assert.Equal(t, float32(1), New(m).Get("value").Float32())
	assert.Equal(t, float32(1), New(m).Get("value").MustFloat32())
	assert.Equal(t, float32(0), New(m).Get("nothing").Float32())
	assert.Equal(t, float32(1), New(m).Get("nothing").Float32(1))

	assert.Panics(t, func() {
		New(m).Get("age").MustFloat32()
	})

}

func TestFloat32Slice(t *testing.T) {

	m := map[string]interface{}{"value": []float32{float32(1)}, "nothing": nil}
	assert.Equal(t, float32(1), New(m).Get("value").Float32Slice()[0])
	assert.Equal(t, float32(1), New(m).Get("value").MustFloat32Slice()[0])
	assert.Equal(t, []float32(nil), New(m).Get("nothing").Float32Slice())
	assert.Equal(t, float32(1), New(m).Get("nothing").Float32Slice([]float32{float32(1)})[0])

	assert.Panics(t, func() {
		New(m).Get("nothing").MustFloat32Slice()
	})

}

func TestIsFloat32(t *testing.T) {

	var o *O

	o = New(float32(1))
	assert.True(t, o.IsFloat32())

	o = New([]float32{float32(1)})
	assert.True(t, o.IsFloat32Slice())

}

func TestFloat64(t *testing.T) {

	m := map[string]interface{}{"value": float64(1), "nothing": nil}
	assert.Equal(t, float64(1), New(m).Get("value").Float64())
	assert.Equal(t, float64(1), New(m).Get("value").MustFloat64())
	assert.Equal(t, float64(0), New(m).Get("nothing").Float64())
	assert.Equal(t, float64(1), New(m).Get("nothing").Float64(1))

	assert.Panics(t, func() {
		New(m).Get("age").MustFloat64()
	})

}

func TestFloat64Slice(t *testing.T) {

	m := map[string]interface{}{"value": []float64{float64(1)}, "nothing": nil}
	assert.Equal(t, float64(1), New(m).Get("value").Float64Slice()[0])
	assert.Equal(t, float64(1), New(m).Get("value").MustFloat64Slice()[0])
	assert.Equal(t, []float64(nil), New(m).Get("nothing").Float64Slice())
	assert.Equal(t, float64(1), New(m).Get("nothing").Float64Slice([]float64{float64(1)})[0])

	assert.Panics(t, func() {
		New(m).Get("nothing").MustFloat64Slice()
	})

}

func TestIsFloat64(t *testing.T) {

	var o *O

	o = New(float64(1))
	assert.True(t, o.IsFloat64())

	o = New([]float64{float64(1)})
	assert.True(t, o.IsFloat64Slice())

}

func TestComplex64(t *testing.T) {

	m := map[string]interface{}{"value": complex64(1), "nothing": nil}
	assert.Equal(t, complex64(1), New(m).Get("value").Complex64())
	assert.Equal(t, complex64(1), New(m).Get("value").MustComplex64())
	assert.Equal(t, complex64(0), New(m).Get("nothing").Complex64())
	assert.Equal(t, complex64(1), New(m).Get("nothing").Complex64(1))

	assert.Panics(t, func() {
		New(m).Get("age").MustComplex64()
	})

}

func TestComplex64Slice(t *testing.T) {

	m := map[string]interface{}{"value": []complex64{complex64(1)}, "nothing": nil}
	assert.Equal(t, complex64(1), New(m).Get("value").Complex64Slice()[0])
	assert.Equal(t, complex64(1), New(m).Get("value").MustComplex64Slice()[0])
	assert.Equal(t, []complex64(nil), New(m).Get("nothing").Complex64Slice())
	assert.Equal(t, complex64(1), New(m).Get("nothing").Complex64Slice([]complex64{complex64(1)})[0])

	assert.Panics(t, func() {
		New(m).Get("nothing").MustComplex64Slice()
	})

}

func TestIsComplex64(t *testing.T) {

	var o *O

	o = New(complex64(1))
	assert.True(t, o.IsComplex64())

	o = New([]complex64{complex64(1)})
	assert.True(t, o.IsComplex64Slice())

}

func TestComplex128(t *testing.T) {

	m := map[string]interface{}{"value": complex128(1), "nothing": nil}
	assert.Equal(t, complex128(1), New(m).Get("value").Complex128())
	assert.Equal(t, complex128(1), New(m).Get("value").MustComplex128())
	assert.Equal(t, complex128(0), New(m).Get("nothing").Complex128())
	assert.Equal(t, complex128(1), New(m).Get("nothing").Complex128(1))

	assert.Panics(t, func() {
		New(m).Get("age").MustComplex128()
	})

}

func TestComplex128Slice(t *testing.T) {

	m := map[string]interface{}{"value": []complex128{complex128(1)}, "nothing": nil}
	assert.Equal(t, complex128(1), New(m).Get("value").Complex128Slice()[0])
	assert.Equal(t, complex128(1), New(m).Get("value").MustComplex128Slice()[0])
	assert.Equal(t, []complex128(nil), New(m).Get("nothing").Complex128Slice())
	assert.Equal(t, complex128(1), New(m).Get("nothing").Complex128Slice([]complex128{complex128(1)})[0])

	assert.Panics(t, func() {
		New(m).Get("nothing").MustComplex128Slice()
	})

}

func TestIsComplex128(t *testing.T) {

	var o *O

	o = New(complex128(1))
	assert.True(t, o.IsComplex128())

	o = New([]complex128{complex128(1)})
	assert.True(t, o.IsComplex128Slice())

}