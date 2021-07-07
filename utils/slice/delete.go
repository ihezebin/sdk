package slice

import (
	"reflect"
)

// By default, the first element matched when traversing slices is deleted;
// If options has a value of true, all matching elements are deleted.
func Delete(s interface{}, elem interface{}, options ...bool) interface{} {
	sV := reflect.ValueOf(s)
	if sV.Kind() != reflect.Slice {
		panic("The first parameter s must be of type slice")
	}
	sLen := sV.Len()
	for i := 0; i < sLen; i++ {
		if sV.Index(i).Interface() == reflect.ValueOf(elem).Interface() {
			sV = reflect.AppendSlice(sV.Slice(0, i), sV.Slice(i+1, sV.Len()))
			if len(options) == 0 || !options[0] {
				return sV.Interface()
			}
			i--
			sLen--
		}
	}
	return sV.Interface()
}

func DeleteInt(s []int, elem int, options ...bool) []int {
	return Delete(s, elem, options...).([]int)
}

func DeleteInt8(s []int8, elem int8, options ...bool) []int8 {
	return Delete(s, elem, options...).([]int8)
}

func DeleteInt16(s []int16, elem int16, options ...bool) []int16 {
	return Delete(s, elem, options...).([]int16)
}

func DeleteInt32(s []int32, elem int32, options ...bool) []int32 {
	return Delete(s, elem, options...).([]int32)
}

func DeleteInt64(s []int64, elem int64, options ...bool) []int64 {
	return Delete(s, elem, options...).([]int64)
}

func DeleteUint8(s []uint8, elem uint8, options ...bool) []uint8 {
	return Delete(s, elem, options...).([]uint8)
}

func DeleteUint16(s []uint16, elem uint16, options ...bool) []uint16 {
	return Delete(s, elem, options...).([]uint16)
}

func DeleteUint32(s []uint32, elem uint32, options ...bool) []uint32 {
	return Delete(s, elem, options...).([]uint32)
}

func DeleteUint64(s []uint64, elem uint64, options ...bool) []uint64 {
	return Delete(s, elem, options...).([]uint64)
}

func DeleteFloat32(s []float32, elem float32, options ...bool) []float32 {
	return Delete(s, elem, options...).([]float32)
}

func DeleteFloat64(s []float64, elem float64, options ...bool) []float64 {
	return Delete(s, elem, options...).([]float64)
}

func DeleteBool(s []bool, elem bool, options ...bool) []bool {
	return Delete(s, elem, options...).([]bool)
}

func DeleteString(s []string, elem string, options ...bool) []string {
	return Delete(s, elem, options...).([]string)
}

func DeleteByte(s []byte, elem byte, options ...bool) []byte {
	return Delete(s, elem, options...).([]byte)
}
func DeleteRune(s []rune, elem rune, options ...bool) []rune {
	return Delete(s, elem, options...).([]rune)
}
