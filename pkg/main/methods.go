package main

import (
	"fmt"
	"scholar-manager/pkg/model"
	"scholar-manager/pkg/util"
)

func createCourse() {
	courseName := util.Scan[string]("Escriba el nombre del curso:")
	courseSummary := util.Scan[string]("Escriba la descripción del curso:")
	newCourse := model.Createcourse(int64(len(courses)), courseName, courseSummary)
	courses = append(courses, newCourse)
	fmt.Println("¡Curso creado exitosamente!")
}

func createStudent() {
	firstName := util.Scan[string]("Escriba el nombre del estudiante")
	lastName := util.Scan[string]("Escriba el apellido del estudiante")
	age := util.Scan[int8]("Escriba la edad del studiante")

	newStudent := model.CreateStudent(int64(len(students)), firstName, lastName, age)
	students = append(students, newStudent)

	fmt.Println("¡Estudiante creado exitosamente!")
}
func addStudentToCourse() {
	courseId := util.Scan[int64]("Escriba el id del curso")

	studentId := util.Scan[int64]("Escriba el id del estudiante")

	courses[courseId].RegisterStudent(students[studentId])
}

func displayCourseStudents() {
	course := courses[util.Scan[int64]("Escriba el id del curso")]
	fmt.Println(course.DisplayCourseStudents())
}
