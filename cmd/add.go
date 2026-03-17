package cmd

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(AddCmd)
}

var AddCmd = &cobra.Command{
	Use:   "add",
	Short: "Agrega una task a tu lista",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		agregartask(args[0])
	},
}

func agregartask(title string) {
	// Obtener la fecha y hora actual
	now := time.Now()
	task := []string{
		fmt.Sprintf("%d", getID()),
		title,
		"false",
		fmt.Sprintf("%d", now.Year()),
		fmt.Sprintf("%d", now.Month()),
		fmt.Sprintf("%d", now.Day()),
		fmt.Sprintf("%d", now.Hour()),
		fmt.Sprintf("%d", now.Minute()),
		fmt.Sprintf("%d", now.Second()),
	}

	// Crear (o reemplazar) el archivo CSV
	file, err := os.OpenFile("tasks.csv", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0o644)
	if err != nil {
		fmt.Println("Error al crear el archivo:", err)
		return
	}

	// Cierra el archivo al final
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	// Crear el escritor CSV
	writer := csv.NewWriter(file)

	// Escribir todas las filas
	err = writer.Write(task)
	if err != nil {
		fmt.Println("Error al escribir:", err)
		return
	}

	// Asegurarse de que se escriban los datos en disco
	writer.Flush()

	// Verificar si hubo errores al escribir
	if err := writer.Error(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("✅ Archivo CSV creado con éxito.")
}

func getID() int {
	// Abrir el archivo CSV para leer
	file, err := os.Open("tasks.csv")
	if err != nil {
		log.Fatal("Error al abrir el archivo")
	}

	// Cierra el archivo al final
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal("Error al cerrar el archivo")
		}
	}()

	// Leer todas las filas del CSV
	r := csv.NewReader(file)
	rows, err := r.ReadAll()
	if err != nil {
		log.Fatal("Error al leer el CSV")
	}

	// Si no hay filas, el ID inicial es 1
	if len(rows) == 0 {
		return 1
	}

	// Obtener el ID de la última fila y sumarle 1
	lastRow := rows[len(rows)-1]
	lastID, err := strconv.Atoi(lastRow[0])
	if err != nil {
		log.Fatal("Error al convertir el ID a entero")
	}

	return lastID + 1
}
