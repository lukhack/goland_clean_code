package tutorial

import "fmt"

func operatorLogic() {
	suma := 1 + 5
	resta := 10 - 5
	multiplicacion := 10 * 5
	divicion := 10 / 5
	resto := 10 % 2

	fmt.Println(suma, resta, multiplicacion, divicion, resto)

	//operadores incremento y decremento
	suma++
	resta--
	fmt.Println(suma, resta)

	suma += 10
	resta -= 2
	fmt.Println(suma, resta)

	fmt.Println(10 > 5)
}
