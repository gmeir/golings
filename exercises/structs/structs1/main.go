// structs1
// Make me compile!
//
package main

import "fmt"

type Person struct {
  name string
  age int
}

func main() {
  person := Person{"Gil", 48}
	fmt.Printf("Person %s and age %d", person.name, person.age)
}