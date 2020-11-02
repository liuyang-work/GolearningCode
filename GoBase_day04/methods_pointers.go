/**
*FileName methods_pointers.go
*Description
*Creat by LiuYang
*Date 2020/11/1 14:59
 */
package main

import (
	"fmt"
	"math"
)

type Vertex1 struct {
	X, Y float64
}

func (v Vertex1) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex1) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex1{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())
}
