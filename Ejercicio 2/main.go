package main

import "fmt"

func factorial(n float64) float64{
	if n == 0 {
	   return 1
	}
	return n * factorial(n-1)
 }

func main() {
	
	var infinito int = 0
	var respuesta float64= 0
	var n float64=0
	fmt.Scan(&infinito)

	for i:= 0; i <= infinito; i++ { //5
		respuesta = respuesta + (1 / factorial(n))
		fmt.Println("->", respuesta)
		n++

	}

	fmt.Println("Resultado ", respuesta)
 }
 
 