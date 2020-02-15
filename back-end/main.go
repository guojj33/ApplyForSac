package main

import (
	"./Models"
	"./Server"
)

// 命令行程序
// func main() {
// 	Models.InitDB()
// 	defer Models.CloseDB()
// 	UI.MainMenu()
// }

const (
	PORT string = "8080"
)

// 服务器
func main() {
	port := PORT

	Models.InitDB()
	defer Models.CloseDB()

	server := Server.NewServer()
	server.Run(":" + port)
}
