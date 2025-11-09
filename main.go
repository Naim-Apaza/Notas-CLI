package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	seleccionarOpcion()
}

func seleccionarOpcion() {
	if len(os.Args) < 2 {
		fmt.Println("Debe poner alguna opcion")
		fmt.Println("Use go run main.go [add|list|remove]")
		os.Exit(1)
	}

	comando := os.Args[1]

	switch comando {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("El comando add requiere una tarea")
			fmt.Println("Ejemplo: go run main.go add \"Comprar leche\"")
			os.Exit(1)
		}
		tarea := os.Args[2]
		add(tarea)
	case "list":
		list()
	case "remove":
		remove()
	default:
		fmt.Printf("❌ Error: Comando desconocido: **%s**\n", comando)
		fmt.Println("Comandos soportados: add, list, remove")
		os.Exit(1)
	}
}

func add(tarea string) {
	file, err := os.OpenFile("notas.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal("Error al abrir el archivo", err)
		return
	}

	defer func() {
		if err := file.Close(); err != nil {
			log.Println("Error al cerrar el archivo", err)
		}
	}()

	// Escribe en el archivo
	// fmt.Println("Escriba una nota!")
	// scanner := bufio.NewScanner(os.Stdin)
	// if !scanner.Scan() {
	// 	fmt.Println("No se pudo leer la nota")
	// 	return
	// }

	// nota := scanner.Text()

	if _, err = file.WriteString(tarea + "\n"); err != nil {
		fmt.Println("Oh a ocurrido un error al escribir en la nota")
		return
	}

	fmt.Println("Exito al escribir la nota")
}

func remove() {
	file, err := os.OpenFile("notas.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		println("Oh a ocurrido un error al abrir el archivo")
		return
	}

	defer func() {
		if err := file.Close(); err != nil {
			log.Println("Error al cerrar el archivo", err)
		}
	}()

}

func list() {
	file, err := os.OpenFile("notas.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error al abrir el archivo", err)
		return
	}

	defer func() {
		if err := file.Close(); err != nil {
			fmt.Println("Error al cerrar el archivo", err)
		}
	}()

	cmd := exec.Command("bat", file.Name())
	if err = cmd.Run(); err != nil {
		fmt.Println("Oh a ocurrido en error al ejecutar el comando", err)
		return
	}
}
