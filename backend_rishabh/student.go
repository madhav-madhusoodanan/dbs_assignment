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

type Student struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Password string `json:"password"`
}