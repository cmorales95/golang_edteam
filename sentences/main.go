package main

import "fmt"

func main() {
	sentencesSwitch()
	sentenceFor()
}

func sentencesSwitch() {
	// We can work withou expresion, we can use the expresion into the each case (multiple sentences
	// Case emoji = "cat" || emoji = "dog"
	emoji := "cat"
	switch emoji {
	case "cat", "dog":
		fmt.Println("It is animal")
	case "banana", "apple":
		fmt.Println("it is a fruit")
	default:
		fmt.Println("no animal or fruit")
	}
}

func sentenceFor() {
	// classic for
	for i := 1; i <= 10;i++ {
		fmt.Println(i)
	}

	// continues for (while)
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	// for forever
	/*for {
		i++
	}*/

	// for range slice
	nums := []uint8{2, 4, 6, 8}
	/*for i, v := range nums {
		fmt.Println("indice:", i, "valor", v)
	}*/

	for i := range nums {
		nums[i] *= 2
	}

	sport := map[string]string{"basketball": "bask"}

	for key, val := range sport {
		fmt.Println(key, val)
	}

	// range string
	hello := "hello"
	for _, v := range hello {
		fmt.Println(string(v))
	}
}