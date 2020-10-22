/**
*FileName defer_multi.go
*Description
*Creat by LiuYang
*Date 2020/10/22 21:21
 */
package main

import "fmt"

func main() {
	fmt.Println("counting")

	for i := 0; i<10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}

//推迟的函数调用会被压入一个栈中。

//当外层函数返回时，被推迟的函数会按照后进先出的顺序调用。
