package utils

import "github.com/gin-gonic/gin"

func HttpResponse(code int, msg string, data interface{}) gin.H {
	return gin.H{
		"code": code,
		"msg":  msg,
		"data": data,
	}
}
