package main

import (
	"errors"
)

func GradingStudents(grades []int) (res []int, err error) {
	if len(grades) < 1 || len(grades) > 60 {
		return nil, errors.New("Number of students is incorrect.")
	}
	var result []int

	for _, grade := range grades {

		if grade < 0 || grade > 100 {
			return nil, errors.New("Student's grade is incorrect.")
		}

		if grade < 38 {
			result = append(result, grade)
			continue
		}

		remainder := grade % 5

		if remainder == 0 {
			result = append(result, grade)
			continue
		}

		multipled_number := grade + (5 - remainder)

		if (multipled_number - grade) < 3 {
			result = append(result, multipled_number)
			continue
		}

		result = append(result, grade)
	}

	return result, nil
}