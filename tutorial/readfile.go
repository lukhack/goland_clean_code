package tutorial

import (
	"fmt"
	"os"
)

func Readfile() {
	var file *os.File
	nuevo_texto := "Angular"

	//escribir := os.WriteFile("fichero.txt", nuevo_texto, 0644)

	file, err := os.OpenFile("./tutorial/fichero.txt", os.O_APPEND, 0777)
	showError(err)
	escribir, errs := file.WriteString(nuevo_texto + "\n")
	showError(file.Close())
	showError(errs)

	println("Estado: ", escribir)
	println("Lector")
	fichero, e := os.ReadFile("./tutorial/fichero.txt")
	showError(e)
	fmt.Println(string(fichero))
}

func showError(e error) {
	if e != nil {
		panic("Error leyendo el archivo")
	}
}
