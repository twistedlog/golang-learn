package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func main(){
	a, b := swap("2", "3")
	fmt.Println(a, b)
}
