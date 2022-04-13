package main

type StudentPreference struct {
	ID string `json:"id"`
	OldCourse string `json:"old_course"`
	NewCourse string `json:"new_course"`
}

type Course struct {
	StudentID string `json:"student_id"`
	CourseID string `json:"course_id"`
}

type RealCourse struct {
	CourseID int `json:"course_id"`
	CourseName string `json:"course_name"`
	Credits int `json:"credits"`
}

type Student struct {
	ID string `json:"id"`
}

type CourseID struct {
	CourseID string `json:"course_id"`
}