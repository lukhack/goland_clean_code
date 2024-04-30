package tutorial

import (
	"fmt"
	"math/rand"
)

type User struct {
	id   int
	Name string
}

type Employee struct {
	User
	Active bool
}

type Cashier interface {
	CalcTotal(item ...float32) float32
	deactivate()
}

type Admin interface {
	DeactiveEmployee(c Cashier)
}

func Newemployee(name string) *Employee {
	return &Employee{
		User: User{
			id:   rand.Intn(1000),
			Name: name,
		},
		Active: true,
	}
}

func (e *Employee) CalcTotal(item ...float32) float32 {

	if !e.Active {
		fmt.Println("Cahier desactivated")
	}

	var sum float32
	for _, v := range item {
		sum += v
	}
	return sum * 1.15
}

func (e *Employee) deactivate() {
	e.Active = false
}

func (e *Employee) DeactiveEmployee(c Cashier) {
	c.deactivate()
}

func interfaces() {
	println("hola")
	var frank Cashier = Newemployee("lukas")
	var robert Admin = Newemployee("Robert")

	var total = frank.CalcTotal(80, 42, 44, 66)
	println(total)

	robert.DeactiveEmployee(frank)

	total = frank.CalcTotal(80, 42, 44, 66)
	println(total)

}
