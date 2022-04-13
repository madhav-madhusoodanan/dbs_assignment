package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	server := gin.Default()
	Interface := New()

	read := server.Group("/read")
	{
		read.POST("/student", Interface.GetStudentCount)
		read.POST("/course", Interface.GetCourses)
	}
	
	write := server.Group("/edit")
	{
		write.POST("/add", Interface.AddCourse)
		write.POST("/substitute", Interface.SubstituteCourse)
	}

	server.Run(":8000")
}