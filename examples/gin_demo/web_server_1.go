package gin_demo

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func PingPong() {
	// 创建服务
	r := gin.Default()

	// 服务地址（路径）， 处理请求函数（匿名函数）
	r.GET("/ping", func(c *gin.Context) {

		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	//r.Run() // 默认监听并在 0.0.0.0:8080 上启动服务
	r.Run(":8090") // 默认监听并在 0.0.0.0:8080 上启动服务

	// http://localhost:8090/ping
	// {"message":"pong"}

}

func RESTfulDemo() {
	// 创建服务
	r := gin.Default()

	// 查询用户
	r.POST("/user", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"action": "新建",
			"name":   "user1",
			"age":    10,
		})
	})

	// 查询用户
	r.GET("/user", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"action": "查询",
			"name":   "user1",
			"age":    10,
		})
	})

	// 修改用户
	r.PUT("/user", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"action": "修改",
			"name":   "user2",
			"age":    20,
		})
	})

	// 删除用户
	r.DELETE("/user", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"action": "删除",
			"name":   "user1",
			"age":    10,
		})
	})
	//r.Run() // 默认监听并在 0.0.0.0:8080 上启动服务
	r.Run(":8090") // 默认监听并在 0.0.0.0:8080 上启动服务

}

func LoadStaticPage() {
	// 创建服务
	r := gin.Default()

	// 加载静态页面
	r.LoadHTMLGlob("examples/gin_demo/templates/*")
	//r.LoadHTMLFiles("examples/gin_demo/templates/index.html")

	r.Static("/w3cs", "examples/statics")

	// 返回静态页面
	r.GET("/index", func(c *gin.Context) {
		//c.JSON(http.StatusOK, gin.H{
		//	"message": "pong",
		//})
		c.HTML(http.StatusOK, "index.html", gin.H{
			"msg": "gin框架后端 LoadStaticPage 传过来的参数",
		})

	})

	r.GET("/rain", func(c *gin.Context) {
		//c.JSON(http.StatusOK, gin.H{
		//	"message": "pong",
		//})
		c.HTML(http.StatusOK, "code_rain.html", gin.H{
			"msg": "gin框架后端 LoadStaticPage 传过来的参数",
		})

	})

	r.GET("/echarts0", func(c *gin.Context) {
		//c.JSON(http.StatusOK, gin.H{
		//	"message": "pong",
		//})
		c.HTML(http.StatusOK, "echarts0.html", gin.H{
			"msg": "gin框架后端 LoadStaticPage 传过来的参数",
		})

	})

	r.GET("/echarts1", func(c *gin.Context) {
		//c.JSON(http.StatusOK, gin.H{
		//	"message": "pong",
		//})
		c.HTML(http.StatusOK, "echarts1.html", gin.H{
			"msg": "gin框架后端 LoadStaticPage 传过来的参数",
		})

	})

	r.GET("/echarts1-1", func(c *gin.Context) {
		//c.JSON(http.StatusOK, gin.H{
		//	"message": "pong",
		//})
		c.HTML(http.StatusOK, "echarts1-1.html", gin.H{
			"msg": "gin框架后端 LoadStaticPage 传过来的参数",
		})

	})

	r.GET("/echarts2", func(c *gin.Context) {
		//c.JSON(http.StatusOK, gin.H{
		//	"message": "pong",
		//})
		c.HTML(http.StatusOK, "echarts2.html", gin.H{
			"msg": "gin框架后端 LoadStaticPage 传过来的参数",
		})

	})

	r.GET("/echarts_profile", func(c *gin.Context) {
		//c.JSON(http.StatusOK, gin.H{
		//	"message": "pong",
		//})
		c.HTML(http.StatusOK, "echarts_profile.html", gin.H{
			"msg": "gin框架后端 LoadStaticPage 传过来的参数",
		})

	})

	//r.Run() // 默认监听并在 0.0.0.0:8080 上启动服务
	r.Run(":8090") // 默认监听并在 0.0.0.0:8080 上启动服务

}

func GetParametersInRequests1() {
	// 创建服务
	r := gin.Default()

	// 获取client参数并传递给前端
	r.GET("/user/info", func(c *gin.Context) {
		userID := c.Query("userid")
		userName := c.Query("username")
		c.JSON(http.StatusOK, gin.H{
			"ID":   userID,
			"Name": userName,
		})
	})

	//r.Run() // 默认监听并在 0.0.0.0:8080 上启动服务
	r.Run(":8090") // 默认监听并在 0.0.0.0:8080 上启动服务

}

func GetParametersInRequests2() {
	// 创建服务
	r := gin.Default()

	// 获取client参数并传递给前端
	r.GET("/user/info/:userid/:username", func(c *gin.Context) {
		userID := c.Param("userid")
		userName := c.Param("username")
		c.JSON(http.StatusOK, gin.H{
			"ID":   userID,
			"Name": userName,
		})
	})

	//r.Run() // 默认监听并在 0.0.0.0:8080 上启动服务
	r.Run(":8090") // 默认监听并在 0.0.0.0:8080 上启动服务

}
