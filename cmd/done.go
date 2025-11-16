package cmd

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(DoneCmd)
}

var DoneCmd = &cobra.Command{
	Use:   "done",
	Short: "Marca una tarea como hecha",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if err := completarTarea(args[0]); err != nil {
			log.Fatal(err)
		}
		fmt.Println("Tarea marcada como completada.")
	},
}

func completarTarea(id string) error {
	// 1. Abrir para leer
	file, err := os.Open("tareas.csv")
	if err != nil {
		return fmt.Errorf("error al abrir el archivo: %v", err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			log.Fatal("Error al cerrar el archivo")
		}
	}()

	reader := csv.NewReader(file)
	rows, err := reader.ReadAll()
	if err != nil {
		return fmt.Errorf("error al leer archivo: %v", err)
	}

	// 2. Buscar y modificar
	found := false
	for i, row := range rows {
		if row[0] == id {
			rows[i][2] = strconv.FormatBool(true)
			found = true
			break
		}
	}

	if !found {
		return fmt.Errorf("no se encontró la tarea con ID %s", id)
	}

	// 3. Reescribir archivo completo
	file2, err := os.Create("tareas.csv")
	if err != nil {
		return fmt.Errorf("error al reemplazar archivo: %v", err)
	}
	defer func() {
		if err := file2.Close(); err != nil {
			log.Fatal("Error al cerrar el archivo")
		}
	}()

	writer := csv.NewWriter(file2)
	defer writer.Flush()

	for _, row := range rows {
		if err := writer.Write(row); err != nil {
			return fmt.Errorf("error al escribir archivo: %v", err)
		}
	}

	return nil
}
