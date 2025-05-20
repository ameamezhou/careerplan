package router

import (
	"github.com/ameamezhou/xiawuyue"
	"github.com/ameamezhou/xiawuyue/xlog"
	"github.com/careerplan/handler"
	"github.com/careerplan/handler/sample"
)

// Handler 路由处理器
type Handler struct {
	Method  string
	Prepare func(x *xiawuyue.Context) (err error) // 单个函数需要使用的中间件
	Do      func(x *xiawuyue.Context)
	URL     string
}

func setGroup(g *xiawuyue.RouterGroup, h Handler) {
	var hd xiawuyue.HandlerFunc
	if h.Prepare != nil {
		hd = func(c *xiawuyue.Context) {
			err := h.Prepare(c)
			if err != nil {
				xlog.Error(err)
				handler.SendResponse(c, err, struct{}{})
				return
			}
			h.Do(c)
		}
	} else {
		hd = h.Do
	}
	if h.Method == "GET" {
		g.GET(h.URL, hd)
	} else if h.Method == "POST" {
		g.POST(h.URL, hd)
	} else {
		xlog.Errorf("%s add fail! dont support %s method, please use GET or POST!", h.URL, h.Method)
	}
}

// GetHandlerList Handler List
func SetHandlerList(x *xiawuyue.Xia) {
	g := x.Group("/public/api")
	// 如果某个路由组需要使用中间件直接用  g.Use()  即可
	setGroup(g, Handler{"GET", nil, sample.SampleHandler,"/simple"})
}
