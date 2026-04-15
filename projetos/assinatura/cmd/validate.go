package cmd

import (
	"fmt"

	"github.com/kyriosdata/assinatura/internal/invoker"
	"github.com/spf13/cobra"
)

var validateContent string
var validateSignature string

var validateCmd = &cobra.Command{
	Use:   "validate",
	Short: "Valida uma assinatura digital simulada",
	Long: `Invoca o assinador.jar para verificar se uma assinatura é válida para o conteúdo fornecido.

Exemplos:
  assinatura validate --content "documento" --signature "MOCKED_SIGNATURE_BASE64_=="`,
	RunE: func(cmd *cobra.Command, args []string) error {
		resp, err := invoker.Validate(validateContent, validateSignature)
		if err != nil {
			return err
		}
		fmt.Println(resp.Format())
		return nil
	},
}

func init() {
	validateCmd.Flags().StringVar(&validateContent, "content", "", "Conteúdo original que foi assinado (obrigatório)")
	validateCmd.Flags().StringVar(&validateSignature, "signature", "", "Assinatura a ser validada (obrigatório)")
	validateCmd.MarkFlagRequired("content")
	validateCmd.MarkFlagRequired("signature")
	rootCmd.AddCommand(validateCmd)
}
