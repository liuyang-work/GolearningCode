/**
*FileName methods_pointers_explained.go
*Description
*Creat by LiuYang
*Date 2020/11/1 15:05
 */
package main

import (
	"fmt"
	"math"
)

type Vertex2 struct {
	X, Y float64
}

func Abs1(v Vertex2) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Scale(v *Vertex2, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex2{3, 4}
	Scale(&v, 10)
	fmt.Println(Abs1(v))
}

/*现在我们要把 Abs 和 Scale 方法重写为函数。

同样，我们先试着移除掉第 16 的 *。

你能看出为什么程序的行为改变了吗？要怎样做才能让该示例顺利通过编译？

（若你不确定，继续往下看。）
*/
