package gin_demo

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetJsonInRequests() {
	// 创建服务
	r := gin.Default()

	// load page
	r.LoadHTMLGlob("examples/gin_demo/templates/*")

	// 获取client json data并传递给前端
	r.POST("/postjson", func(c *gin.Context) {

		// request.body
		body, _ := c.GetRawData()

		var m map[string]interface{}
		_ = json.Unmarshal(body, &m)
		m["Response"] = "testing"

		//c.JSON(http.StatusOK, m)
		c.HTML(http.StatusOK, "pass_data.html", m)
	})

	//r.Run() // 默认监听并在 0.0.0.0:8080 上启动服务
	r.Run(":8090") // 默认监听并在 0.0.0.0:8080 上启动服务

}

func FormDemo_submit_form_data() {
	// 创建服务
	r := gin.Default()

	// load page
	r.LoadHTMLGlob("examples/gin_demo/templates/*")

	// 返回静态页面
	r.GET("/index", func(c *gin.Context) {
		//c.JSON(http.StatusOK, gin.H{
		//	"message": "pong",
		//})
		c.HTML(http.StatusOK, "index.html", gin.H{
			"msg": "gin框架后端 FormDemo 传过来的参数",
		})

	})

	// 获取client json data并传递给前端
	r.POST("/user/add", func(c *gin.Context) {

		username := c.PostForm("username")
		password := c.PostForm("password")
		email := c.PostForm("useremail")

		c.JSON(http.StatusOK, gin.H{
			"msg":      "ok",
			"username": username,
			"password": password,
			"email":    email,
		})

	})

	//r.Run() // 默认监听并在 0.0.0.0:8080 上启动服务
	r.Run(":8090") // 默认监听并在 0.0.0.0:8080 上启动服务

}

func FormDemo_submit_form_to_new_page() {
	// 创建服务
	r := gin.Default()

	// load page
	r.LoadHTMLGlob("examples/gin_demo/templates/*")

	// 返回静态页面
	r.GET("/index", func(c *gin.Context) {
		//c.JSON(http.StatusOK, gin.H{
		//	"message": "pong",
		//})
		c.HTML(http.StatusOK, "index.html", gin.H{
			"msg": "gin框架后端 FormDemo 传过来的参数",
		})

	})

	// 获取client json data并传递给前端
	r.POST("/user/add", func(c *gin.Context) {

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

	//r.Run() // 默认监听并在 0.0.0.0:8080 上启动服务
	r.Run(":8090") // 默认监听并在 0.0.0.0:8080 上启动服务

}
