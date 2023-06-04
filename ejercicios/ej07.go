package ejercicios

// Escriba un método recursivo que tome dos
// números enteros y calcule la multiplicación
// entre ellos, usando sólo sumas
func Multiplicar(a, b int) int {
	if b == 0 || a == 0 {
		return 0
	}
	if b == 1 {
		return a
	}

	return a + Multiplicar(a, b-1)
}
