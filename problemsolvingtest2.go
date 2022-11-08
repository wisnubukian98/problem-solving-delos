package main

import "fmt"

func problemsolvingtest2(s, c, fs int) int {
	fmt.Println("Enter number of students:")
	fmt.Scanln(&s)
	fmt.Println("Enter number of candies:")
	fmt.Scanln(&c)
	fmt.Println("Enter number of candies that the first student will receive:")
	fmt.Scanln(&fs)
	result := candiesdistribution(s, c, fs)
	return result
}

func candiesdistribution(students, candies, firststudent int) int {
	laststudent := students - firststudent + 1
	totalcandy := candies - laststudent
	sourcandystudent := totalcandy % students
	if sourcandystudent != 0 {
		return sourcandystudent
	} else {
		return students
	}
}
