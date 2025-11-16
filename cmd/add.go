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
	Short: "Agrega una tarea a tu lista",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		agregarTarea(args)
	},
}

func agregarTarea(args []string) {
	id, err := obtenerSiguienteID()
	if err != nil {
		fmt.Println("Error al obtener el id", err)
		return
	}
	ah := time.Now()

	tarea := []string{strconv.Itoa(id), args[0], "false", ah.Format(time.TimeOnly)}

	// Crear (o reemplazar) el archivo CSV
	file, err := os.OpenFile("tareas.csv", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println("Error al crear el archivo:", err)
		return
	}

	//Cierra el archivo al final
	defer func() {
		if err := file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	// Crear el escritor CSV
	writer := csv.NewWriter(file)

	// Escribir todas las filas
	err = writer.Write(tarea)
	if err != nil {
		fmt.Println("Error al escribir:", err)
		return
	}

	// Asegurarse de que se escriban los datos en disco
	writer.Flush()

	if err := writer.Error(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("✅ Archivo CSV creado con éxito.")

}

func obtenerSiguienteID() (int, error) {
	file, err := os.Open("tareas.csv")
	if err != nil {
		// Si no existe el archivo, empezamos desde 1
		if os.IsNotExist(err) {
			return 1, nil
		}
		return 0, err
	}

	defer func() {
		if err := file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	reader := csv.NewReader(file)
	rows, err := reader.ReadAll()
	if err != nil {
		return 0, err
	}

	// Si el archivo existe pero está vacío
	if len(rows) == 0 {
		return 1, nil
	}

	// Obtener la última fila
	ultima := rows[len(rows)-1]
	lastID, err := strconv.Atoi(ultima[0])
	if err != nil {
		return 0, err
	}

	return lastID + 1, nil
}
