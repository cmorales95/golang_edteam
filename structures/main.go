package main

import "fmt"

func main() {
	// create a basic structure
	type course struct {
		Name      string
		Professor string
		Country   string
	}

	db := course{
		Name:      "Bases de Datos",
		Professor: "Alexys",
		Country:   "Colombia",
	}

	fmt.Printf("%+v\n", db)

	// Instance the structure without call the keys
	// Literal Structure
	git := course{"Git", "Beto", "Bolivia"}
	fmt.Printf("%+v\n",git)

	// Instance only one attribute
	css := course{Professor: "Alvaro"}
	fmt.Printf("%+v\n",css)

	// Print a specific attribute
	fmt.Printf("%+v\n", db.Name)
	fmt.Printf("%+v\n", git.Country)
	fmt.Printf("%+v\n", css.Professor)

	// Pointers
	// Dereference
	p := &db
	p.Professor = "Alvaro" // It's not necessary

	fmt.Printf("%+v\n", db)
	fmt.Printf("%+v\n", p)

}
