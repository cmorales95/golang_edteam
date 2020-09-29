package main

import "fmt"
import "greeting"

func main() {
	//var dog string = "dog"
	// bool, string, numeric
	
	var a bool = true
	var b string = "EDteam"
	var c uint8 = 100
	var d uint16 = 23000
	e := uint16(c) + d
	_ = 234

	fmt.Printf("Tipo: %T, Valor: %v\n", a, a)
	fmt.Printf("Tipo: %T, Valor: %v\n", b, b)
	fmt.Printf("Tipo: %T, Valor: %v\n", e, e)

	// declaración post-incremento y post-decremento: ++, --
	var f = 3
	//f = f++ + 2 incorrecto
	f++
	println(f)

	fmt.Println(4 == 4) //

	// operadores and(&&) or(||)
	// operador unario !

}


b58f770b-6350-413f-9c38-a77b0a0a8941