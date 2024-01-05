package model

type Grade struct {
	gradeId int64
	name    string
	grade   float32
	summary string
}

type Course struct {
	courseId int64
	name     string
	summary  string
	students []Student
	grades   map[int64][]Grade
}

type Student struct {
	studentId int64
	firstName string
	lastName  string
	age       int8
	courses   []int64
}
