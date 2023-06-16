package main

import (
    "fmt"
)

func main(){
    var name string = "Joe"
    var age int = 45
    var isCool = true
    fmt.Println(name, age)
    fmt.Printf("%T\n", name)
    fmt.Printf("%t\n", isCool)

    const height = 5.9
    fmt.Printf("%T\n", height) // %T prints the type of the variable

    // shorthand
    size := 1.3
    fmt.Printf("%T\n", size)

    // multiple variables
    name1, email := "Joe", "jo@gmail.com"
    fmt.Println(name1, email)

//Note: You can only use shorthand inside a function body

    var cgpa float32 = 9.5
    fmt.Printf("%T\n", cgpa)
}
