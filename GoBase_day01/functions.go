/*
* @Author: liuyang
* @Date:   2020-10-21 18:12:51
* @Last Modified by:   liuyang
* @Last Modified time: 2020-10-21 18:20:32
*/
package main

import "fmt"

func add(x int,y int) int {
	return x+y
}

func main() {
	fmt.Println("The add number is ",add(4,5))
}

/*
函数可以没有参数或接受多个参数。

在本例中，add 接受两个 int 类型的参数。

注意类型在变量名 之后。

*/