package main

import "fmt"

func main() {
	i, j := 3, 4

	p := &i
	fmt.Println(*p)
	*p = 21
	fmt.Println(i)

	p = &j
	*p = *p + 30
	fmt.Println(j)
}
