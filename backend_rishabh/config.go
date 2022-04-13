package main 

import (
	"github.com/gin-gonic/gin"
)

var Database = New()

func Add(ctx *gin.Context) {
		/* TODO: create the student account, id name and password */
	Database.AddCourse(ctx)
}

func Substitute(ctx *gin.Context) {
		// give the preference list and check if availabilities are present
		// gives the number of seats that are empty
	Database.SubstituteCourse(ctx)
}

func GetCourses(ctx *gin.Context) {
		// get the courses that are available
	Database.GetCourses(ctx)
}

func GetStudentCount(ctx *gin.Context) {
		// get the students that are available
	Database.GetStudentCount(ctx)
}
