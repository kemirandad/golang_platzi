// Paquete principal
package main

// Modulo para imprimir en consola y mas...
import (
	"fmt"
	"strings"
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

	/* Bloque 2: fmt
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

	// Scanf
	var name string
	var age int
	fmt.Sscanf("Kim is 22 years old", "%s is %d years old", &name, &age)
	fmt.Printf("%s, %d\n", name, age)
	*/
	/* Bloque 3: funciones
	normalFunction("Hola Mundo")
	tripeArgument(1, 2, "Kenny")

	value := returnValue(2)
	fmt.Println("Value: ", value)

	value1, value2 := doubleReturn(2)
	fmt.Println("Value 1: ", value1, "***  Value 2: ", value2)

	_, value3 := doubleReturn(5)
	fmt.Println("Value 1: ", nil, "***  Value 2: ", value3)
	*/

	/* Bloque 4: Bucle For
	// For condicional
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}

	fmt.Println()

	// For while
	counter := 0
	for counter < 10 {
		fmt.Println(counter)
		counter++
	}

	fmt.Println()

	// For forever
	counterForever := 0
	for {
		fmt.Println(counterForever)
		counterForever++
		if counterForever == 20 {
			break
		}
	}

	fmt.Println()

	// Reto
	anotherCounter := 10
	for i := anotherCounter; i >= 0; i-- {
		fmt.Println(i)
	}

	listaNumerosPares := []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}
	for i, par := range listaNumerosPares {
		fmt.Printf("posicion %d número par: %d \n", i, par)
	}
	*/

	/* Bloque 5: operadores lógicos
	valor1 := 1
	valor2 := 2

	if valor1 == 1 {
		fmt.Println("Es 1")
	} else {
		fmt.Println("No es 1")
	}

	// With and
	if valor1 == 1 && valor2 == 3 {
		fmt.Println("Es verdad")
	}

	// With or
	if valor1 == 0 || valor2 == 2 {
		fmt.Println("Esto tambien es verdad")
	}

	// Reto
	for i := 0; i <= 4; i++ {
		isPar(i)
	}

	fmt.Println("No autorizado: ", auth("kemirandad", "123"))
	fmt.Println("Autorizado: ", auth("kemirandad", "1234"))
	*/

	/* Bloque 6: Uso de switch
	modulo := 4 % 2
	switch modulo {
	case 0:
		fmt.Println("Es par")
	default:
		fmt.Println("Es impar")
	}

	switch modulo := 4 % 2; modulo {
	case 0:
		fmt.Println("Es par")
	default:
		fmt.Println("Es impar")
	}

	// sin condicion
	value := 50
	switch {
	case value > 100:
		fmt.Println("Es mayor a 100")
	case value < 0:
		fmt.Println("Es menor a 0")
	default:
		fmt.Println("No condicion")
	}
	*/
	/* Bloque 7: Keywords
	// Defer
	defer fmt.Println("Hola")
	fmt.Println("Mundo")

	// Continue y break
	for i := 0; i < 10; i++ {
		fmt.Println(i)

		// Continue
		if i == 2 {
			fmt.Println("Es dos")
			continue
		}

		// Break
		if i == 8 {
			fmt.Println("Break")
			break
		}
	}
	*/
	/* Bloque 8: Array y Slices
	// Array
	var array [4]int
	array[0] = 1
	array[1] = 2
	fmt.Println(array, len(array), cap(array))

	// Slice
	slice := []int{0, 1, 2, 3, 4, 5, 6}
	fmt.Println(slice, len(slice), cap(slice))

	// Métodos en el slice
	fmt.Println(slice[0])
	fmt.Println(slice[:3])
	fmt.Println(slice[2:4])
	fmt.Println(slice[4:])

	// Append
	slice = append(slice, 7)
	fmt.Println(slice)

	// Append nueva list
	newSlice := []int{8, 9, 10}
	slice = append(slice, newSlice...)
	fmt.Println(slice)
	*/

	//slice := []string{"hola", "que", "hace"}

	//for i := range slice {
	//	fmt.Println(i)
	//}

	isPalindromo("Ama")
	// ama o ama
	// amor o roma

}

func isPalindromo(text string) {
	var textReverse string
	//text = strings.ToLower(text)

	for i := len(text) - 1; i >= 0; i-- {
		textReverse += string(text[i])
	}

	if strings.EqualFold(text, textReverse) {
		fmt.Println("Es palindromo")
	} else {
		fmt.Println("No es un palindromo")
	}
}

func tripeArgument(a, b int, c string) {
	fmt.Println(a, b, c)
}

func normalFunction(message string) {
	fmt.Println(message)
}

func returnValue(a int) int {
	return a * 2
}

func doubleReturn(a int) (c, d int) {
	return a, a * 2
}

func isPar(a int) (c bool) {
	if a%2 == 0 {
		fmt.Println("Es par")
		return true
	} else {
		fmt.Println("No es par")
		return false
	}
}

func auth(user, password string) (out bool) {
	if user == "kemirandad" && password == "1234" {
		return true
	}
	return false
}
