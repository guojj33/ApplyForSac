package main

import (
	"./Models"
	"./Server"
)

const (
	PORT string = "8081"
)

// 服务器
func main() {
	port := PORT

	Models.InitDB()
	defer Models.CloseDB()

	server := Server.NewServer()
	server.Run(":" + port)
}
