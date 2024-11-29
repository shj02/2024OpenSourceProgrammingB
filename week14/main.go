package main

import "fmt"

func main() {
	var student1 struct { // 구조체
		id   int
		name string
		gpa  float32
	}
	student1.id = 202444040
	student1.name = "Son Hyeji"
	student1.gpa = 4.5
	fmt.Println(student1.gpa)

	var student2 struct {
		id   int
		name string
		gpa  float32
	}
	student2.id = 202444040
	student2.name = "Son Hyeji"
	student2.gpa = 4.5
	fmt.Println(student2.id)
}
