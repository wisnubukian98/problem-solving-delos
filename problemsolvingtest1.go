package main

import (
	"fmt"
	"time"
)

func problemsolvingtest1() {
	var y1, m1, d1, y2, m2, d2 int

	fmt.Println("Enter year, month, and date of expected return date:")
	fmt.Scanln(&y1, &m1, &d1)

	fmt.Println("Enter year, month, and date of actual return date:")
	fmt.Scanln(&y2, &m2, &d2)

	if y2 > y1 {
		fine := 12000
		fmt.Println("total fine:", fine)
	}

	t1 := Date(y1, m1, d1)
	t2 := Date(y2, m2, d2)

	yearDiff, monthDiff, dayDiff := diff(t1, t2)
	if yearDiff == 0 {
		if monthDiff > 0 {
			monthFine := 500 * monthDiff
			fmt.Println("total fine:", monthFine)
		}
		if dayDiff > 0 && monthDiff <= 0 {
			dayFine := 15 * dayDiff
			fmt.Println("total fine:", dayFine)
		}
	}
	if yearDiff > 0 {
		fmt.Println("total fine:", 12000)
	}
}

func Date(year, month, day int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}

func diff(a, b time.Time) (year, month, day int) {

	if a.After(b) {
		a, b = b, a
	}
	y1, m1, d1 := a.Date()
	y2, m2, d2 := b.Date()

	year = int(y2 - y1)
	month = int(m2 - m1)
	day = int(d2 - d1)

	if day < 0 {
		t := time.Date(y1, m1, 32, 0, 0, 0, 0, time.UTC)
		day += 32 - t.Day()
		month--
	}
	if month < 0 {
		month += 12
		year--
	}

	return
}
