package ejercicios

// Escriba un método recursivo que calcule Fibonacci de n.
func Fibonacci(n int) int {
	if n < 3 {
		return 1
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}
