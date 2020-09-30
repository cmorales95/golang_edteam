package main

import (
	"fmt"
)

func main() {
	// slice no poseen datos, son apuntadores a un Array
	secondExample()
}

func firstExample() {
	set := [7]string{"lion", "horse", "cow", "butterfly", "bird", "airplane_1", "airplane_2"}
	animals := set[0:5]
	fly := set[3:7]
	fly[0] = "eagle" // Cambiar el valor el array inicial, este modifica el elemento en memoria
	fmt.Println(set)
	fmt.Println(animals)
	fmt.Println(fly)
}

func secondExample() {
	// len() # (longitud)de elementos
	// cap() # (capacidad)numero de elementos del array donde apunta el slice, a partir del indice de donde se creo
	/*food := [5]string{"hotdog", "strowberry", "lemon", "hamburger", "pizza"}
	fruits := food[1:3] // "strowberry, lemon
	fruits = append(fruits, "pinnaple", "cocomelon", "peak", "a", "b", "c", "d", "f") // duplicate

	fmt.Println(food)
	fmt.Println(fruits)
	fmt.Println("fruits size:", len(fruits))
	fmt.Println("fruits capacity: ", cap(fruits))*/

	// Slice literal (o dinámico)
	// fruits := []string{}
	//fruits := []string{"strowberry", "lemon"}
	fruits := make([]string, 3)
	fruits = append(fruits, "pinnapple", "a", "b")
	fmt.Println(fruits)
	fmt.Println("fruits size:", len(fruits))
	fmt.Println("fruits capacity: ", cap(fruits))

}

