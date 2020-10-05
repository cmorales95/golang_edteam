package main

import "fmt"

func main() {
	// Instance
	// We could decline the use of the identifier
	Go := Course{
		Name:    "Go desde cero",
		Price:   12.34,
		IsFree:  true,
		UserIds: []uint{12, 56, 89},
		Classes: map[uint]string{
			1: "Introducción",
			2: "Estructura",
			3: "Maps",
		},
	}


	// We could instance specific attributes, 'Zero Value' will assign
	css := Course{Name: "Css desde cero", IsFree: true}

	// Empty instance
	js := Course{}
	js.Name = "Curso JS"
	js.UserIds = []uint{12, 67}

	// Call the struct instance
	fmt.Println(Go.Name)
	fmt.Printf("%+v\n", css)
	fmt.Printf("%+v\n", js)

	// Call a method
	Go.PrintClasses()

	// Go automatically send the pointer
	Go.ChangePrice(67.14)
	fmt.Println(Go.Price)
}
