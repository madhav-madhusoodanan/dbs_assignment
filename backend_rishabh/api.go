/* Rachit bhai, just Search for "Query"  and whatever is the string in the 1st argument, that goes to the database...take lite about the vaiables involved*/

package main

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
)

type API struct {
	Db *sql.DB
}

func New() API {
	var api API
	db, err := sql.Open("mysql", "user:password@/dbname")
	if err != nil {
		panic(err)
	}

	api.Db = db
	return api
}


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
		// break

	}
	
}

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