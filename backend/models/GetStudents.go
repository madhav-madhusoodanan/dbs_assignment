package models

import (
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func (api *API) GetStudent(ctx *gin.Context) {

	// gives me all the courses that a strudent has registered for 
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

	courses := []Course{}
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

		crs := Course{
			CourseID: CourseID,
			StudentID: StudentID,
		}
		courses = append(courses, crs)
	}
	var msg struct {
		Message string `json:"message"`
		Data []Course `json:"data"`
	}
	msg.Message = "success"
	msg.Data = courses
	
	ctx.JSON(200, msg)
	
}