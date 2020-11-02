/**
*FileName stringer.go
*Description
*Creat by LiuYang
*Date 2020/11/1 13:37
 */
package main

import (
	"fmt"
)

/*
type Stringer interface {
	String() string
}
*/

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a, z)

}

/*
fmt 包中定义的 Stringer 是最普遍的接口之一。

type Stringer interface {
    String() string
}

Stringer 是一个可以用字符串描述自己的类型。fmt 包（还有很多包）都通过此接口来打印值。
*/
