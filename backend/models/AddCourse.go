package models

import (
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func (api *API) AddCourse(ctx *gin.Context) {
	var course Course
	err := ctx.BindJSON(&course)
	if err != nil {
		log.Println(err)
	}
		
	rows, err := api.Db.Query("call add_crs(?, ?)", course.StudentID, course.CourseID)
	if err != nil {
		log.Println(err)
		panic(err)
	}

	for rows.Next() {
		var String string
		_ = rows.Scan(&String)
		if String == "Addition Successful" {

			// ther is a problem with the database
			ctx.JSON(500, gin.H {
				"status": false,
			})
			break
		}

		// no problem
		ctx.JSON(200, gin.H {
			"status": true,
		})

	}
	
}
