package main

import (
	"fmt"
	"math"
)

func PrintfInt() {
	var (
		int_8  int8
		int_16 int16
		int_32 int32
		int_64 int64
		int_   int
	)

	int_8 = math.MaxInt8
	int_16 = math.MaxInt16
	int_32 = math.MaxInt32
	int_64 = math.MaxInt64
	int_ = math.MaxInt64
	fmt.Printf("%T max:%d\n", int_8, int_8)
	fmt.Printf("%T max:%d\n", int_16, int_16)
	fmt.Printf("%T max:%d\n", int_32, int_32)
	fmt.Printf("%T max:%d\n", int_64, int_64)
	fmt.Printf("%T max:%d\n", int_, int_)
}

func PrintfUint() {
	var (
		uint_8  uint8
		uint_16 uint16
		uint_32 uint32
		uint_64 uint64
		uint_   uint
	)

	uint_8 = math.MaxUint8
	uint_16 = math.MaxUint16
	uint_32 = math.MaxUint32
	uint_64 = math.MaxUint64
	uint_ = math.MaxUint64
	fmt.Printf("%T max:%d\n", uint_8, uint_8)
	fmt.Printf("%T max:%d\n", uint_16, uint_16)
	fmt.Printf("%T max:%d\n", uint_32, uint_32)
	fmt.Printf("%T max:%d\n", uint_64, uint_64)
	fmt.Printf("%T max:%d\n", uint_, uint_)
}

func PrintfRuneAndByte() {
	var (
		rune_ rune
		byte_ byte
	)
	rune_ = math.MaxInt32
	byte_ = math.MaxInt8
	fmt.Printf("%T max:%v\n", rune_, rune_)
	fmt.Printf("%T max:%v\n", byte_, byte_)
}

func PirntfFloat() {
	var (
		float_32 float32
		float_64 float64
	)
	float_32 = math.MaxFloat32
	float_64 = math.MaxFloat64
	fmt.Printf("%T max:%f\n", float_32, float_32)
	fmt.Printf("%T max:%f\n", float_64, float_64)
}

func main() {
	PrintfInt()
	PrintfUint()
	PrintfRuneAndByte()
	PirntfFloat()

	// 	int8	Yes	8
	// int16	Yes	16
	// int32	Yes	32
	// int64	Yes	64
	// uint8	No	8
	// uint16	No	16
	// uint32	No	32
	// uint64	No	64
	// int	Yes	等于cpu位数
	// uint	No	等于cpu位数
	// rune	Yes	与 int32 等价
	// byte	No	与 uint8 等价
	// uintptr	No	-

}
