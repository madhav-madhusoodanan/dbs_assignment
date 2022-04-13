package controllers

import (
	"github.com/gin-gonic/gin"
)

func GetStudentCount(ctx *gin.Context) {
		/* TODO: create the student account, id name and password */
	Api.GetStudentCount(ctx)
}

func GetCourses(ctx *gin.Context) {
		// give the preference list and check if availabilities are present
		// gives the number of seats that are empty
	Api.GetCourses(ctx)
}

