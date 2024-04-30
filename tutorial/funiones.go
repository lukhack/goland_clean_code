package tutorial

import "fmt"

func calcular(number1 int, number2 int) int {
	sum := number1 + number2
	return sum
}

func calcular2(number1 int, number2 int) (numeror1 int, numeror2 int, result int) {
	numeror1 = number1
	numeror2 = number2

	result = number1 + number2
	return
}

func Pruebafucion() {

	fmt.Println("numero calculado:", calcular(10, 20))
	sum := calcular(100, 499)
	fmt.Println(sum)

	numeror1, numeror2, result := calcular2(100, 499)

	println(numeror1, numeror2, result)

}
