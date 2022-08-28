package main

import "fmt"

func main() {
	var horasTrabadas int
	var costoHora float32
	var sueldo float32

	var tipo int

	fmt.Print(" Ingrese la cantidad de horas trabajdas")

	fmt.Scan(&horasTrabadas)
	fmt.Print("Ingrese el pago por hora")
	fmt.Scan(&costoHora)

	sueldo = float32(horasTrabadas) * costoHora

	if sueldo > 80 {
		fmt.Print("UWUW")
	} else {
		fmt.Print("DOM DOM")
	}

	fmt.Print("Sueldo total", tipo)
}
