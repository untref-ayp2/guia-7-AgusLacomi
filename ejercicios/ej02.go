package ejercicios

// Escriba un método recursivo que tome un entero n
// devuelva su factorial
func Factorial(a int) int {
	if a == 1 {
		return 1
	}
	if a == 0 {
		return 0
	}

	return a * Factorial(a-1)
}
