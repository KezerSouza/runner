package cmd

import (
	"fmt"

	sim "github.com/kyriosdata/assinatura/internal/simulador"
	"github.com/spf13/cobra"
)

// Porta usada no comando start
var port int

// Comando principal: assinatura simulador
var simuladorCmd = &cobra.Command{
	Use:   "simulador",
	Short: "Gerencia o simulador.jar",
}

// Inicia o simulador
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Inicia o simulador",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := sim.Start(port); err != nil {
			return err
		}

		fmt.Println("[OK] Simulador iniciado")
		return nil
	},
}

// Encerra o simulador
var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Encerra o simulador",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := sim.Stop(); err != nil {
			return err
		}

		fmt.Println("[OK] Simulador encerrado")
		return nil
	},
}

// Exibe status atual
var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Exibe status do simulador",
	RunE: func(cmd *cobra.Command, args []string) error {
		running, pid := sim.Status()

		if running {
			fmt.Printf("[OK] Simulador ativo (PID %d)\n", pid)
		} else {
			fmt.Println("[INFO] Simulador parado")
		}

		return nil
	},
}

// Registra comandos no CLI principal
func init() {
	startCmd.Flags().IntVar(&port, "port", 8443, "Porta do simulador")

	simuladorCmd.AddCommand(startCmd)
	simuladorCmd.AddCommand(stopCmd)
	simuladorCmd.AddCommand(statusCmd)

	rootCmd.AddCommand(simuladorCmd)
}
