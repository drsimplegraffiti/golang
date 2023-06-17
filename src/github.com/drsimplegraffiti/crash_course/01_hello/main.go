package main

import (
    "fmt"
)

func main(){
 fmt.Println("Hello world😀😀")
 a := 50
 b := &a // the & operator gives the address of the variable

 c := a + *b // the * operator gives the value stored at an address when you have the address
 fmt.Println(c)
}
