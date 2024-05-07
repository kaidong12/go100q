package gin_demo

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Redirect_to_another_page() {
	// 创建服务
	r := gin.Default()

	// load page
	r.LoadHTMLGlob("examples/gin_demo/templates/*")

	r.GET("/goto1", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "http://www.baidu.com")
	})

	r.GET("/goto2", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/index2")
	})

	// 返回静态页面
	r.GET("/index2", func(c *gin.Context) {
		//c.JSON(http.StatusOK, gin.H{
		//	"message": "pong",
		//})
		c.HTML(http.StatusOK, "index.html", gin.H{
			"msg": "gin框架后端 Redirect_to_another_page 传过来的参数",
		})

	})

	//r.Run() // 默认监听并在 0.0.0.0:8080 上启动服务
	r.Run(":8090") // 默认监听并在 0.0.0.0:8080 上启动服务

}

func Route_group() {
	// 创建服务
	r := gin.Default()

	userGroup := r.Group("/user")
	{
		userGroup.GET("/add")
		userGroup.POST("/login")
		userGroup.POST("/logout")

	}

	orderGroup := r.Group("/order")
	{
		orderGroup.GET("/add")
		orderGroup.POST("/list")
		orderGroup.POST("/edit")

	}

	//r.Run() // 默认监听并在 0.0.0.0:8080 上启动服务
	r.Run(":8090") // 默认监听并在 0.0.0.0:8080 上启动服务

}

func Redirect_to_404_page() {
	// 创建服务
	r := gin.Default()

	// load page
	r.LoadHTMLGlob("examples/gin_demo/templates/*")

	r.GET("/goto3", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "http://www.bbbbbb.com")
	})

	r.GET("/goto4", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/index4")
	})

	r.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusNotFound, "404.html", nil)
	})

	//r.Run() // 默认监听并在 0.0.0.0:8080 上启动服务
	r.Run(":8090") // 默认监听并在 0.0.0.0:8080 上启动服务

}
