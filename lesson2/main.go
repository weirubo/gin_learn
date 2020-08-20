package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 响应处理
func main() {
	r := gin.Default()
	// String
	// c.String() 第一个参数是code，第二个参数是格式化字符串，第三个开始的若干参数支持任何数据类型
	r.GET("/username", func(c *gin.Context) {
		c.String(http.StatusOK, "用户名：%s", "frank")
	})
	// JSON
	// map
	r.GET("/user", func(c *gin.Context) {
		user := make(map[string]interface{})
		user["name"] = "frank"
		user["email"] = "frank@gmail.com"
		user["age"] = 18

		c.JSON(http.StatusOK, user)

	})

	// gin.H{}
	r.GET("/user1", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"name":  "bob",
			"email": "bob@gmail.com",
			"age":   19,
		})
	})

	// struct
	// 字段名首字母必须大写，否则 c.JSON() 获取不到
	// 如果响应字段的首字母需要小写，可以在 struct 使用 tag 设置
	type User struct {
		Name  string `json:"name"`
		Email string
		Age   int
	}
	r.GET("/user2", func(c *gin.Context) {
		user2 := User{
			Name:  "lucy",
			Email: "lucy@gmail.com",
			Age:   17,
		}
		c.JSON(http.StatusOK, user2)
	})

	// JSONP
	// curl -X GET http://localhost:8081/user3?callback=abc
	// abc({"age":16,"email":"lily@gmail.com","name":"lily"});
	r.GET("/user3", func(c *gin.Context) {
		c.JSONP(http.StatusOK, gin.H{
			"name":  "lily",
			"email": "lily@gmail.com",
			"age":   16,
		})
	})

	r.GET("/user4", func(c *gin.Context) {
		// {"html":"\u003cb\u003ehello\u003c/b\u003e"}
		/* c.JSON(http.StatusOK, gin.H{
			"html": "<b>hello</b>",
		}) */
		// {"html":"<b>hello</b>"}
		c.PureJSON(http.StatusOK, gin.H{
			"html": "<b>hello</b>",
		})
	})

	// r.SecureJsonPrefix(")]}',\n")
	r.GET("/user5", func(c *gin.Context) {
		arr := [...]string{"a", "b", "c"}
		c.SecureJSON(http.StatusOK, arr)
	})

	r.GET("/user6", func(c *gin.Context) {
		c.AsciiJSON(http.StatusOK, gin.H{
			"name":  "张三",
			"email": "zs@gmail.com",
			"age":   20,
		})
	})

	// xml
	r.GET("/user7", func(c *gin.Context) {
		c.XML(http.StatusOK, gin.H{
			"name":  "apple",
			"email": "apple@gmail.com",
		})
	})

	// yaml
	r.GET("/user8", func(c *gin.Context) {
		c.YAML(http.StatusOK, gin.H{
			"name":  "orange",
			"email": "orange@gmailcom",
		})
	})

	_ = r.Run(":8081")
}
