package cmd

import (
	"fmt"

	sim "github.com/kyriosdata/assinatura/internal/simulador"
	"github.com/spf13/cobra"
)

var simPort int

var simuladorCmd = &cobra.Command{
	Use:   "simulador",
	Short: "Gerencia o simulador.jar do HubSaúde",
}

var simStartCmd = &cobra.Command{
	Use:   "start",
	Short: "Inicia o simulador",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := sim.Start(simPort); err != nil {
			return err
		}
		fmt.Println("[OK] Simulador iniciado")
		return nil
	},
}

var simStopCmd = &cobra.Command{
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

var simStatusCmd = &cobra.Command{
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

func init() {
	simStartCmd.Flags().IntVar(&simPort, "port", 8443, "Porta do simulador")

	simuladorCmd.AddCommand(simStartCmd)
	simuladorCmd.AddCommand(simStopCmd)
	simuladorCmd.AddCommand(simStatusCmd)

	rootCmd.AddCommand(simuladorCmd)
}
