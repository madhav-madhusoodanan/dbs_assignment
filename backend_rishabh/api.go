package main 

import (
	"log"
	"database/sql"
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
	var crs Course
	err := ctx.BindJSON(&crs)
	if err != nil {
		log.Println(err)
	}
		
	_, err = api.Db.Query("call add_crs(?, ?)", crs.StudentID, crs.CourseID)
	if err != nil {
		log.Println(err)
	}

	// what is the output?
	
	ctx.JSON(200, gin.H {
		"message": "course addition successful",
	})
}

func (api *API) SubstituteCourse(ctx *gin.Context) {
	var std StudentPreference
	err := ctx.BindJSON(&std)
	if err != nil {
		log.Println(err)
	}
	
	rows, err := api.Db.Query("call substi_crs(?, ?, ?)", std.ID, std.OldCourse, std.NewCourse)
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

func (api *API) GetCourse(ctx *gin.Context) {
	var std StudentPreference
	err := ctx.BindJSON(&std)
	if err != nil {
		log.Println(err)
	}
	
	rows, err := api.Db.Query("call substi_crs(?, ?, ?)", std.ID, std.OldCourse, std.NewCourse)
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

func (api *API) GetStudent(ctx *gin.Context) {
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