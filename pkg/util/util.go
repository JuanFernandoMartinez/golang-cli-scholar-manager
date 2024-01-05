package util

import "fmt"

//generic function to display comparable type lists
func DisplayList[K comparable](items []K) {
	for i := 0; i < len(items); i++ {
		fmt.Printf("%v\n", items[i])
	}
}

func Scan[T comparable](message string) T {

	fmt.Println(message)
	var value T
	fmt.Scanf("%v\n", &value)

	return value
}
