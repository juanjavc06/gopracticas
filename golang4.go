package main

import "fmt"

func main() {

	var costoHora float32

	horasTrabajadas := 10

	fmt.Println("Indique las horas trabajadas")

	fmt.Scan(&costoHora)

	totalHoras := costoHora * float32(horasTrabajadas)

	fmt.Println("HORAS TRABAJADAS", totalHoras)

}
