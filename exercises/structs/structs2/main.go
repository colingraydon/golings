// structs2
// Make me compile!
package main

import "fmt"

type Details struct {
	phone string
}

type Person struct {
	// don't just create the phone field here. embed a new struct
	name string
	age  int
	Details
}

func main() {
	// contactDetails := ContactDetails{}
	person := Person{name: "John", age: 32}
	fmt.Printf("%s is %d years old and his phone is %s\n", person.name, person.age, person.phone)
}
