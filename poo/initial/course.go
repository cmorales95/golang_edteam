package main

import "fmt"

// With Upper to expose
// we could add slice, arrays, functions, etc
type Course struct {
	Name    string
	Price   float64
	IsFree  bool
	UserIds []uint
	Classes map[uint]string
}

// To define this function as method change the order of parenthesis
// Please, don't use an argument called 'this' or 'self' word use the first letter of the struct
// Normal function: func PrintClasses(c Course) {
// Add method: func (c Course) PrintClasses() {
func (c *Course) PrintClasses() {
	text := "Las clases son:"
	for _, class := range c.Classes {
		text += class + ", "
	}
	fmt.Println(text[:len(text)-2])
}

// We need to use pointers if we want to update attributes through methods
func (c *Course) ChangePrice(price float64) {
	c.Price = price
}

// Tip: When a method use a pointer, please change all of your methods to use pointer (Interface class more information)