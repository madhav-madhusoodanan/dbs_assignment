package models

import (
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func (api *API) GetStudentCount(ctx *gin.Context) {
	// gets me the number of available seats for a course
	var course CourseID
	err := ctx.BindJSON(&course)
	if err != nil {
		log.Println(err)
	}
	
	rows, err := api.Db.Query("select avail_seats from course where course.id = ?", course.CourseID)
	if err != nil {
		log.Println(err)
		return
	}

	for rows.Next() {

		// find out what are the columns of the table being returned
		var count int 

		err = rows.Scan(&count)
		if err != nil {
			ctx.JSON(500, gin.H {
				"error": "no records found",
			})
			continue
		}
		var msg struct {
			Message string `json:"message"`
			Count int `json:"count"`
		}
		msg.Message = "success"
		msg.Count = count
		
		ctx.JSON(200, msg)
	}
	
}