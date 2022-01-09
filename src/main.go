// Paquete principal
package main

// Modulo para imprimir en consola y mas...
import "fmt"

func main() {
	// Declaración de constantes
	const pi float64 = 3.14
	const pi2 = 3.1416

	// Método para imprimir en una linea nueva
	fmt.Println("Pi :", pi)
	fmt.Println("Pi 2:", pi2)

	// Declaración de variables enteras
	base := 12
	var altura int = 14
	var area int
	var ancho = 25

	fmt.Println(base, altura, area, ancho)

	// Zero values
	var a int
	var b float64
	var c string
	var d bool

	fmt.Println(a, b, c, d)

	// Area cuadrado
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado
	fmt.Println("Area del Cuadrado: ", areaCuadrado)
}
