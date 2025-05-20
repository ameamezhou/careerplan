package sample

import (
	"github.com/ameamezhou/xiawuyue"
	"github.com/ameamezhou/xiawuyue/xlog"
	"github.com/careerplan/handler"
	"github.com/careerplan/pkg/errno"
)

func SampleHandler(c *xiawuyue.Context){
	var err error = nil
	var data interface{} = nil
	var listParam handler.QueryParam

	defer func() {
		if err != nil {
			xlog.Errorf("%s\n", err)
		}
		handler.SendResponse(c, err, data)
	}()

	if e := c.FormUnmarshal(&listParam); e != nil {
		err = errno.New(errno.SampleError, e)
		return
	}
	// 具体逻辑

	data = handler.QueryParam{0}
}
