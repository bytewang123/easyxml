package easyxml

import (
	"reflect"
	"strconv"
	"unsafe"
)

func equalStr(b *[]byte, s string) bool {
	return *(*string)(unsafe.Pointer(b)) == s
}

func parseFloat(b *[]byte) (float64, error) {
	return strconv.ParseFloat(*(*string)(unsafe.Pointer(b)), 64)
}

// A hack until issue golang/go#2632 is fixed.
// See: https://github.com/golang/go/issues/2632
func bytesToString(b *[]byte) string {
	return *(*string)(unsafe.Pointer(b))
}

func StringToBytes(s string) []byte {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := reflect.SliceHeader{
		Data: sh.Data,
		Len:  sh.Len,
		Cap:  sh.Len,
	}
	return *(*[]byte)(unsafe.Pointer(&bh))
}
