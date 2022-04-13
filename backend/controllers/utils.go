package controllers

import (
	"github.com/gin-gonic/gin"
)

func GetCourse(ctx *gin.Context) {
		/* TODO: create the student account, id name and password */
	Api.GetCourse(ctx)
}

func GetStudent(ctx *gin.Context) {
		// give the preference list and check if availabilities are present
		// gives the number of seats that are empty
	Api.GetStudent(ctx)
}

