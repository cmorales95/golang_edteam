package main

import "fmt"

func main() {
	fruit := "apple"
	// puntero
	var p *string
	p = &fruit // & Operador de dirección
	// modifixar el valor del puntero desde el apuntador
	*p = "pineapple"

	fmt.Printf("Tipo %T, Valor %v, Dirección %v\n", fruit, fruit, &fruit)
	fmt.Printf("Tipo %T, Valor %v, Desreferenciación %s\n", p, &p, *p)

}
