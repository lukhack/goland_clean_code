package tutorial

import "fmt"

func Slice() {
	peliculas := []string{
		"La ivda es bella"}

	peliculas = append(peliculas, "Avenger")

	fmt.Println(peliculas)

	fmt.Println(peliculas[0:2])

}
