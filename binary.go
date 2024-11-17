package gulc

import (
	"reflect"
	"unsafe"
)

// LowBits 返回最低位的1
func LowBits(num int) int {
	return num & (-num)
}

// ReverseBits 反转二进制位
func ReverseBits(num int) int {
	num = ((num << 1) & 0xaaaaaaaa) | ((num & 0xaaaaaaaa) >> 1)
	num = ((num << 2) & 0xcccccccc) | ((num & 0xcccccccc) >> 2)
	num = ((num << 4) & 0xf0f0f0f0) | ((num & 0xf0f0f0f0) >> 4)
	num = ((num << 8) & 0xff00ff00) | ((num & 0xff00ff00) >> 8)
	num = ((num << 16) & 0xffff0000) | ((num & 0xffff0000) >> 16)
	return num
}


// HighBits 返回数字的最高位
func HighBits(num int) int {
	return ReverseBits(LowBits(ReverseBits(num)))
}

// LE2Int64以小端的顺序创造data,转换为int64
func LE2Int64(data []byte) uint64 {
	addr := (*reflect.SliceHeader)(unsafe.Pointer(&data)).Data
	int64Addr := (*uint64)(unsafe.Pointer(addr))
	return *int64Addr
}

// Int2Byte以小端方式将num转换为[]byte
func Int2Byte(num int) []byte {
	sliceHeader := reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(&num)),
		Len:  4,
		Cap:  4,
	}
	return *(*[]byte)(unsafe.Pointer(&sliceHeader))
}