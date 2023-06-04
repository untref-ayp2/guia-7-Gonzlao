package ejercicios

/*func main() {
	arr := []int{7, 4, 6, 8, 5, 9, 3, 1}
	suma := Sumatoria(arr)
	fmt.Println(suma)

}*/

func Sumatoria(arr []int) int {
	if len(arr) == 0 {
		return 0
	} else if len(arr) == 1 {
		return arr[0]
	}

	medio := len(arr) / 2
	suma_izq := Sumatoria(arr[:medio])
	suma_der := Sumatoria(arr[medio:])

	return suma_izq + suma_der
}
