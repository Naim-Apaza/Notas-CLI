package cmd

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
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
	file, err := os.Open("tasks.csv")
	if err != nil {
		log.Fatal("Error al abrir el archivo")
	}
	defer func() {
		if err = file.Close(); err != nil {
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

	if _, err := fmt.Fprintln(w, "ID\tDescription\tDone?\tDate"); err != nil {
		log.Fatal("Error al escribir en tabwriter")
	}

	for i := range rows {
		if len(rows[i]) < 9 {
			continue
		}

		id := rows[i][0]
		desc := rows[i][1]
		comp := rows[i][2]

		y, _ := strconv.Atoi(rows[i][3])
		mo, _ := strconv.Atoi(rows[i][4])
		d, _ := strconv.Atoi(rows[i][5])
		h, _ := strconv.Atoi(rows[i][6])
		m, _ := strconv.Atoi(rows[i][7])
		s, _ := strconv.Atoi(rows[i][8])

		now := time.Now()
		t := time.Date(y, time.Month(mo), d, h, m, s, 0, now.Location())

		hecho := "Pendiente"
		if comp == "true" {
			hecho = "Hecho"
		}

		// CORRECCIÓN AQUÍ: 5 columnas en el Header = 5 variables en el Fprintf
		// ID, Description, Done?, Date, Hour
		if _, err := fmt.Fprintf(w, "%s\t%s\t%s\t%s\t\n", id, desc, hecho, timediff.TimeDiff(t)); err != nil {
			log.Fatal("Error al escribir en tabwriter")
		}
	}

	if err := w.Flush(); err != nil {
		log.Fatal("Error al formatear la salida")
	}
}
