package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// cookie 读写

/* func (*gin.Context).Cookie(name string) (string, error)
(gin.Context).Cookie on pkg.go.dev

Cookie returns the named cookie provided in the request or ErrNoCookie if not found.
And return the named cookie is unescaped.
If multiple cookies match the given name, only one cookie will be returned. */

/* func (*gin.Context).SetCookie(name string, value string, maxAge int, path string, domain string, secure bool, httpOnly bool)
(gin.Context).SetCookie on pkg.go.dev

SetCookie adds a Set-Cookie header to the ResponseWriter's headers.
The provided cookie must have a valid Name.
Invalid cookies may be silently dropped. */
func main() {
	r := gin.Default()
	r.GET("/cookie", func(c *gin.Context) {
		// 读取 cookie
		val, err := c.Cookie("username4")
		if err != nil {
			// cookie 不存在，设置 cookie，过期时间 10s
			c.SetCookie("username4", "lucy", 10, "/", "localhost", false, true)
		}
		fmt.Println(val)
	})
	_ = r.Run(":8081")
}
