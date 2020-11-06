/**
*FileName exercise_fibonacci_closure.go
*Description
*Creat by LiuYang
*Date 2020/11/5 20:38
 */
package main

import (
	"fmt"
)

// 返回一个“返回int的函数”
func fibonacci() func() int {
	res1 := 0
	res2 := 1
	return func() int {
		temp := res1
		res1, res2 = res2, res1+res2
		return temp
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

/*
让我们用函数做些好玩的事情。

实现一个 fibonacci 函数，

它返回一个函数（闭包），

该闭包返回一个斐波纳契数列 `(0, 1, 1, 2, 3, 5, ...)`。
*/
