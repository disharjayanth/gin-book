package main

import "github.com/gin-gonic/gin"

func IndexHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello world",
	})
}

func GetNameHandler(c *gin.Context) {
	name := c.Params.ByName("name")
	c.JSON(200, gin.H{
		"message": "Hello " + name,
	})
}

func main() {
	router := gin.Default()

	router.GET("/", IndexHandler)
	router.GET("/:name", GetNameHandler)

	router.Run()
}
