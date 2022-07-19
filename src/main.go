package main

import (
	"fmt"
	"strings"
	"sync"
)

// INTERFACES
type figures2d interface {
	area() float64
}

// STRUCTS
type car struct {
	brand string
	year  uint
}

type computer struct {
	ram   uint
	disk  uint
	brand string
}

// Métodos del struct computer
func (pc computer) ping() {
	fmt.Println(pc.brand, "Pong")
}

// func (pc computer) duplicateRAM() { // * Con esto solo ejecutara la función pero no modificara la propiedad de los datos instanciados. Es útil para retornar valores más no modificar la instancia
func (pc *computer) duplicateRAM() { // * Con  esto podemos alterar la propiedad de la instancia
	pc.ram = pc.ram * 2
}

func (pc computer) String() string { // * Con este método modificas la forma en cómo se muestra el struct al momento de llamarlo con un fmt.Println
	return fmt.Sprintf("Tengo %d GB de RAM, %d GB de Disco y soy marca %s", pc.ram, pc.disk, pc.brand)
}

type square struct {
	base float64
}

func (self square) area() float64 {
	return self.base * self.base
}

type rectangule struct {
	base   float64
	height float64
}

// * FUNCIÓN IMPLEMENTANDO LA INTERFAZ
func calculateArea(figure figures2d) {
	fmt.Println("Area:", figure.area())
}

func (self rectangule) area() float64 {
	return self.base * self.height
}

func normalFunction(msg string) {
	fmt.Println(msg)
}

// func tripleArgument(a int, b int, c string) {
func tripleArgument(a, b int, c string) { // * a y b son ambos int
	fmt.Println(a, b, c)
}

func returnValue(a int) int {
	return a * 2
}

func doubleReturn(a int) (b, c int) { // * se pueden retornar 2 valores
	return a, a * 2
}

func isPalindrome(text string) bool {
	var textReverse string
	textToLower := strings.ToLower(text)

	for i := len(text) - 1; i >= 0; i-- {
		textReverse += string(textToLower[i])
	}

	return textToLower == textReverse
}

// GOROUTINES (functions) wg
func say(text string, wg *sync.WaitGroup) {
	defer wg.Done() // * Indicamos a la WG que terminó la Goroutine. Por detrás decrementar el contador de las WG
	// * Defer para ejecutarlo al final de la función
	fmt.Println(text)
}

// GOROUTINES (functions) channel
// func say(text string, c chan<- string) {
// 	// * <-chan salida de datos
// 	// * chan<- entrada de datos
// 	c <- text
// }

func main() {
	// CONSTANTES
	// const pi float64 = 3.14
	// const pi2 = 3.1415

	// fmt.Println("pi:", pi)
	// fmt.Println("pi2:", pi2)

	// ENTEROS
	// base := 12 // * Se crea sin indicar el tipo de dato y se asigna el valor.
	// base = 12  // * Esto es solo asignación.
	// var height int = 14 // * Declaración y asignación.
	// var area int        // * Solamente declaración.

	// fmt.Println(base, height, area)

	// ZERO VALUES
	// var a int
	// var b float64
	// var c string
	// var d bool

	// fmt.Println(a, b, c, d)

	// Area
	// const squareBase = 10
	// squareArea := squareBase * squareBase
	// fmt.Println("Area de un cuadrado es:", squareArea)

	// OPERADORES ARITMETICOS
	// Suma
	// x := 10
	// y := 50
	// result := x + y
	// fmt.Println("Suma:", result)

	// Resta
	// result = y - x
	// fmt.Println("Resta:", result)

	// Multiplicación
	// result = y * x
	// fmt.Println("Multiplicación:", result)

	// División
	// result = y / x
	// fmt.Println("División:", result)

	// Resto
	// result = y % x
	// fmt.Println("Resto:", result)

	// Incrementar
	// x++
	// fmt.Println("Incrementar:", x)

	// Decrementar
	// x--
	// fmt.Println("Decrementar:", x)

	// TIPOS DE DATOS PRIMITIVOS
	// Numeros enteros
	// int = Depende del OS (32 o 64 bits)
	// int8 = 8bits = -128 a 127
	// int16 = 16bits = -2^15 a 2^15-1
	// int32 = 32bits = -2^31 a 2^31-1
	// int64 = 64bits = -2^63 a 2^63-1

	// Optimizar memoria cuando sabemos que el dato simpre va ser positivo
	//	// uint = Depende del OS (32 o 64 bits)
	//	// uint8 = 8bits = 0 a 127
	//	// uint16 = 16bits = 0 a 2^15-1
	//	// uint32 = 32bits = 0 a 2^31-1
	//	// uint64 = 64bits = 0 a 2^63-1

	//	// numeros decimales
	//	// float32 = 32 bits = +/- 1.18e^-38 +/- -3.4e^38
	//	// float64 = 64 bits = +/- 2.23e^-308 +/- -1.8e^308

	//	// textos y booleanos
	//	// string = ""
	//	// bool = true or false

	//	// numeros complejos
	//	// Complex64 = Real e Imaginario float32
	//	// Complex128 = Real e Imaginario float64
	//	// Ejemplo : c:=10 + 8i

	// FMT
	// Println
	// helloMessage := "Hello"
	// worldMessage := "World"

	// fmt.Println(helloMessage, worldMessage)

	// Printf
	// name := "Platzi"
	// courses := 500
	// fmt.Printf("%s tiene más de %d cursos\n", name, courses)
	// fmt.Printf("%v tiene más de %d cursos\n", name, courses) // * v indica que no conocemos el tipo de dato que pueda estár allí

	// Slice String
	// fmt.Println(string(name[0])) // * Al solicitar un caracter según su índice, se regresa en valor decimal, por eso es necesario usar la función string para convertirlo
	// fmt.Println(name[:3])        // Mostrar un caracter del string
	// fmt.Println(name[:3])
	// fmt.Println(name[2:4])
	// fmt.Println(name[4:])
	// fmt.Println(name[4:10]) // * Esto generaría un panic porque se sale del rango solicitado

	// Sprintf
	// message := fmt.Sprintf("%s tiene más de %d cursos", name, courses)
	// fmt.Println(message)

	// Indicar tipo de dato de la variable
	// fmt.Printf("helloMessage: %T\ncursos: %T\n", helloMessage, courses)

	// FUNCIONES
	// normalFunction("epa")
	// tripleArgument(1, 2, "hola")

	// value := returnValue(2)
	// fmt.Println("Value:", value)

	// value1, value2 := doubleReturn(2) // * esta función retorna 2 valores
	// value1, _ := doubleReturn(2) // * así podemos usar un solo valor en vez de los dos
	// fmt.Println("Value1:", value1, "Value2:", value2)

	// CICLOS
	// For condicional
	// for i := 1; i < 10; i++ {
	// 	fmt.Println(i)
	// }

	// For while
	// counter := 1
	// for counter <= 10 {
	// 	fmt.Println(counter)
	// 	counter++
	// }

	// For forever
	// counterForever := 1
	// for {
	// 	fmt.Println(counterForever)
	// 	if counterForever == 10 { // * Si no hago esto se ejecuta el ciclo por siempre
	// 		break
	// 	}
	// 	counterForever++
	// }

	// CONDICIONALES
	// If
	// const value = 1
	// const value2 = 2

	// if value == 1 {
	// 	fmt.Println("Es uno.")
	// } else {
	// 	fmt.Println("No es uno.")
	// }

	// if value == 1 && value2 == 2 {
	// 	fmt.Println("Es verdad, AND")
	// }

	// if value == 1 || value2 == 2 {
	// 	fmt.Println(" Es verdad, OR")
	// }

	// value, err := strconv.Atoi("53")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("Value:", value)

	// Switch
	// switch module := 5 % 2; module {
	// case 0:
	// 	fmt.Println("Es par")
	// default:
	// 	fmt.Println("Es impar")
	// }
	// Switch sin condición
	// value := 200
	// switch {
	// case value > 100:
	// 	fmt.Println("Es mayor a 100")
	// case value < 0:
	// 	fmt.Println("Es menor a 0")
	// default:
	// 	fmt.Println("No condición")
	// }

	// DEFER
	// defer fmt.Println("Hola") // * Con defer la función se ejecuta al terminar la función
	// fmt.Println("Mundo")
	// * Se puede usar más de un defer dentro de una función, pero cómo buena práctica es debemos usar uno solo por función

	// CONTINUE
	// for i := 1; i < 10; i++ {
	// 	if i == 2 {
	// 		continue // * Sirve para continuar hacia el siguiente ciclo y evitar terminar la ejecución
	// 	}
	// 	fmt.Println(i)
	// }

	// BREAK
	// for i := 1; i < 10; i++ {
	// 	if i == 8 {
	// 		break // * Sirve para culminar el ciclo por completo
	// 	}
	// 	fmt.Println(i)
	// }

	// ARRAYS
	// var array [4]int
	// * Los arrays tienen un tamaño definido y no puede varias, lo que puede variar es el valor de sus elementos
	// array[0] = 1
	// array[1] = 2
	// fmt.Println(array)
	// fmt.Println(len(array)) // * Tamaño actual del array
	// fmt.Println(cap(array)) // * Tamaño máximo posible del array

	// * También se puede declarar y asignar un array directamente
	// var array [2]string = [2]string{"Hola", "Mundo"}
	// array := [2]string{"Hola", "Mundo"}
	// fmt.Println(array)

	// SLICE
	// slice := []int{0, 1, 2, 3, 4, 5, 6}
	// fmt.Println(len(slice))
	// fmt.Println(cap(slice))

	// Métodos
	// fmt.Println(slice[0])
	// fmt.Println(slice[:3])
	// fmt.Println(slice[2:4])
	// fmt.Println(slice[4:])
	// fmt.Println(slice[4:10]) // * Esto generaría un panic porque se sale del rango solicitado

	// Append
	// slice = append(slice, 7)
	// fmt.Println(slice)
	// Append nueva lista
	// newSlice := []int{8, 9, 10}
	// slice = append(slice, newSlice...) // * newSlice... es la forma de mandar una destructuración de un array
	// fmt.Println(slice)

	// Recorriendo slices con range
	// slice := []string{"Hola,", "que", "hace?"}
	// for i, value := range slice {
	// 	fmt.Println(i, value)
	// }

	// fmt.Println(isPalindrome("Ama"))

	// MAPS
	// m := make(map[string]int)
	// m["Mario"] = 25
	// m["Jesus"] = 33
	// fmt.Println(m)

	// Otra forma de declarar un map
	// m := map[string]int{
	// 	"Mario": 25,
	// 	"Jesus": 33,
	// }
	// fmt.Println(m)

	// Recorriendo map con range
	// for k, v := range m { // * k = key v = value
	// 	fmt.Println(k, v)
	// }

	// Obteniendo un valor directamente
	// value := m["Mario"]
	// fmt.Println(value)
	// value = m["Mari"] // * No existe la clave Mari, por ende nos devolverá un Zero Value
	// fmt.Println(value)
	// value2, ok := m["Mari"] // * Con el ok podemos verficar si la llave existe
	// fmt.Println(value2, ok)

	// * DATO MEGA IMPORTANTE:
	/*
		La estructura de datos MAP son más eficientes que los Array o Slices,
		ya que usan de forma nativa concurrencia para ejecutar las operaciones.
	*/

	// STRUCTS
	// myCar := car{brand: "Ford", year: 2020} // * Instance car
	// fmt.Println(myCar)

	// var otherCar car
	// otherCar.brand = "Toyota"
	// fmt.Println(otherCar)

	// PUNTEROS
	// a := 50
	// b := &a // * Con el & se accede a la dirección de memoria

	// fmt.Println(a)
	// fmt.Println(b)
	// fmt.Println(*b) // * Con el * se accede al valor de la dirección de memoria (&)

	// *b = 100 // * Debido a que a y b está apuntando a la misma dirección de memoria, alterar uno alterará a el otro
	// fmt.Println(a)

	// myPC := computer{ram: 16, disk: 200, brand: "Lenovo"}
	// fmt.Println(myPC)

	// myPC.ping()

	// myPC.ram = 20 // * Esta es otra forma de alterar una propiedad de una instancia

	// fmt.Println(myPC)
	// myPC.duplicateRAM()
	// fmt.Println(myPC)
	// myPC.duplicateRAM()
	// fmt.Println(myPC)

	// * & es para pasar una referencia
	// * * es para pasar una desreferencia

	// STRINGERS
	// myPc2 := computer{ram: 16, brand: "Mac", disk: 100}
	// fmt.Println(myPc2)

	// mySquare := square{base: 2}
	// myRect := rectangule{base: 2, height: 4}

	// * Forma sin implementar una función con interfaz
	// squareArea := mySquare.area()
	// rectArea := myRect.area()
	// fmt.Println("Área del cuadrado:", squareArea)
	// fmt.Println("Área del rectángulo:", rectArea)

	//  * Forma implementando una función con interfaz
	// calculateArea(mySquare)
	// calculateArea(myRect)
	// * Estructuras de datos que comparten mismos comportamientos deben tener interfaces para ser más eficientes

	// LISTA DE INTERFACES
	// myInterface := []interface{}{"Hola", 12, 4.9} // * interface{} indica una interfaz de cualquier tipo
	// fmt.Println(myInterface)

	// GOROUTINES
	// Wait Group
	var wg sync.WaitGroup
	// * sync permite interactuar de forma primitiva con la Goroutines, lo cuál lo hace más eficiente, pero a costa de que es más dificil mantener
	// * Wait Group permite almacenar goroutines que el hilo principal esperará a que se terminen de ejecutar

	fmt.Println("Hello")
	wg.Add(2) // * Agregamos una Goroutine en el ciclo del wait group
	go say("World", &wg)
	// time.Sleep(time.Second * 1) // * Funciona pero no es la mejor práctica

	go func(text string, wg *sync.WaitGroup) { // * Goroutine con función anónima
		defer wg.Done()
		fmt.Println(text)
	}("Adios", &wg)
	wg.Wait() // * Esperamos a que cada Goroutine termine

	// Channels
	// c := make(chan string, 1) // * Cantidad de Goroutines que recibirá el channel
	// fmt.Println("Hello")
	// go say("Bye", c)
	// fmt.Println(<-c)

	// * RECOMENDACIÓN EN CUANTO A CONCURRENCIA: se recomienda el uso de los WaitGroup debido a la eficiencia que aporta.
}
