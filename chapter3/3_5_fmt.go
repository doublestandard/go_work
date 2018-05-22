package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello Golang!")

	fmt.Println("My", "name", " is", " Taro")

	// printf
	fmt.Printf("数値 = %d \n", 5)

	fmt.Printf("10進数=%d 2進数=%b 8進数=%o 16進数=%x \n", 17, 17, 17, 17)
	// => 10進数=17 2進数=10001 8進数=21 16進数=11

	//パラメータ不足
	fmt.Printf("%d年%d月%d日 \n", 2018, 5)
	// 2018年5月%!d(MISSING)日

	// パラメータ多過ぎ
	fmt.Printf("%d年%d月%d日 \n", 2018, 5, 1, 9)
	// => 2018年5月1日 %!(EXTRA int=9)

	fmt.Printf("\n")

	fmt.Printf("数値=%v 文字列=%v 配列=%v \n", 5, "Go lang", [...]int{1, 2, 3})
	// => 数値=5 文字列=Go lang 配列=[1 2 3]

	// %#はGoのリテラル表現でデータを埋め込む
	fmt.Printf("数値=%#v 文字列=%#v 配列=%#v \n", 5, "Go lang", [...]int{1, 2, 3})
	//=>数値=5 文字列="Go lang" 配列=[3]int{1, 2, 3}

	//%Tはデータの方情報を埋め込む
	fmt.Printf("数値=%T 文字列=%T 配列=%T \n", 5, "Go lang", [...]int{1, 2, 3})
	//=>数値=int 文字列=string 配列=[3]int

	// printlnは改行付き
	fmt.Print("Test!")
	fmt.Println("TEST!")
	//=> Test!TEST! (改行)
}
