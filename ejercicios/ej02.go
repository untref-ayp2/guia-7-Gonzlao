package ejercicios

// Escriba un método recursivo que tome un entero n
// devuelva su factorial
func Factorial(n int) int {
	if n == 1 {
		return n
	}

	return n *Factorial(n-1)
}
