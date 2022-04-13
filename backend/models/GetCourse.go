package models

import (
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func (api *API) GetCourse(ctx *gin.Context) {
	// gets me the number of available seats for a course
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

	products := []Course{}
	for rows.Next() {

		// find out what are the columns of the table being returned
		var CourseID string 
		var StudentID string 

		err2 := rows.Scan(&CourseID, &StudentID)
		if err2 != nil {
			ctx.JSON(500, gin.H {
				"error": "no records found",
			})
			continue
		}

		product := Course{
			CourseID: CourseID,
			StudentID: StudentID,
		}
		products = append(products, product)
	}
	var msg struct {
		Message string `json:"message"`
		Data []Course `json:"data"`
	}
	msg.Message = "success"
	msg.Data = products
	
	ctx.JSON(200, msg)
	
}