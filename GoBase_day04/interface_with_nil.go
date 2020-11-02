/**
*FileName interface_with_nil.go
*Description
*Creat by LiuYang
*Date 2020/11/1 12:52
 */
package main

import "fmt"

type I2 interface {
	M2()
}

type T2 struct {
	S1 string
}

func (t *T2) M2() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S1)
}

func main() {
	var i1 I2

	var t *T2
	i1 = t
	describe1(i1)
	i1.M2()

	i1 = &T2{"hello"}
	describe1(i1)
	i1.M2()
}

func describe1(i I2) {
	fmt.Printf("(%v, %T)", i, i)
}

/*
即便接口内的具体值为 nil，方法仍然会被 nil 接收者调用。

在一些语言中，这会触发一个空指针异常，

但在 Go 中通常会写一些方法来优雅地处理它（如本例中的 M 方法）。

注意: 保存了 nil 具体值的接口其自身并不为 nil
*/
