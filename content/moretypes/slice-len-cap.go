// +build OMIT

package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// Отрежете отрязъка, за да му зададете нулева дължина.
	s = s[:0]
	printSlice(s)

	// Удължете го.
	s = s[:4]
	printSlice(s)

	// Изхвърлете първите му две стойности.
	s = s[2:]
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
