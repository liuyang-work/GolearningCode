/*
* @Author: liuyang
* @Date:   2020-10-21 20:14:31
* @Last Modified by:   liuyang
* @Last Modified time: 2020-10-21 20:17:17
*/
package main

import "fmt"

const Pi = 3.14

func main() {
	const World = "世界"
	fmt.Println("Hello",World)
	fmt.Println("Happy",Pi,"day")

	const Truth = true
	fmt.Println("Go rules?",Truth)
}

/*
常量的声明与变量类似，只不过是使用 const 关键字。

常量可以是字符、字符串、布尔值或数值。

常量不能用 := 语法声明。
*/