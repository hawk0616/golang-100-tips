package main

import "fmt"

// スライス
// スライスは配列で裏付けられ、スライスのデータ構造は配列のデータ構造に裏付けられる

 func main() {
	s := make([]int, 3, 6)
	fmt.Println(s)
	s[0] = 23
	fmt.Println(s)
	// 長さの範囲外の要素へのアクセスはすでにメモリが確保されているのに禁止されている

	// スライスの残りの領域の使い方
	// append
	s = append(s, 3)
	fmt.Println(s)
	// append によって長さが追加差あれた

	// 配列は固定サイズの構造体
	// Goでは要素がいっぱいになったスライスに要素を追加しようとすると内部的に配列をもう一つ作成し、容量を二倍にして操作分を追加する
	// Goではスライスは要素数が1024になるまで2倍でふえ、それ以降は25%ずつ増えていく

	// スライスは基底配列を参照する
	// 以前の規程配列が参照されなくなった場合、ヒープ上二割与えられていれば、GCにより解放される

	s2 := make([]int, 3, 6)
	s3 := s2[1:3]
	// 上記はともに規程配列を参照している
	// 故にs3はs2とは異なるインデックスから始まり、長さ２で容量が5のスライス

	s3 = append(s3, 2)
	fmt.Println(s2)
	fmt.Println(s3)
	// 上記では同じ規程配列を参照しているが、s3のみ更新される
	fmt.Println(len(s3))

	// また、規程配列の容量を超えてs3に要素を追加し続けると、別の規程配列を作成することになる
	s3 = append(s3, 4)
	s3 = append(s3, 4)
	s3 = append(s3, 4)
	s3 = append(s3, 4)

	fmt.Println(len(s3))
	fmt.Println(s3)
}