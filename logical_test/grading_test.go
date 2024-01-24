package main

import (
	"reflect"
	"testing"
)

func TestGradingStudents(t *testing.T) {
	testCases := []struct {
		name	 string
		input    []int
		expected []int
		errExpected bool
	}{
		{"Test 1", []int{84, 29, 57}, []int{85, 29, 57}, false},
		{"Test 2", []int{73, 67, 38, 33}, []int{75, 67, 40, 33}, false},
		{"Test 3", []int{37, 38}, []int{37, 40}, false},
		{"Test 4", []int{35, 40, 45}, []int{35, 40, 45}, false},
		{"Test student's grades less than 0", []int{73, 67, -1, 33}, nil, true},
		{"Test student's grades greater than 100", []int{73, 67, 38, 105}, nil, true},
		{"Test number of students greater less than 1", []int{}, nil, true},
		{"Test number of students greater than 60", make([]int, 62), nil, true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := GradingStudents(tc.input)

			if tc.errExpected {
				if err == nil {
					t.Errorf("Expected no error, but got: %v", err)
				}
			} else {
				if err != nil {
					t.Errorf("Expected no error, but got: %v", err)
				}

				if !reflect.DeepEqual(result, tc.expected) {
					t.Errorf("Expected %v, but got %v", tc.expected, result)
				}
			}
		})
	}
}