package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	add()
	list()
}

func add() {
	file, err := os.OpenFile("notas.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal("Error al abrir el archivo", err)
		return
	}
	defer file.Close()

	// Escribe en el archivo
	fmt.Println("Escriba una nota!")
	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		fmt.Println("No se pudo leer la nota")
		return
	}

	nota := scanner.Text()

	if _, err = file.WriteString(nota + "\n"); err != nil {
		fmt.Println("Oh a ocurrido un error al escribir en la nota")
		return
	}

	fmt.Println("Exito al escribir la nota")
}

func list() {
	file, err := os.OpenFile("notas.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal("Error al abrir el archivo", err)
		return
	}
	defer file.Close()

	cmd := exec.Command("bat", file.Name())

	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin

	if err := cmd.Run(); err != nil {
		fmt.Println("Error al ejecutar el comando", err)
		return
	}
}
