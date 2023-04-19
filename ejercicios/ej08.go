package ejercicios

// Escribir un método recursivo que devuelva el
// cociente y el resto de la división entera mediante
// restas sucesivas
func DivisionEntera(dividendo, divisor int) (cociente, resto int) {
	panic("Not implemented")
}

func dividir(a, b int) int {

	if a < b || b == 0 {
		return 0
	}

	return 1 + dividir(a-b, b)
}
