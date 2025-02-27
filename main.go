package main

import (

	"gin-ranking/router"
)

func main() {
	//进入路由
	r := router.Router()

	//服务端口
	r.Run(":9999")
}
