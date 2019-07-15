package benchmark

import (
	"testing"
	"unsafe"
)

func Str2Bytes(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}
func Bytes2Str(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func BenchmarkB1(b *testing.B) {
	b.StopTimer()
	str := "fjdklsjakfjdkljkslajfkldjksajkl";
	var bs []byte = []byte(str);
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		str = Bytes2Str(bs)
		bs = Str2Bytes(str)
	}
}

func BenchmarkB2(b *testing.B) {
	b.StopTimer()
	str := "fjdklsjakfjdkljkslajfkldjksajkl";
	var bs []byte = []byte(str);
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		str = string(bs[:])
		bs = []byte(str[:])
	}
}




