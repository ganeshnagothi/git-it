package main

import "fmt"

//area of rectangle
func main() {
	var l int
	var area int
	var w int
	fmt.Print("Enter the value of length :")
	fmt.Scanln(&l)
	fmt.Print("Enter the value of Width :")
	fmt.Scanln(&w)
	area = l * w
	fmt.Printf("Area = %d", area)

}
