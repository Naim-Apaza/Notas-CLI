package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(DeleteCmd)
}

var DeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Agrega una tarea a tu lista",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("La tarea se agrego correctamente")
	},
}
