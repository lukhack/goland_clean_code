package tutorial

import "fmt"

type Cliente struct {
	nombre   string
	telefono string
	next     *Cliente
}

var primerCliente *Cliente

func Load() {
	
}

func Load2() {
	var v int = 4
	var p *int

	p = &v
	*p = 8

	incValor(v)
	incReferencia(&v)
	incReferencia(&v)

	println("*p: ", *p)
}

func incValor(v int) {
	v++
	fmt.Println("v1: ", v)
}

func incReferencia(v *int) {
	*v++
	fmt.Println("v2: ", *v)
}
