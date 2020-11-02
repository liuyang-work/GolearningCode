/**
*FileName interface_values.go
*Description
*Creat by LiuYang
*Date 2020/11/1 16:14
 */
package main

import (
	"fmt"
	"math"
)

type I1 interface {
	MM2()
}

type T1 struct {
	S string
}

func (t *T1) MM2() {
	fmt.Println(t.S)
}

type F float64

func (f F) MM2() {
	fmt.Println(f)
}

func main() {
	var i I1

	i = &T1{"Hello"}
	//fmt.Println(i)
	describe(i)
	i.MM2()

	i = F(math.Pi)
	describe(i)
	i.MM2()
	//fmt.Println(i)
}

func describe(i1 I3) {
	fmt.Printf("(%v, %T)\n", i1, i1)
}

/*
接口也是值。它们可以像其它值一样传递。

接口值可以用作函数的参数或返回值。

在内部，接口值可以看做包含值和具体类型的元组：

(value, type)

接口值保存了一个具体底层类型的具体值。

接口值调用方法时会执行其底层类型的同名方法。
*/
