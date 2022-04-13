package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	server := gin.Default()

	read := server.Group("/read")
	{
		read.POST("/student", GetStudentCount)
		read.POST("/course", GetCourses)
	}
	
	write := server.Group("/edit")
	{
		write.POST("/add", Add)
		write.POST("/substitute", Substitute)
	}

	server.Run(":8000")
}