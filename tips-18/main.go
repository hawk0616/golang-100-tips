package main

import (
	"fmt"
	"math"
)

// 正数のオーバーフローを無視する

// Goには10個の整数型がある
// それぞれの型には固有の範囲がある

func main() {
	var counter int32 = math.MaxInt32
	fmt.Println(counter)
	counter = counter + 1
	fmt.Println(counter)

	// 上記コードはコンパイルされ、panicにはならないが、整数のオーバーフローが発生している
	// このようなオーバーフローは算術演算が与えられた型の範囲を超えた場合に発生する

	// Goではコンパイル時に検出可能なオーバーフローはpanicを引き起こす
	// example: var counter02 int32 = math.MaxInt32 + 1
}

// 加算された時に正数のオーバーフローを検出
func Inc32 (counter int32) int32 {
	if counter == math.MaxInt32 {
		panic("counter overflow")
	}
	return counter + 1
}

// 符号なしの場合も一緒
func IncUint(counter uint32) uint32 {
	if counter == math.MaxUint32 {
		panic("counter overflow")
	}

	return counter + 1
}

// ２つの整数加算時にオーバーフローを検出
func AddInt(a, b int) int {
	if (b < 0 && a > math.MaxInt - b ) || (b > 0 && a < math.MaxInt - b ) {
		panic("integer overflow")
	}
	return a + b
}

// 常山時のオーバーフローを検出
func MultiplyInt(a, b int) int {
	if a == 0 || b == 0 {
		return 0
	}
	result := a * b
	if a == 1 || b == 1 {
		return result
	}
	if a == math.MaxInt || b == math.MaxInt{
		panic("integer overflow")
	}
	// 常山が正数のオバーフローを発生させているかを検出
	if result/a != b {
		panic("integer overflow")
	}
	return result
}