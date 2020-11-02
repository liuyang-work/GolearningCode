/**
*FileName interfaces.go
*Description
*Creat by LiuYang
*Date 2020/11/1 15:25
 */
package main

import (
	"fmt"
	"math"
)

type Abuser interface {
	Abs() float64
}

func main() {
	var a Abuser
	f := MyFloat64(-math.Sqrt(2))
	v := Vertex6{3, 4}

	a = f  //a MyFloat 实现了 Abuser
	a = &v //a *Vertex6 实现了 Abuser

	//下面一行，v是一个Vertex6（而不是*Vertex6）
	//所以没有实现 Abuser
	a = v

	fmt.Println(a.Abs())
}

type MyFloat64 float64

func (f MyFloat64) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex6 struct {
	X, Y float64
}

/*func (v *Vertex6) Abs() float64 {
	return math.Sqrt(v.X*v.X+v.Y*v.Y)
}*/

func (v Vertex6) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

/*接口类型 是由一组方法签名定义的集合。

接口类型的变量可以保存任何实现了这些方法的值。

注意: 示例代码的 22 行存在一个错误。

由于 Abs 方法只为 *Vertex （指针类型）定义，

因此 Vertex（值类型）并未实现 Abuser。
*/
