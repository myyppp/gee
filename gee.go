package gee

import (
	"fmt"
	"log"
	"net/http"
)

// HandlerFunc 定义请求用的handler
type HandlerFunc func(http.ResponseWriter, *http.Request)

// Engine 实现ServeHTTP接口
type Engine struct {
	// 路由映射表，key 由请求方法和静态路由地址构成，
	// 例如GET-/、GET-/hello、POST-/hello，
	// 针对相同的路由，如果请求方法不同,可以映射不同的处理方法(Handler)，
	// value 用户映射的处理方法。
	router map[string]HandlerFunc
}

// 构造函数
func New() *Engine {
	return &Engine{
		router: make(map[string]HandlerFunc),
	}
}

func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	key := method + "-" + pattern
	log.Printf("Route %4s - %s", method, pattern)
	engine.router[key] = handler
}

// GET 请求
// 将路由和处理方法注册到映射表 router 中
func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.addRoute("GET", pattern, handler)
}

// POST 请求
func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.addRoute("POST", pattern, handler)
}

// 开启一个http服务
func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}

// 解析请求的路径，查找路由映射表
// 如果找到，执行注册的处理方法
// 找不到，404
func (engine *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	key := r.Method + "-" + r.URL.Path
	if handler, ok := engine.router[key]; ok {
		handler(w, r)
	} else {
		w.WriteHeader(http.StatusNotFound) // 设置返回码
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", r.URL)
	}

}
