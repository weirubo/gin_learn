package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// HTTP 自定义配置

func main() {
	r := gin.Default()

	r.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{})
	})

	s := &http.Server{
		Addr:           ":8081",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1024 * 1024,
	}

	_ = s.ListenAndServe()

	/* r := gin.Default()
	r.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{})
	})
	_ = http.ListenAndServe(":8081", r) */
}
