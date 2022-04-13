package controllers

import (
	"github.com/gin-gonic/gin"
)

func AddCourse(ctx *gin.Context) {
		/* TODO: create the student account, id name and password */
	Api.AddCourse(ctx)
}

func SubstituteCourse(ctx *gin.Context) {
		// give the preference list and check if availabilities are present
		// gives the number of seats that are empty
	Api.SubstituteCourse(ctx)
}

