package main

import "fmt"
import "time"

func main(){
	fmt.Println("***** MI PROGRAMA CON GO ****")

	/*
	fmt.Println("Hola "+os.Args[1]+" bienvenido al programa en GO")

	edad,_ := strconv.Atoi(os.Args[2])

	if edad >= 18 && edad != 25 && edad != 99 {
		fmt.Println("Eres mayor de edad")
	}else if edad == 25 {
		fmt.Println("Tienes 25 años")
	}else if edad == 99 {
		fmt.Println("Pronto morirás")
	}else{
		fmt.Println("Eres menor de edad")
	}
	*/

	/*
	numero,_ := strconv.Atoi(os.Args[1])

	if numero%2 == 0{
		fmt.Println("El número es par")
	}else{
		fmt.Println("El número es IMPAR")
	}
	*/

	//peliculas := []string{"Pelicula 1", "El club de la lucha", "A todo gas", "Gran torino"}

	/*
	for i := 0; i < len(peliculas); i++{
		if i%2 == 0{
			fmt.Println("La pelicula -> "+peliculas[i]+" es par:", i)
		}else{
			fmt.Println("El pelicula -> "+peliculas[i]+" es IMPAR:", i)
		}
	}
	*/

	/*
	for _, pelicula := range peliculas {
		fmt.Println(pelicula);
	}
	*/

	momento := time.Now()
	hoy := momento.Weekday();

	switch hoy {
		case 0:
			fmt.Println("Hoy es domingo")

		case 1:
			fmt.Println("Hoy es Lunes")

		case 2:
			fmt.Println("Hoy es Martes")

		case 3:
			fmt.Println("Hoy es Miercoles")

		default:
			fmt.Println("Hoy es otro dia de la semana")
	}

}