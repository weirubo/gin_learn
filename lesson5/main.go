package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// 绑定参数

type user struct {
	Name string `form:"username" json:"username" xml:"username" uri:"username"`
	Age  int    `form:"age" json:"age" xml:"age" uri:"age"`
}

type student struct {
	UserInfo user `form:"userinfo"`
	Score    int  `form:"score" json:"score" xml:"score" uri:"score"`
}

/* func (*gin.Context).ShouldBind(obj interface{}) error
(gin.Context).ShouldBind on pkg.go.dev

ShouldBind checks the Content-Type to select a binding engine automatically, Depending the "Content-Type" header different bindings are used:

    "application/json" --> JSON binding

    "application/xml" --> XML binding

otherwise --> returns an error It parses the request's body as JSON if Content-Type == "application/json" using JSON or XML as a JSON input. It decodes the json payload into the struct specified as a pointer. Like c.Bind() but this method does not set the response status code to 400 and abort if the json is not valid. */

// 要将请求体绑定到结构体中，使用模型绑定。 Gin目前支持JSON、XML、YAML和标准表单值的绑定（foo=bar＆boo=baz）。
// Gin使用 go-playground/validator.v8 进行验证。
// 结构体字段首字母必须大写。
// 使用时，需要在要绑定的所有字段上，设置相应的tag。 例如，使用 JSON 绑定时，设置字段标签为 json:"fieldname"。
// 你也可以指定必须绑定的字段。 如果一个字段的 tag 加上了 binding:"required"，但绑定时是空值, Gin 会报错。
func main() {
	r := gin.Default()
	// Must bind 和 Should bind
	var user1 user
	// Bind()
	// 这些方法属于 MustBindWith 的具体调用。
	// 如果发生绑定错误，则请求终止，并触发 c.AbortWithError(400, err).SetType(ErrorTypeBind)。响应状态码被设置为 400 并且 Content-Type 被设置为 text/plain; charset=utf-8。
	// 如果您在此之后尝试设置响应状态码，Gin会输出日志 [GIN-debug] [WARNING] Headers were already written. Wanted to override status code 400 with 422。

	// BindJSON()

	// BindXML()

	// BindQuery()

	// BindYAML()

	// ShouldBind()
	// Gin 会尝试根据 Content-Type 推断如何绑定。
	// 这些方法属于 ShouldBindWith 的具体调用。
	// 如果发生绑定错误，Gin 会返回错误并由开发者处理错误和请求。
	// 如果是 `GET` 请求，只使用 `Form` 绑定引擎（`query`）。
	// 如果是 `POST` 请求，首先检查 `content-type` 是否为 `JSON` 或 `XML`，然后再使用 `Form`（`form-data`）。
	// 不能多次绑定，c.ShouldBind 使用了 c.Request.Body，不可重用。
	r.GET("/user", func(c *gin.Context) {
		if err := c.ShouldBind(&user1); err != nil {
			c.String(http.StatusBadRequest, "err:%v", err)
			return
		}
		c.JSON(http.StatusOK, user1)
	})

	// ShouldBindJSON()
	/* func (*gin.Context).ShouldBindJSON(obj interface{}) error
	(gin.Context).ShouldBindJSON on pkg.go.dev

	ShouldBindJSON is a shortcut for c.ShouldBindWith(obj, binding.JSON). */
	r.POST("/user1", func(c *gin.Context) {
		if err1 := c.ShouldBindJSON(&user1); err1 != nil {
			c.String(http.StatusBadRequest, "err:%v", err1)
			return
		}
		c.JSON(http.StatusOK, user1)
	})

	// ShouldBindXML()
	/* func (*gin.Context).ShouldBindXML(obj interface{}) error
	(gin.Context).ShouldBindXML on pkg.go.dev

	ShouldBindXML is a shortcut for c.ShouldBindWith(obj, binding.XML). */
	r.POST("/user2", func(c *gin.Context) {
		if err2 := c.ShouldBindXML(&user1); err2 != nil {
			c.String(http.StatusBadRequest, "err:%v", err2)
			return
		}
		c.JSON(http.StatusOK, user1)
	})

	// ShouldBindQuery()
	// ShouldBindQuery 如果 url 查询参数和 post 数据都存在，函数只绑定 url 查询参数而忽略 post 数据。
	/* func (*gin.Context).ShouldBindQuery(obj interface{}) error
	(gin.Context).ShouldBindQuery on pkg.go.dev

	ShouldBindQuery is a shortcut for c.ShouldBindWith(obj, binding.Query). */
	r.POST("/user3", func(c *gin.Context) {
		if err3 := c.ShouldBindQuery(&user1); err3 != nil {
			c.String(http.StatusBadRequest, "err:%v", err3)
			return
		}
		c.JSON(http.StatusOK, user1)
	})

	// ShouldBindYAML()
	/* func (*gin.Context).ShouldBindYAML(obj interface{}) error
	(gin.Context).ShouldBindYAML on pkg.go.dev

	ShouldBindYAML is a shortcut for c.ShouldBindWith(obj, binding.YAML). */
	r.POST("/user4", func(c *gin.Context) {
		if err4 := c.ShouldBindYAML(&user1); err4 != nil {
			c.String(http.StatusBadRequest, "err:%v", err4)
			return
		}
		c.JSON(http.StatusOK, user1)
	})

	// ShouldBindWith()
	/* func (*gin.Context).ShouldBindWith(obj interface{}, b binding.Binding) error
	(gin.Context).ShouldBindWith on pkg.go.dev

	ShouldBindWith binds the passed struct pointer using the specified binding engine. See the binding package. */
	// 如果想要多次绑定，可以使用 ShouldBindWith()，它会在绑定之前将 body 存储到上下文中，但会对性能造成轻微影响。
	// 只有某些格式需要此功能，如 JSON, XML, MsgPack, ProtoBuf。 对于其他格式, 如 Query, Form, FormPost, FormMultipart 可以多次调用 c.ShouldBind() 而不会造成任任何性能损失
	r.POST("/user5", func(c *gin.Context) {
		if err5 := c.ShouldBindWith(&user1, binding.JSON); err5 != nil {
			c.String(http.StatusBadRequest, "err:%v", err5)
			return
		}
		c.JSON(http.StatusOK, user1)
	})

	// ShouldBindUri()
	/* func (*gin.Context).ShouldBindUri(obj interface{}) error
	(gin.Context).ShouldBindUri on pkg.go.dev

	ShouldBindUri binds the passed struct pointer using the specified binding engine. */
	r.GET("/user6/:username/:age", func(c *gin.Context) {
		if err6 := c.ShouldBindUri(&user1); err6 != nil {
			c.String(http.StatusBadRequest, "err:%v", err6)
			return
		}
		c.JSON(http.StatusOK, user1)
	})

	// 绑定数据到嵌套结构体
	r.GET("/demo", func(c *gin.Context) {
		s1 := student{}
		if err7 := c.ShouldBind(&s1); err7 != nil {
			c.String(http.StatusBadRequest, "err:%v", err7)
		}
		c.JSON(http.StatusOK, s1)
	})

	_ = r.Run(":8081")
}
