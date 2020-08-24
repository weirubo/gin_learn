package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HTTP 方法

func main() {
	r := gin.Default()
	// 查询
	r.GET("/test", func(c *gin.Context) {
		c.String(http.StatusOK, "get")
	})
	// 新增
	r.POST("/test", func(c *gin.Context) {
		c.String(http.StatusOK, "post")
	})
	// 更新
	r.PUT("/test", func(c *gin.Context) {
		c.String(http.StatusOK, "put")
	})
	// 更新某一个资源的某一项
	r.PATCH("/test", func(c *gin.Context) {
		c.String(http.StatusOK, "patch")
	})
	// 删除
	r.DELETE("/test", func(c *gin.Context) {
		c.String(http.StatusOK, "delete")
	})
	_ = r.Run(":8081")
}
