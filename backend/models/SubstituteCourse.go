package models

import (
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func (api *API) SubstituteCourse(ctx *gin.Context) {
	var student StudentPreference
	err := ctx.BindJSON(&student)
	if err != nil {
		log.Println(err)
	}
	
	rows, err := api.Db.Query("call substi_crs(?, ?, ?)", student.ID, student.OldCourse, student.NewCourse)
	if err != nil {
		log.Println(err)
		return
	}

	for rows.Next() {

		// find out what are the columns of the table being returned
		var String string 

		_ = rows.Scan(&String)
		if String == "Substitution Successful" {
			ctx.JSON(500, gin.H {
				"status": false,
			})
			break
		}
		
		ctx.JSON(200, gin.H {
			"status": true,
		})

	}
	
}