package main

import (
	"fmt"
	"scholar-manager/pkg/model"
	"scholar-manager/pkg/util"
)

// database
var courses = []model.Course{}
var students = []model.Student{}

// hashmap of menus created to stored all needed menu's on the application
var menus = map[string][]string{
	"main":    {"1. opciones de cursos", "2. opciones de estudiante", "3. opciones de Calificación"},
	"course":  {"1. crear curso", "2. matricular estudiante en curso", "3. ver estudiantes de curso", "4. ver calificaciones", "ver lista de cursos"},
	"student": {"1. crear estudiante", "2. ver cursos inscritos", "3. ver datos de estudiante", "4. ver lista de estudiantes"},
	"grade":   {"1. crear calificación"},
}

// validate if menu input is right
func validarEntrada(menu int, choice int8) bool {
	if choice > 0 && int(choice) <= menu {
		return true
	}

	return false
}

// displays a menu and receives the user's input
func menu(menu string) int8 {
	options := menus[menu]

	util.DisplayList[string](options)
	var choice int8
	fmt.Scanf("%d\n", &choice)

	for validarEntrada(len(options), choice) != true {
		fmt.Println("Invalid selection \n please try again")
		util.DisplayList[string](options)
		fmt.Scanf("%d\n", &choice)
	}

	return choice
}

// main execution method for application running
func main() {

	mainHandler(menu("main"))

}
