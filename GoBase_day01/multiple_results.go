/*
* @Author: liuyang
* @Date:   2020-10-21 18:29:16
* @Last Modified by:   liuyang
* @Last Modified time: 2020-10-21 18:32:42
*/
package main

import "fmt"

func swap(x,y string) (string,string) {
	return y,x
}

func main() {
	a,b := swap("hello","world")
	fmt.Println(a,b)
}

/*

函数可以返回任意数量的返回值。

swap 函数返回了两个字符串。

*/