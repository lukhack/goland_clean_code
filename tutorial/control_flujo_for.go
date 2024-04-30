package tutorial

import "fmt"

func control_flujo_for() {
	for i := 0; i < 10; i++ {

		if i == 2 {
			continue
		} else if i == 5 {
			break
		}
		fmt.Println("i:", i)

	}

	nombre := "lukas"
	for i := 0; i < len(nombre); i++ {
		fmt.Println("i:", string(nombre[i]))
	}

	num := 0

	for num != 100 {
		fmt.Println("i", num)
		num++
	}

}
