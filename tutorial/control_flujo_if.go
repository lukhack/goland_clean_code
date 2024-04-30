package tutorial

import "fmt"

func controlFlujo_if() {
	numero := 45

	if numero == 45 {
		fmt.Println("valor es igual")
		fmt.Println(numero)
	} else if numero == 40 {
		fmt.Println("valor no es igual a 40")

	} else {
		fmt.Println("valor no es igual")

	}

	edad := 30

	if edad >= 18 {
		fmt.Println("Eres mayor de edad")
	} else if edad <= 18 && edad >= 0 {
		fmt.Println("Eres menor de edad")
	} else {
		fmt.Println("Edad no es valida")
	}

}
