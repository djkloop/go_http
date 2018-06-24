package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)


func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()

	// post
	router.POST("/user", CreaterUser)

	// 带参数的post请求
	router.POST("/user/:user_name", Login)
	return router
}

func main() {
	r := RegisterHandlers()

	http.ListenAndServe(":8999", r)
}
