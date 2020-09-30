package main

import "fmt"

func main() {
	
	// mapa
	animals := make(map[string]string)
	animals["cat"] = "cat"
	animals["dog"] = "dog"
	fmt.Println(animals)

	// mapa literal
	fruits := map[string]string {
		"apple": "apple",
		"banana": "banana",
	}
	fmt.Println(fruits)

	// borrar un elemento
	delete(fruits, "banana")
	fmt.Println(fruits)

	// obtener un elemento
	fmt.Println(animals["cat"])

	// experimientando con una llave vacía
	// validamos si un valor existe :O
	//_, ok := animals["gorilla"]
	if _, ok := animals["gorrilla"]; !ok {
		animals["gorilla"] = "gorilla"
	}
	fmt.Println(animals)

}
