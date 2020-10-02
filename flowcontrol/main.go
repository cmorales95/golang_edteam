package main

import (
	"fmt"
	"os"
)

func main() {
	//defer fmt.Println(3)
	//defer fmt.Println(2)
	//defer fmt.Println(1)
	//
	// defer save the state of the arguments
	createFile()

	fmt.Println(division(10, 2))
	fmt.Println(division(40, 2))
	fmt.Println(division(12, 0))
	fmt.Println(division(22, 4))
}


// Example of use: defer
// Closing file with secure
func createFile() {
	file, err := os.Create("hello.txt")
	if err != nil {
		fmt.Println("A error has ocurr %v", err)
		return
	}
	// Close Secure
	defer file.Close()

	_, err = file.Write([]byte("Hello Cristian to this new file"))
	if err != nil {
		fmt.Println("Ocurrió un error al escribir el archivo: %v", err)
		return
	}
}


func division(dividendo, divisor int) int {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovery from the panic", r)
		}
	}()

	panicFunction(dividendo, divisor)
	return dividendo / divisor
}

func panicFunction(dividendo, divisor int) {
	if  divisor == 0 {
		panic("Vales madres")
	}

}