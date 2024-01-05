package main

import "fmt"

func courseHandler(choice int8) {
	switch choice {

	}
}
func studentHandler(choice int8) {
	fmt.Println(choice)
}

func mainHandler(choice int8) {
	switch choice {
	case 1:
		courseHandler(menu("course"))
		break
	case 2:
		studentHandler(menu("student"))
		break
	}
}
