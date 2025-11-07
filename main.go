package main

import (
	"fmt"
	"os"
)

func main() {
	a, err := os.Create("notas.txt")

	if err != nil {
		fmt.Printf("error al crear el archivo")
		return
	}

	data := "Aprendiendo a escribir en un archivo"

	_, err = a.WriteString(data)
	if err != nil {
		fmt.Printf("Error al escribir el archivo %v", err)
		return
	}

	data = "\nEscribo otra cosa, donde se escribira? abajo o al lado?"
	_, err = a.WriteString(data)
	if err != nil {
		fmt.Printf("Error al escribir el archivo %v", err)
		return
	}

	fmt.Printf("Se creo el archivo %v y se escribio", a.Name())
}
