package main

import "fmt"

func main() {
	// Slice no poseen datos, son apuntadores a un array
	set := [7]string{"cat", "horse", "cow", "butterfly", "bird", "airplane1", "airplane2"}
	animals := set[0:4]
	fmt.Println(animals)
}
