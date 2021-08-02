package main

import (
	"docker-test/db"
	"docker-test/file"
	"io/ioutil"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
	db.Setup()
	file.Setup()
}

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	r.GET("/db/add", func(c *gin.Context) {
		data := c.Query("data")

		_, err := db.AddTest(data)
		if err != nil {
			c.JSON(400, gin.H{"message": "error", "error": err.Error()})
			return
		}

		c.JSON(200, gin.H{"message": "success"})
	})

	r.GET("/db/get", func(c *gin.Context) {

		data, err := db.GetTest()
		if err != nil {
			c.JSON(400, gin.H{"message": "error", "error": err.Error()})
			return
		}

		c.JSON(200, gin.H{"message": "success", "data": data})
	})

	r.GET("/file/get", func(c *gin.Context) {
		b, err := ioutil.ReadFile("./storage/test.txt")
		if err != nil {
			c.JSON(400, gin.H{"message": "error", "error": err.Error()})
			return
		}

		c.JSON(200, gin.H{"message": "success", "data": string(b)})
	})

	r.GET("/file/write", func(c *gin.Context) {
		data := c.Query("data")

		err := ioutil.WriteFile("./storage/test.txt", []byte(data), 0644)
		if err != nil {
			c.JSON(400, gin.H{"message": "error", "error": err.Error()})
			return
		}

		c.JSON(200, gin.H{"message": "success"})
	})

	r.Run(":8000")
}
