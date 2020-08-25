package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 接收参数，读取 reader

func main() {
	r := gin.Default()
	// 表单提交
	/* func (*gin.Context).Param(key string) string
	(gin.Context).Param on pkg.go.dev

	Param returns the value of the URL param. It is a shortcut for c.Params.ByName(key)

	    router.GET("/user/:id", func(c *gin.Context) {

	     // a GET request to /user/john

	     id := c.Param("id") // id == "john"

	    }) */
	// multipart/form-data
	// application/x-www-form-urlencoded
	// c.PostForm()
	/* func (*gin.Context).PostForm(key string) string
	(gin.Context).PostForm on pkg.go.dev

	PostForm returns the specified key from a POST urlencoded form or multipart form when it exists, otherwise it returns an empty string `("")`. */

	// c.DefaultPostForm()
	/* func (*gin.Context).DefaultPostForm(key string, defaultValue string) string
	(gin.Context).DefaultPostForm on pkg.go.dev

	DefaultPostForm returns the specified key from a POST urlencoded form or multipart form when it exists, otherwise it returns the specified defaultValue string. See: PostForm() and GetPostForm() for further information. */

	/* func (*gin.Context).Query(key string) string
	(gin.Context).Query on pkg.go.dev

	Query returns the keyed url query value if it exists, otherwise it returns an empty string `("")`. It is shortcut for `c.Request.URL.Query().Get(key)`

	     GET /path?id=1234&name=Manu&value=

	     c.Query("id") == "1234"

	     c.Query("name") == "Manu"

	     c.Query("value") == ""

		 c.Query("wtf") == "" */

	/* func (*gin.Context).DefaultQuery(key string, defaultValue string) string
	 (gin.Context).DefaultQuery on pkg.go.dev

	 DefaultQuery returns the keyed url query value if it exists, otherwise it returns the specified defaultValue string. See: Query() and GetQuery() for further information.

		 GET /?name=Manu&lastname=

		 c.DefaultQuery("name", "unknown") == "Manu"

		 c.DefaultQuery("id", "none") == "none"

		 c.DefaultQuery("lastname", "none") == "" */

	/* func (*gin.Context).QueryMap(key string) map[string]string
	(gin.Context).QueryMap on pkg.go.dev

	QueryMap returns a map for a given query key. */

	/* func (*gin.Context).PostFormMap(key string) map[string]string
	(gin.Context).PostFormMap on pkg.go.dev

	PostFormMap returns a map for a given form key. */

	/* func (*gin.Context).DataFromReader(code int, contentLength int64, contentType string, reader io.Reader, extraHeaders map[string]string)
	(gin.Context).DataFromReader on pkg.go.dev

	DataFromReader writes the specified reader into the body stream and updates the HTTP code. */
	r.POST("/:version/user", func(c *gin.Context) {
		version := c.Param("version")
		name := c.PostForm("name")
		age := c.DefaultPostForm("age", "18")
		email := c.Query("email")
		tel := c.Request.URL.Query().Get("tel")
		sex := c.DefaultQuery("sex", "boy")
		score := c.QueryMap("score")
		level := c.PostFormMap("level")

		c.JSON(http.StatusOK, gin.H{
			"version": version,
			"name":    name,
			"age":     age,
			"email":   email,
			"tel":     tel,
			"sex":     sex,
			"score":   score,
			"level":   level,
		})
	})

	// 从 render 读取数据
	r.GET("/avatar", func(c *gin.Context) {
		resp, err := http.Get("http://dingyue.ws.126.net/2020/0823/ef9c1c17j00qfhow6001cd000hs00lbp.jpg")
		if err != nil || resp.StatusCode != http.StatusOK {
			c.Status(http.StatusServiceUnavailable)
			return
		}
		reader := resp.Body
		contentLength := resp.ContentLength
		contentType := resp.Header.Get("Content-Type")

		extraHeaders := map[string]string{
			"Content-Disposition": `attachment; filename="yangchaoyue.jpg"`,
		}

		c.DataFromReader(http.StatusOK, contentLength, contentType, reader, extraHeaders)

	})
	_ = r.Run(":8081")
}
