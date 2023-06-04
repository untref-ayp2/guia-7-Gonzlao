package ejercicios

// Escribir un método recursivo que devuelva el
// cociente y el resto de la división entera mediante
// restas sucesivas
func DivisionEntera(dividendo, divisor int) (cociente, resto int) {
	if divisor > dividendo {
		return 0, dividendo
	}

	coc, res := DivisionEntera(dividendo-divisor, divisor)

	return 1 + coc, res
}
