package models

import (
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func (api *API) GetCourses(ctx *gin.Context) {

	// gives me all the courses that a strudent has registered for 
	var student Student
	err := ctx.BindJSON(&student)
	if err != nil {
		log.Println(err)
	}
	
	rows, err := api.Db.Query("call CrsDetails(?)", student.ID)
	if err != nil {
		log.Println(err)
		return
	}

	courses := []RealCourse{}
	for rows.Next() {

		// find out what are the columns of the table being returned
		var courseName string 
		var _avail int 
		var credits int
		var id int
		var _tot int 

		err := rows.Scan(&courseName, &_avail, &credits, &id, &_tot)
		if err != nil {
			ctx.JSON(500, gin.H {
				"error": "no records found",
			})
			continue
		}

		crs := RealCourse{
			CourseID: id,
			Credits: credits,
			CourseName: courseName,
		}
		courses = append(courses, crs)
	}
	var msg struct {
		Message string `json:"message"`
		Data []RealCourse `json:"data"`
	}
	msg.Message = "success"
	msg.Data = courses
	
	ctx.JSON(200, msg)
	
}