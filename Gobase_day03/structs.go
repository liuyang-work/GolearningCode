/**
*FileName structs.go
*Description
*Creat by LiuYang
*Date 2020/10/28 12:07
 */
package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	fmt.Println(Vertex{1,2})
}

//一个结构体（struct）就是一组字段（field）