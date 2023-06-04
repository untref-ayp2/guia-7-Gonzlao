package ejercicios

// Escribir un método recursivo que tome un
// array de números enteros y devuelva la suma
// de todos sus elementos
func SumaArray(v []int) int {
	if len(v) == 0 {
		return 0
	}
	if len(v) == 1 {
		return v[0]
	}

	return v[len(v)-1] + SumaArray(v[0:len(v)-1])
}
