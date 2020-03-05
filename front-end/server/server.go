package service

import (
	"net/http"
	"net/http/httputil"
	"os"

	"net/url"

	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"github.com/urfave/negroni"
)

// NewServer configures and returns a Server.
func NewServer() *negroni.Negroni {

	formatter := render.New(render.Options{
		IndentJSON: true,
	})

	n := negroni.Classic()
	mx := mux.NewRouter()

	initRoutes(mx, formatter)

	n.UseHandler(mx)
	return n
}

type apiHandle struct {
	host string
	port string
}

func (this *apiHandle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	remote, err := url.Parse("http://" + this.host + ":" + this.port)
	if err != nil {
		panic(err)
	}
	println(remote.Host, remote.Path)
	println(r.URL.Path)
	proxy := httputil.NewSingleHostReverseProxy(remote)
	proxy.ServeHTTP(w, r)
}

type fileNotExistHandle struct {
	host string
	port string
}

func (this *fileNotExistHandle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	println("redirect to self")
	println(r.URL.Path)
	http.Redirect(w, r, "http://"+this.host+":"+this.port+"/index.html", http.StatusFound)
	return
}

func initRoutes(mx *mux.Router, formatter *render.Render) {
	webRoot := os.Getenv("WEBROOT")
	if len(webRoot) == 0 {
		if root, err := os.Getwd(); err != nil {
			panic("Could not retrive working directory")
		} else {
			webRoot = root
			//fmt.Println(root)
		}
	}

	apih := &apiHandle{host: "127.0.0.1", port: "8081"} //后端在本机运行，后端的地址和端口
	//api 代理
	mx.PathPrefix("/api").Handler(apih)
	//非静态资源
	// fneh := &fileNotExistHandle{host: "47.112.254.255", port: "8080"} //服务器公有地址与前端端口
	// mx.PathPrefix("/user").Handler(fneh)
	// mx.PathPrefix("/admin").Handler(fneh)
	//静态资源
	mx.PathPrefix("/").Handler(http.FileServer(http.Dir(webRoot + "/assets/")))
}
