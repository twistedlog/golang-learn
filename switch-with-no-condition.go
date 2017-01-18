package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Printf("Good Morning")
	case t.Hour() < 17:
		fmt.Printf("Good Evening")
	default:
		fmt.Printf("Good Night")
	}
}
