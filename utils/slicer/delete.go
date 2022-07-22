package slicer

import (
	"reflect"
)

// Delete By default, the first element matched when traversing slices is deleted;
// If options has a value of true, all matching elements are deleted.
func Delete(s interface{}, elem interface{}) interface{} {
	sV := reflect.ValueOf(s)
	if sV.Kind() != reflect.Slice {
		panic("The first parameter s must be of type slicer")
	}
	sLen := sV.Len()
	for i := 0; i < sLen; i++ {
		if sV.Index(i).Interface() == reflect.ValueOf(elem).Interface() {
			sV = reflect.AppendSlice(sV.Slice(0, i), sV.Slice(i+1, sV.Len()))
			return sV.Interface()
		}
	}
	return sV.Interface()
}

func DeleteAll(s interface{}, elem interface{}) interface{} {
	sV := reflect.ValueOf(s)
	if sV.Kind() != reflect.Slice {
		panic("The first parameter s must be of type slicer")
	}
	sLen := sV.Len()
	for i := 0; i < sLen; i++ {
		if sV.Index(i).Interface() == reflect.ValueOf(elem).Interface() {
			sV = reflect.AppendSlice(sV.Slice(0, i), sV.Slice(i+1, sV.Len()))
			i--
			sLen--
		}
	}
	return sV.Interface()
}
