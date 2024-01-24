package main

import "fmt"

func main() {
	var numStudent int
	var studentGrades []int

	fmt.Scanln(&numStudent)

	i := 1
	for i <= numStudent {
		var grade int

		fmt.Scanln(&grade)
		studentGrades = append(studentGrades, grade)

		i += 1
	}

	fmt.Println(studentGrades)

	result, err := GradingStudents(studentGrades)

	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}
