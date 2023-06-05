package main

import "C"

/*
	"C" をimportすることでbuild時にdllに変換することができる

	例: go build -buildmode=c-shared -o main-cgo.dll export-cgo.go
	上記コマンドによって生成されるファイルは main-cgo.dll と main-cgo.hファイル

	.dll：その名の通りdllファイル
	.h  ：Cで記述されたファイル
*/

// mainがないとbuildできない
func main() {}

/*
	exportしたい関数に
	//export {functionName}
	とコメントを入れる

	※// と exportの間にスペースを入れてしまうと .h ファイルが作成されないため注意
*/

//export addition
func addition(vals []int) int {
	var result int
	for _, v := range vals {
		result += v
	}
	return result
}

//export subtraction
func subtraction(vals []int) int {
	var result int
	for _, v := range vals {
		result -= v
	}
	return result
}

//export multiplication
func multiplication(vals []int) int {
	var result int
	for _, v := range vals {
		result *= v
	}
	return result
}

//export division
func division(a, b float64) float64 {
	if b == 0 {
		panic(`0では割れない`)
	}
	return a / b
}
