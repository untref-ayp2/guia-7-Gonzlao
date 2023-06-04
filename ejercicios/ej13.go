package ejercicios

// Escriba un método recursivo para determinar si
// un elemento está en un arreglo utilizando búsqueda binaria. usar tajada
func BusquedaBinaria(arreglo []int, elemento int) bool {
	numero_encontrado := busqueda(arreglo, elemento)
	encontrado := true
	if numero_encontrado == -1 {
		encontrado = false
	}
	return encontrado

}

func busqueda(array []int, x int) int {
	if len(array) <= 0 {
		return -1
	}
	medio := len(array) / 2
	if array[medio] > x {
		return busqueda(array[:medio], x)
	} else if array[medio] < x {
		return busqueda(array[medio+1:], x)
	}
	return medio
}
