// Paquete principal
package main

// Modulo para imprimir en consola y mas...
import (
	"fmt"
)

func main() {

	/* Bloque 1: Datos primitivos
	// Declaración de constantes
	const pi_man float64 = 3.14
	const pi2 = 3.1416

	// Método para imprimir en una linea nueva
	fmt.Println("Pi :", pi_man)
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

	// // Operadores Aritmeticos
	x := 10
	y := 50

	// Suma
	result := x + y
	fmt.Println("Suma: ", result)

	// Multiplicacion
	result = x * y
	fmt.Println("Multiplicacion: ", result)

	// Division
	result = y / x
	fmt.Println("Division: ", result)

	// Modulo
	result = y % x
	fmt.Println("Modulo: ", result)

	// Incremental
	x++
	fmt.Println("Incremental: ", x)

	// Decremental
	x--
	fmt.Println("Decremental: ", x)

	// Retos
	// Area de un rectangulo, trapecio y de un circulo
	var alturaCuadrado int = 10
	areaDeRectangulo := baseCuadrado * alturaCuadrado
	fmt.Println("Area de rectangulo: ", areaDeRectangulo)

	// Area de Trapecio
	var baseInferior int = 20
	var baseSuperior = 10
	alturaTrapecio := 5
	areaDeTrapecio := (baseInferior + baseSuperior) * alturaTrapecio / 2
	fmt.Println("Area de trapecio: ", areaDeTrapecio)

	// Area de Circulo
	radio := 8
	var pi float64 = math.Pi
	areaDelCirculo := pi * math.Pow(float64(radio), 2)
	fmt.Println("Area del circulo: ", areaDelCirculo)
	*/

	// Declaración de variables
	helloMessage := "Hello"
	worldMessage := "World"

	// Println
	fmt.Println(helloMessage, worldMessage)

	// Printf
	nombre := "Platzi"
	cursos := 500
	// % s: para string %d: para enteros %f para flotantes
	// %b: para booleanos %c: para caracteres
	fmt.Printf("%s tiene más de %d cursos\n", nombre, cursos)

	// %v: lo usamos para cuando no tenemos certeza del tipo de dato
	// a imprimir
	fmt.Printf("%v tiene más de %v cursos\n", nombre, cursos)

	// Sprintf
	message := fmt.Sprintf("%s tiene más de %d cursos", nombre, cursos)
	fmt.Println(message)

	// Tipo de datos
	fmt.Printf("helloMessage: %T\n", helloMessage)
	fmt.Printf("cursos: %T\n", cursos)
}
