/**
*FileName map_literals_continued.go
*Description
*Creat by LiuYang
*Date 2020/10/28 14:00
 */
package main

import "fmt"

type Vertex6 struct {
	Lat, Long float64
}

var (
	m1 = map[string]Vertex6{
		"Bell Labs": { 40.68433,  -74.39967},
		"Google":    {37.42202, -122.08408},
	}
)

func main() {
	fmt.Println(m1)
}

/*若顶级类型只是一个类型名，你可以在文法的元素中省略它。*/