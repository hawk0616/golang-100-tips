package main

import "fmt"

// GOではfloat32とfloat64の2種類の浮動小数点数型がある
// 浮動小数点は小数値を表現できないという整数型の制約を解消するために導入された

func main() {
	var n float32 = 1.0001
	fmt.Println(n* n)

	var a float64
	positiveInf := 1 / a
	negativeInf := -1 / a
	nan := a / a
	fmt.Println(positiveInf, negativeInf, nan)

	// 10進から浮動小数点数への変換は精度を失う

	f1 := f1(100)
	f2 := f2(100)

	fmt.Println(f1, f2)
}

func f1(n int) float64 {
	result := 10_000.
	for i := 0; i < n; i++ {
		result += 1.0001
	}

	return result
}

func f2(n int) float64 {
	result := 0.
	for i := 0; i < n; i++ {
		result += 1.0001
	}

	return result + 10_000.
}