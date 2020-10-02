package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	functionsWithParameters("Cristian", "Morales")

	emoji := "dog"
	change(&emoji)
	fmt.Println(emoji)
	fmt.Println(functionWithReturn(2,3)) //nolint:gofmt

	text := "CrISTiaN"
	mins, mays := functionReturnMultipleValues(text)

	fmt.Println(mins, mays, "\n")


	controlErrors()
	functionReceivingFunctions()

	fmt.Println(sum2(1,2,3,4,5))

	x := func(name string) {
		fmt.Println("Hello from func autoinvocate", name)
	}

	x("Good job")
}

func functionsWithParameters(firstName string, lastName string) {
	fmt.Printf("Hello %s %s\n", firstName, lastName)
}

// functions receiving pointers
// param by value
//func change(value *string) {
func change(value *string) {
	*value = "smile" // don't forge the dereference symbol * to change the value from pointer
}

// Golang allow us to indicate the datatype onlye for the end
// of calling num int, num 2 int = bellow
func functionWithReturn(num1, num2 int) int {
	return num1 + num2
}


// WHen we return multiple values we have to specify the
// The N values to specify
func functionReturnMultipleValues(text string) (string, string) {
	may := strings.ToUpper(text)
	min := strings.ToLower(text)
	return min, may
}


func controlErrors() {
	// In general, some standar funcitonalities return the content and error
	// if the error is different of nil, an error was encounter
	content, err := ioutil.ReadFile("./functions/hello.txt")
	if err != nil {
		fmt.Printf("Error %v", err)
		return
		// Is a common practice validate first the erros and then the right value
	}

	// convert the result
	fmt.Println(string(content))
}

// personalize function
func personalFunctionDivision(dividendo, divisor int) (int, error) {
	if divisor == 0{
		return 0, errors.New("You can not divide by zero")
	}
	return dividendo / divisor, nil
}

// personlize funciont v2 ... not recommended
func personalFunctionDivision2(dividendo, divisor int) (result int, err error) {
	if divisor == 0{
		err = errors.New("You can not divide by zero")
		return
	}
	result = dividendo / divisor
	return
}


func functionReceivingFunctions() {
	nums := []int{1, 4, 70, 5, 67, 90, 2}
	result := filter(nums, func(num int) bool {
		return num >= 10
	})

	fmt.Println(result)

	x := hello("Cristian")
	fmt.Println(x("Morales"))
}


func filter(nums []int, callback func(int) bool) []int {
	var result []int
	for _, v := range nums {
		if callback(v) {
			result = append(result, v)
		}
	}
	return result
}

func hello(name string) func(lastname string) string {
	return func(lastname string) string {
		return "Hello " + name + " " + lastname
	}
}

// Variatic Functions
// ... spread operator
func sum2(nums ...int) int {
	total := 0
	for _, v := range nums {
		total += v
	}
	return total
}