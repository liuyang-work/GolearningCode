/**
*FileName map_literals.go
*Description
*Creat by LiuYang
*Date 2020/10/28 13:54
 */
package main

import "fmt"

type Vertex5 struct {
	Lat,Long float64
}

var n = map[string]Vertex5{
	"Bell Labs":Vertex5{40.68433,-74.39967},
	"Google":Vertex5{37.42202, -122.0840},
}

func main() {
	fmt.Println(n)
}

//映射的文法与结构体相似，不过必须有键名。
