package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{3, 4}
	v.X = 7
	fmt.Println(v.X)
}
