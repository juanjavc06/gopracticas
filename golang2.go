package main

import "fmt"

func main() {
	var horasTrabadas int
	var costoHora float32
	var sueldo float32

	fmt.Print(" Ingrese la cantidad de horas trabajdas")

	fmt.Scan(&horasTrabadas)
	fmt.Print("Ingrese el pago por hora")
	fmt.Scan(&costoHora)

	sueldo = float32(horasTrabadas) * costoHora

	fmt.Print("Sueldo total", sueldo, " Amigos")
}
