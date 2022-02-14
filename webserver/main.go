package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Welcome to the world of gin !")
	})
	r.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s !", name)
	})
	r.GET("/weather/:city", func(c *gin.Context) {
		city := c.Param("city")
		c.JSON(http.StatusOK, gin.H{
			"city_name": city,
			"weather":   "sunny",
		})
	})
	/**
	http://127.0.0.1:8080/weather/wuhan

	{
	    "city_name": "wuhan",
	    "weather": "sunny"
	}
	*/
	//http://127.0.0.1:8080/users?name=Tom&role=students
	r.GET("/users", func(c *gin.Context) {
		name := c.Query("name")
		role := c.DefaultQuery("role", "teacher")
		c.String(http.StatusOK, "%s is a %s", name, role)
	})

	r.POST("/form", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.DefaultPostForm("password", "000000")
		c.JSON(http.StatusOK, gin.H{
			"username": username,
			"password": password,
		})
	})
	// POST接口定义
	r.POST("/", func(c *gin.Context) {
		json := make(map[string]interface{})
		c.BindJSON(&json)
		c.JSON(http.StatusOK, gin.H{
			"server_info": json["server"],
			"ip_info":     json["ip"],
		})

	})

	r.Run(":9933") // listen and serve on 0.0.0.0:9933
}
