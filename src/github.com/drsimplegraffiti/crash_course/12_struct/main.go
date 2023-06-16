package main


import (
    "fmt"
    "strconv"
)

type Person struct {
    firstName string
    lastName string
    age int
    city string
    gender string

}

// Greeting method (value receiver)
func (p Person) greet() string{
    return "Hello, my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age) + " years old."
}


// hasBirthday method (pointer receiver
func (p *Person) hasBirthday(){
    p.age++
}

// getMarried (pointer receiver)
func (p *Person) getMarried(spouseLastName string){
    if p.gender == "M"{
        return
    } else {
        p.lastName = spouseLastName
    }
}


func main(){
    person1 := Person{
        firstName: "John",
        lastName: "Doe",
        age: 25,
        city: "New York",
    }
    
    fmt.Println(person1)

    firstName := "John"
    lastName := "Doe"
    age := 25
    city := "New York"
    gender := "F"

    person2 := Person{
        firstName,
        lastName,
        age,
        city,
        gender,
        
}

    fmt.Println(person2)
    fmt.Println(person2.firstName)
    person2.age++
    fmt.Println(person2.age)
    fmt.Println(person2.greet())
    person2.hasBirthday()
    person2.getMarried("Williams")
    fmt.Println(person2.greet())
}

