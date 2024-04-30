package tutorial

import "fmt"

type Person struct {
	name string
	age  int
}

func (p Person) saludar(saludo string) {
	println(saludo + ", " + p.name)
}

func (p Person) cumple() int {
	return p.age + 1
}

func structs() {
	persona1 := Person{name: "Alice", age: 30}
	fmt.Println(persona1)

	persona1.age = 40
	fmt.Println(persona1)

	persona1.saludar("hola como estas")
	fmt.Println("edad: ", persona1.cumple())
}
