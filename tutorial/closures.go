package tutorial

import "fmt"

func gorras(pedido float32, moneda string) (string, float32, string) {
	precio := func(p float32) float32 {
		return pedido * p
	}

	return "El numero de gorras pedidas es:", precio(2), moneda
}

func patalones(color string, es_corto bool) {
	fmt.Println(color, es_corto)
}

func patalonesv2(caracteristicas ...string) {
	for _, caracteristica := range caracteristicas {
		fmt.Println(caracteristica)
	}

}
