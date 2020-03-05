### 后端

- 需要支持 go 的环境

- 可能还需要安装的包
  ```
	github.com/boltdb/bolt
  
	github.com/urfave/negroni
	github.com/gorilla/mux
	github.com/unrolled/render

	github.com/dgrijalva/jwt-go/request
	github.com/dgrijalva/jwt-go
  ```
  
- 运行
  
  直接运行
  
```shell
  go run main.go
  ```
  
  或者生成可执行文件再运行
  
  ```
  go build
  ./back-end.exe
  ```
  
  
  
- 默认在 ``localhost:8081`` 运行

