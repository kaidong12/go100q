package gin_demo

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func My_Handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("usersession", "userid-1")
		c.Next()
		//if xxxx {
		//	c.Next()
		//}
		c.Abort()
	}
}

func Middleware_demo() {
	// 创建服务
	r := gin.Default()

	// 注册中间件
	r.Use(My_Handler())

	// 获取client json data并传递给前端
	r.GET("/user/inf", My_Handler(), func(c *gin.Context) {

		usersession := c.MustGet("usersession").(string)
		log.Println("===================", usersession)

		username := c.PostForm("username")
		password := c.PostForm("password")
		email := c.PostForm("useremail")

		m := make(map[string]interface{})
		//m["ID"] = 111
		m["Name"] = username
		m["Password"] = password
		m["Email"] = email
		m["Response"] = "testing"

		//c.JSON(http.StatusOK, m)
		c.HTML(http.StatusOK, "pass_data.html", m)

	})
}
