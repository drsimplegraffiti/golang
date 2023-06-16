   package main

   import (
       "fmt"
   )

   func main() {
       x := 5
       y := 7

       if x < y {
           fmt.Printf("%d is less than %d\n", x, y)
       } else if x > y {
           fmt.Printf("%d is greater than %d\n", x, y)
       } else {
           fmt.Printf("%d is equal to %d\n", x, y)
       }


       // Switch
       color := "red"

       switch color {
           case "red":
               fmt.Println("color is red")
           case "blue":
            fmt.Println("color is blue")
           default:
                fmt.Println("color is not red or blue")
       }
   }
