package mygin

import (
	"fmt"
	"net/http"
	"path"
	"sync"
)

//上下文context 简单的他添加response request engine指针 isabort就可以支持最简单的流程
type Context struct {
	Request       *http.Request
	ResponseWrite http.ResponseWriter
	engine        *Engine
	isAbort       bool
}

type HandlerFun func(ctx *Context)

type HandlerList []HandlerFun

type Engine struct {
	RouterGroup
	Handlers []HandlerFun
	router   map[string]HandlerList
	pool     sync.Pool
}

type Message struct {
	Message string
}

type IRouter interface {
	Use(...HandlerFun) IRouter
	GET(string, ...HandlerFun) IRouter
	Group(string, ...HandlerFun) *RouterGroup
}

type RouterGroup struct {
	Handlers []HandlerFun
	engine   *Engine
	basePath string
}

func NewEngine() (*Engine) {
	en := new(Engine)
	en.router = make(map[string]HandlerList)
	en.pool.New = func() interface{} {
		return en.allocateContext()
	}
	en.RouterGroup = RouterGroup{
		basePath: "/",
		Handlers: nil,
		engine:   en,
	}

	return en
}

func (engine *Engine) Run(addr string) (err error) {
	fmt.Println("Listening and serving HTTP on", addr)
	err = http.ListenAndServe(addr, engine)
	return
}

//继承http包中的handler接口，在run中即可传入engine
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := engine.pool.Get().(*Context)
	c.ResponseWrite = w
	c.Request = req
	engine.handleHTTPRequest(c)

	engine.pool.Put(c)
}

//客户端请求之后具体执行的函数 之前文章所说的获取所有handler 一个个执行
// 这里简单用了for循环 判断isabort属性来判断是否停止
func (engine *Engine) handleHTTPRequest(c *Context) {
	httpMethod := c.Request.Method
	path := c.Request.URL.Path

	if handlers, ok := engine.router[httpMethod+"^"+path]; ok {
		for _, fu := range handlers {
			fu(c)
			if c.isAbort {
				return
			}
		}
	}
}

func (engine *Engine) allocateContext() *Context {
	return &Context{engine: engine}
}

func (engine *Engine) addRoute(httpMethod, absolutePath string, handlers HandlerList) {
	engine.router[httpMethod+"^"+absolutePath] = handlers
}

//添加group方法 设置group的basepath 和handler
func (routerGroup *RouterGroup) Group(path string, handlers ...HandlerFun) *RouterGroup {
	rg := RouterGroup{}
	rg.Handlers = routerGroup.CombineHandlers(handlers)
	rg.basePath = path
	rg.engine = routerGroup.engine

	return &rg
}

func (routerGroup *RouterGroup) Use(handlers ...HandlerFun) IRouter {
	routerGroup.Handlers = append(routerGroup.Handlers, handlers...)
	return routerGroup
}

func (group *RouterGroup) calculateAbsolutePath(relativePath string) string {
	return joinPaths(group.basePath, relativePath)
}

func joinPaths(absolutePath, relativePath string) string {
	if relativePath == "" {
		return absolutePath
	}

	finalPath := path.Join(absolutePath, relativePath)

	appendSlash := lastChar(relativePath) == '/' && lastChar(finalPath) != '/'
	if appendSlash {
		return finalPath + "/"
	}

	return finalPath
}

//工具方法 获取字符串最后一个字符
func lastChar(str string) uint8 {
	if str == "" {
		panic("The length of the string can't be 0")
	}

	return str[len(str)-1]
}

//计算路径合并handler 然后添加到map中
func (group *RouterGroup) handle(httpMethod, relativePath string, handlers HandlerList) IRouter {
	absolutePath := group.calculateAbsolutePath(relativePath)
	handlers = group.CombineHandlers(handlers)
	group.engine.addRoute(httpMethod, absolutePath, handlers)

	return group
}

//合并handler 之后返回
func (group *RouterGroup) CombineHandlers(handlers HandlerList) HandlerList {
	finalSize := len(group.Handlers) + len(handlers)
	mergedHandler := make(HandlerList, finalSize)
	copy(mergedHandler, group.Handlers)
	copy(mergedHandler[len(group.Handlers):], handlers)
	return mergedHandler
}

//添加get method路由
func (group *RouterGroup) GET(path string, handlers ...HandlerFun) (IRouter) {
	group.handle("GET", path, handlers)

	return group
}
