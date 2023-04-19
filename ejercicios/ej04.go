package ejercicios

// Escriba un método recursivo que tome una cadena y
// devuelva como resultado la cadena invertida.
func Invertir(palabra string) string {
	if len(palabra) <= 1 {
		return palabra
	}

	return Invertir(palabra[1:]) + palabra[:1]
}
