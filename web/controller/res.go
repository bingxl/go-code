// 一些用于处理 http 返回的函数
package controller

import (
	"gin-sites/service"
	"net/http"
)

// 请求参数错误
func resBad(msg string) (int, service.ResT) {
	return http.StatusBadRequest,
		service.ResT{
			Code:   http.StatusBadRequest,
			Status: 0,
			Msg:    msg,
			Data:   nil,
		}
}

// 请求成功
func resOK() (int, service.ResT) {
	return http.StatusOK,
		service.ResT{
			Code:   http.StatusOK,
			Status: 1,
			Msg:    "请求成功",
			Data:   nil,
		}
}

// 通过 c.Status 来设置 http status
func res(c service.ResT) (int, service.ResT) {
	httpCode := http.StatusOK
	if c.Status == 0 {
		httpCode = http.StatusInternalServerError
	}

	return httpCode, c
}
