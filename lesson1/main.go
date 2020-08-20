package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 安装 gin，快速启动
func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	_ = r.Run(":8081")
}
