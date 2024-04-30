package tutorial

import "fmt"

func listas() {
	frutas := []string{"Pera", "Manzana", "Mango", "Naranja"}

	fmt.Println(frutas)

	for i := 0; i < len(frutas); i++ {
		fmt.Println(frutas[i])
	}
	fmt.Println("_________________________________________")
	frutas = append(frutas, "Banano")
	for _, fruta := range frutas {
		fmt.Println(fruta)
	}
}
