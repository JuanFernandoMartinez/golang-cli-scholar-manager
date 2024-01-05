package model

import (
	"strconv"
)

func Createcourse(courseId int64, name string, summary string) Course {
	newCourse := Course{
		courseId: courseId,
		name:     name,
		summary:  summary,
		students: []Student{},
		grades:   map[int64][]Grade{},
	}

	return newCourse
}

func CreateStudent(studentId int64, fn string, ln string, age int8) Student {
	return Student{
		studentId: studentId,
		firstName: fn,
		lastName:  ln,
		age:       age,
		courses:   []int64{},
	}
}

func (student Student) PutIntoCourse(courseId int64) {
	student.courses = append(student.courses, courseId)
}

func (course Course) RegisterStudent(student Student) {
	student.courses = append(student.courses, course.courseId)
	course.students = append(course.students, student)
	course.grades[student.studentId] = []Grade{}
}

func (course Course) AddGrade(grade Grade, studentId int64) {
	course.grades[studentId] = append(course.grades[studentId], grade)
}

func (student Student) toString() string {
	stringValue := "Student ID: " + strconv.Itoa(int(student.studentId)) + "First Name: " + student.firstName + "\nLast Name: " + student.lastName + "\nAge: " + strconv.Itoa(int(student.age))

	return stringValue
}

func (course Course) DisplayCourseStudents() string {
	output := ""
	for i := 0; i < len(course.students); i++ {
		output += course.students[i].toString()
	}

	output += "\n"

	return output
}
