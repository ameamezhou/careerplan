package handler

import (
	"github.com/ameamezhou/xiawuyue"
	"github.com/careerplan/pkg/errno"
	"net/http"
)

type QueryParam struct {
	Data 	interface{} `json:"data" form:"data"`
}

// SendResponse response http requests
func SendResponse(c *xiawuyue.Context, err error, data interface{}) {
	var code int
	var msg string
	var status int
	if err == nil || err == (*errno.Errno)(nil) {
		code, msg = errno.DecodeErr(errno.OK)
		status = http.StatusOK
	} else {
		code, msg = errno.DecodeErr(err)
		if code > 0 && code <= 20000 {
			// server error
			status = http.StatusInternalServerError
		} else {
			if code == 20318 {
				status = http.StatusOK
			} else {
				// client error
				status = http.StatusBadRequest
			}
		}
	}

	c.JSON(status, xiawuyue.ResponseXia{
		Code:    code,
		Message: msg,
		Data:    data,
	})
}

// Tpl 读取template的缓存
var Tpl = new(xiawuyue.BuildTemplate)

// GetMainPage return template page
func GetMainPage(c *xiawuyue.Context) {
	data := xiawuyue.ResponseXia{
		Code:    0,
		Message: "ok",
	}
	c.WriteTpl(Tpl, "main.html", data)
}

