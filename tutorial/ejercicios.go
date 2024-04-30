package tutorial

import "fmt"

func ejercicios() {
	fmt.Println("Ejercicios")

	for i := 0; i < 20; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

}
