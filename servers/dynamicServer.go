package servers

import (
	"net/http"
	"strings"
)

//动态服务器结构
type DynamicServer struct {
	execFunc func(http.ResponseWriter, *http.Request)
}

//判断路由是否属于动态数据请求
func (this *DynamicServer) CheckDynamic(path string) bool {
	path = strings.Trim(path, "/")
	if function, ok := Web[path]; ok {
		this.execFunc = function
		return true
	}
	if function, ok := Api[path]; ok {
		this.execFunc = function
		return true
	}
	return false

}

//判断路由是否属于websocket
func (this *DynamicServer) CheckWebsocket(path string) bool {
	path = strings.Trim(path, "/")
	if function, ok := Ws[path]; ok {
		this.execFunc = function
		return true
	}
	return false
}

//执行动态方法
func (this *DynamicServer) Execute(w http.ResponseWriter, r *http.Request) {
	this.execFunc(w, r)
}

func (this *DynamicServer) getAlias(req string) string {
	if val, ok := Alias[req]; ok {
		return val
	}
	return req
}

//返回一个动态服务器实例
func NewDynamicServer() *DynamicServer {
	return &DynamicServer{}
}
