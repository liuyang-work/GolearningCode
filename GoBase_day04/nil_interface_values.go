/**
*FileName nil_interface_values.go
*Description
*Creat by LiuYang
*Date 2020/11/1 13:03
 */
package main

import "fmt"

type I3 interface {
	M3()
}

func main() {
	var i2 I3
	describe2(i2)
	i2.M3()
}

func describe2(i I3) {
	fmt.Printf("(%v, %T)\n", i, i)
}

/*
nil 接口值既不保存值也不保存具体类型。

为 nil 接口调用方法会产生运行时错误，

因为接口的元组内并未包含能够指明该调用哪个具体方法的类型。
*/
