/**
*FileName web_servers.go
*Description
*Creat by LiuYang
*Date 2020/11/2 12:02
 */
package main

import (
	"fmt"
	"log"
	"net/http"
)

type Hello struct{}

func (h Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "Hello!")
}

func main() {
	var h Hello
	err1 := http.ListenAndServe("localhost:4000", h)
	if err1 != nil {
		log.Fatal(err1)
	}
}

/*
包 http 通过任何实现了 http.Handler 的值来响应 HTTP 请求：

package http

type Handler interface {
    ServeHTTP(w ResponseWriter, r *Request)
}
在这个例子中，类型 Hello 实现了 `http.Handler`。

访问 http://localhost:4000/ 会看到来自程序的问候。
*/
