package cmd

import (
	"fmt"

	"github.com/kyriosdata/assinatura/internal/invoker"
	"github.com/spf13/cobra"
)

var signContent string
var signToken string

var signCmd = &cobra.Command{
	Use:   "sign",
	Short: "Cria uma assinatura digital simulada",
	Long: `Invoca o assinador.jar para criar uma assinatura digital simulada para o conteúdo fornecido.

Exemplos:
  assinatura sign --content "documento a assinar"
  assinatura sign --content "documento a assinar" --token "meu-token"`,
	RunE: func(cmd *cobra.Command, args []string) error {
		resp, err := invoker.Sign(signContent, signToken)
		if err != nil {
			return err
		}
		fmt.Println(resp.Format())
		return nil
	},
}

func init() {
	signCmd.Flags().StringVar(&signContent, "content", "", "Conteúdo a ser assinado (obrigatório)")
	signCmd.Flags().StringVar(&signToken, "token", "", "Token de autenticação (opcional)")
	signCmd.MarkFlagRequired("content")
	rootCmd.AddCommand(signCmd)
}
