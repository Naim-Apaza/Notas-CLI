package cmd

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"text/tabwriter"
	"time"

	"github.com/mergestat/timediff"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(ListCmd)
}

var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "Muestra las tareas de tu lista",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		mostrarLista()
	},
}

func mostrarLista() {
	file, err := os.Open("tareas.csv")
	if err != nil {
		log.Fatal("Error al abrir el archivo")
	}
	defer func() {
		if err := file.Close(); err != nil {
			log.Fatal("Error al cerrar el archivo")
		}
	}()

	r := csv.NewReader(file)
	rows, err := r.ReadAll()
	if err != nil {
		log.Fatal("Error al leer el CSV")
	}

	padding := 3
	w := tabwriter.NewWriter(os.Stdout, 0, 0, padding, ' ', 0)

	if _, err := fmt.Fprintln(w, "ID\tDescripción\tCompletado\tCreado\t"); err != nil {
		log.Fatal("Error al escribir en tabwriter")
	}

	for i := range rows {
		id := rows[i][0]
		desc := rows[i][1]
		comp := rows[i][2]
		fecha := rows[i][3]

		ParseTime, err := time.Parse(time.TimeOnly, strings.TrimSpace(fecha))
		if err != nil {
			log.Fatal("Error al convertir la fecha")
		}
		now := time.Now()

		t := time.Date(
			now.Year(),
			now.Month(),
			now.Day(),
			ParseTime.Hour(),
			ParseTime.Minute(),
			ParseTime.Second(),
			0,
			now.Location(),
		)

		hecho := ""
		s, err := strconv.ParseBool(comp)
		if err != nil {
			log.Fatal("Error al convertir a booleano")
		}

		if s {
			hecho = "Hecho"
		} else {
			hecho = "Pendiente"
		}

		if _, err := fmt.Fprintf(w, "%s\t%s\t%s\t%v\t\n", id, desc, hecho, timediff.TimeDiff(t)); err != nil {
			log.Fatal("Error al escribir en tabwriter")
		}

	}

	if err := w.Flush(); err != nil {
		log.Fatal("error al escribir el archivo")
	}

}
