package main

import (
	"fmt"

	"github.com/ilhamnyto/sprint_case_study/grading"
)

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

	result, err := grading.GradingStudents(studentGrades)

	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}
