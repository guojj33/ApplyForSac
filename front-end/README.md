### 前端

#### 生产环境部署

- 打包

  在 front-end-dev 目录下执行进行打包

  ```
  cnpm run build
  ```

  成功后会在 front-end-dev 目录下将生成 dist 文件夹，将其复制到当前目录的 assets 目录下

- 运行

  直接运行

  ```
  go run main.go
  ```

  生成可执行文件再运行

  ```
  go build
  ./front-end.exe
  ```