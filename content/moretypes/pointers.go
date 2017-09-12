// +build OMIT

package main

import "fmt"

func main() {
	i, j := 42, 2701

	p := &i         // p сочи към i
	fmt.Println(*p) // прочита се i през указателя
	*p = 21         // задава се i през указателя
	fmt.Println(i)  // показва новата стойност на i

	p = &j         // p сочи към j
	*p = *p / 37   // деление на j чрез указателя
	fmt.Println(j) // показва новата стойност на j
}
