package main

import (
	"fmt"
	"os"
)



func main() {
	fmt.Println("Hello World")
	f, err := os.Open("test.txt")
	if err != nil {
		fmt.Println("Error" , err)
	}
	defer f.Close() // defer will execute at the end of the function

	// Read the file
	rf := make([]byte, 128)
	f.Read(rf)
	fmt.Println(string(rf))
}
