package gee

import (
	"log"
	"net/http"
)

// HandlerFunc 定义请求用的handler
type HandlerFunc func(*Context)

// Engine 实现ServeHTTP接口
type Engine struct {
	// 路由映射表，key 由请求方法和静态路由地址构成，
	// 例如GET-/、GET-/hello、POST-/hello，
	// 针对相同的路由，如果请求方法不同,可以映射不同的处理方法(Handler)，
	// value 用户映射的处理方法。
	router *router
}

// 构造函数
func New() *Engine {
	return &Engine{
		router: newRouter(),
	}
}

func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	log.Printf("Route %4s - %s", method, pattern)
	engine.router.addRoute(method, pattern, handler)
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

// 接管所有的 HTTP 请求
// 解析请求的路径，查找路由映射表
// 如果找到，执行注册的处理方法
// 找不到，404
func (engine *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := newContent(w, r)
	engine.router.handle(c)
}
