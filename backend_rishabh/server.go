package main

import (

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	server := gin.Default()

	read := server.Group("/read")
	{
		read.POST("/student", GetStudent)
		read.POST("/course", GetCourse)
	}
	
	write := server.Group("/edit")
	{
		write.POST("/add", Add)
		write.POST("/remove", Substitute)
	}

	server.Run(":8000")
}