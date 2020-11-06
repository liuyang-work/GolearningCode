/**
*FileName exercise_slices.go
*Description
*Creat by LiuYang
*Date 2020/11/4 21:34
 */
package main

/*

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	var sli [][]uint8			//新建一个切片
	//sli :=make([][]unit8)
	for i := 0; i < dy; i++ {
		var one_line []uint8
		for j := 0; j < dx; j++ {
			one_line = append(one_line, (uint8)(i*j))		//有趣的函数有(x+y)/2, x*y, x^y, x*log(y) 和 x%(y+1)
		}
		sli = append(sli, one_line)
	}
	return sli
}

func main() {
	pic.Show(Pic)
}
*/
/*
实现 Pic。它应当返回一个长度为 dy 的切片，

其中每个元素是一个长度为 dx，元素类型为 uint8 的切片。

当你运行此程序时，它会将每个整数解释为灰度值（好吧，其实是蓝度值）并显示它所对应的图像。

图像的选择由你来定。几个有趣的函数包括 (x+y)/2, x*y, x^y, x*log(y) 和 x%(y+1)。

（提示：需要使用循环来分配 [][]uint8 中的每个 []uint8；

请使用 uint8(intValue) 在类型之间转换；你可能会用到 math 包中的函数。）
*/
