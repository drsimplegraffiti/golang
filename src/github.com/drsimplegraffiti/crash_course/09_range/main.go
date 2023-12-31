package main

import (
    "fmt"
)

func main() {
    ids := []int{34,454,5,65,65,343}

    // Loop through ids
    for i, id := range ids {
        fmt.Printf("%d - ID: %d\n", i, id)
    }

    // Not using index
    for _, id := range ids{
        fmt.Printf("ID: %d\n", id)
    }


    // Add ids together
    sum := 0
    for _, id := range ids {
        sum += id
    }
    fmt.Println("Sum", sum)


   // Range with map
   emails := map[string]string{"Bob": "b@g.com", "Sharon": "s@c.c"}

   for k, v := range emails {
       fmt.Printf("%s: %s\n", k, v)
   }

   for k := range emails {
       fmt.Println("Name: " + k)
   }

}
