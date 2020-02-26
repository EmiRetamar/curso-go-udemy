package main

import "fmt"

type Gorra struct {
	marca string
	color string
	precio float32
	plana bool
}

func main() {
	/*
	var gorra_negra = Gorra{ "Adidas", "Roja", 20.50, true}

	fmt.Println(gorra_negra.marca)
	fmt.Println(gorra_negra.precio)
	*/

 	/*
	var numero1 float32 = 10;
	var numero2 float32 = 6;

	fmt.Println("Calculadora 1");
	calculadora(numero1, numero2);
	fmt.Println("-------------");

	var numero3 float32 = 44;
	var numero4 float32 = 7;
	fmt.Println("Calculadora 2");
	calculadora(numero3, numero4);
	*/

	/*
	fmt.Print("Pedido 1 -->");
	fmt.Println(gorras(45, "€"));

	fmt.Println("-------------");

	fmt.Print("Pedido 2 -->");
	fmt.Println(gorras(24, "$"));
	*/
	
	/*
	var peliculas [3][2]string;
	peliculas[0][0] = "La verdad duele"
	peliculas[0][1] = "Mientras duermes"

	peliculas[1][0] = "Ciudano ejemplar"
	peliculas[1][1] = "El señor de los anillos"

	peliculas[2][0] = "Gran torino"
	peliculas[2][1] = "A todo gas"
	*/
	
	peliculas := []string{
		"La verdad duele", 
		"Ciudadano Ejemplar", 
		"Batman",
	    "Superman"}

	peliculas = append(peliculas, "Sin limites");
	peliculas = append(peliculas, "Camp Rock");
	peliculas = append(peliculas, "A todo gas");

	fmt.Println(peliculas[0:3])

}

func pantalon(caracteristicas ...string){
	for _, caracteristica := range caracteristicas{
		fmt.Println(caracteristica);
	}
}

func gorras(pedido float32, moneda string) (string, float32, string){

	precio := func() float32{
		return pedido*7
	}

	return "El precio del pedido:", precio(), moneda
}

func devolverTexto() (dato1 string, dato2 int) {
	dato1 = "Víctor"
	dato2 = 99

	return
}

func holaMundo(){
	fmt.Println("Hola Mundo!!");
}

func operacion(n1 float32, n2 float32, op string) float32 {
	var resultado float32

	if(op == "+"){
		resultado = n1 + n2
	}

	if(op == "-"){
		resultado = n1 - n2
	}

	if(op == "*"){
		resultado = n1 * n2
	}

	if(op == "/"){
		resultado = n1 / n2
	}

	return resultado
}

func calculadora(numero1 float32, numero2 float32){
	// Suma
	fmt.Print("La suma es: ");
	fmt.Println(operacion(numero1, numero2, "+"));

	// Resta
	fmt.Print("La resta es: ");
	fmt.Println(operacion(numero1, numero2, "-"));

	// Multiplicación
	fmt.Print("La multiplicación es: ");
	fmt.Println(operacion(numero1, numero2, "*"));

	// División
	fmt.Print("La division es: ");
	fmt.Println(operacion(numero1, numero2, "/"));
}

