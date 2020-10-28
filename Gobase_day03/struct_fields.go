/**
*FileName struct_fields.go
*Description
*Creat by LiuYang
*Date 2020/10/28 12:10
 */
package main

import "fmt"

type Vertex1 struct {
	X int
	Y int
}

func main() {
	v := Vertex1{1,2}
	v.X=4
	fmt.Println(v.X)
}

//结构体字段使用点号来访问
