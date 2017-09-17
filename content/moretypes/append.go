// +build OMIT

package main

import "fmt"

func main() {
	var s []int
	printSlice(s)

	// append работи и с нулирани отрязъци.
	s = append(s, 0)
	printSlice(s)

	// Отрязъкът расте колкото е необходимо.
	s = append(s, 1)
	printSlice(s)

	// Можем да добавим няколко елемента наведнъж.
	s = append(s, 2, 3, 4)
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
