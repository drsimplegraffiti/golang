package main

import (
    "fmt"
)

func main(){
    a := 5
    b := &a // 0xc0000a6058 this is a memory address of a 

    fmt.Println(a,b)
    fmt.Printf("%T\n", a)
    fmt.Printf("%T\n", b) //*int represents a pointer

    // read the value from the address using *
    fmt.Println(*b)
    fmt.Println(*&a)


    // change value with pointer
    *b = 12
    fmt.Println(a)
}
