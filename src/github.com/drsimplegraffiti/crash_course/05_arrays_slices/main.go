package main
import (
	"fmt"
)




func main(){
	// Arrays
	var myArray [2]string

	// Assigning values
	myArray[0] = "Hello"
	myArray[1] = "World"

	fmt.Println(myArray[0], myArray[1])
	fmt.Println(myArray)

	// Declare and assign
	myArray2 := [2]string{"peanut", "butter"}
	fmt.Println(myArray2)

	// Slices
	mySlice := []string{"peanut", "butter", "jelly"}
	fmt.Println(mySlice)

	// length
	fmt.Println(len(mySlice))

	// range
	fmt.Println(mySlice[1:3])
}
