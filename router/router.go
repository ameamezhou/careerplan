package router

import (
	"github.com/ameamezhou/xiawuyue"
	"net/http"

	"git.woa.com/weworkoss/forecast/handler/bot"
	"git.woa.com/weworkoss/forecast/handler/buildabolish"
	"git.woa.com/weworkoss/forecast/handler/charts"
	"git.woa.com/weworkoss/forecast/handler/index"
	"git.woa.com/weworkoss/forecast/handler/insforecast"
	"git.woa.com/weworkoss/forecast/handler/modulestock"
)

// Handler 路由处理器
type Handler struct {
	Method  string
	Prepare func(x *xiawuyue.Context) (err error)
	Do      func(x *xiawuyue.Context)
	URL     string
	Name    string
}

// HandlerList Handler list
type HandlerList struct {
	public  string
	handles []*Handler
	prepare func(http.ResponseWriter, *http.Request) error
}


// GetHandlerList Handler List
func GetHandlerList() []*Handler {

	handle_list := &HandlerList{
		public: "/public/api",
	}
	return handle_list.handles
}
