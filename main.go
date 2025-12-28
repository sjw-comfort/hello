package main

import (
	"fmt"
)

func main() {
	
	a := 100
	b := &a
	fmt.Println( a, &a) 
	fmt.Println( b, &b) 
	
}
