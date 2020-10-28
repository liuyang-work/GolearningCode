/*
* @Author: liuyang
* @Date:   2020-10-21 18:08:23
* @Last Modified by:   liuyang
* @Last Modified time: 2020-10-21 18:19:16
*/
package main

import (
	"fmt"
	"math"
)

func  main()  {
	fmt.Print("Now you have %g problems.\n",math.Sqrt(7))
}

/*此代码用圆括号组合了导入，这是“分组”形式的导入语句。

当然你也可以编写多个导入语句，例如：

import "fmt"

import "math"

不过使用分组导入语句是更好的形式。*/