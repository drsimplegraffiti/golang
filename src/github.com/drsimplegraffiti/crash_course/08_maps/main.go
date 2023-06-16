package main

import (
    "fmt"
)

func main(){

    email := make(map[string]string)
    email["Bob"] = "abayomi@yopmail.com"
    email["Sharon"] = "sharon@gmail.com"

    fmt.Println(email)
    fmt.Println(len(email))
    fmt.Println(email["Bob"])
    delete(email, "Bob")
    fmt.Println(email)
    
    // declare and initialize a map
    email2 := map[string]string{"James": "j@gmail.com"}
    email2["Bob"] = "bob@gmail.com"
    fmt.Println(email2)
}
