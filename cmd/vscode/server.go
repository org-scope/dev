package vscode

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	VscodeCmd.AddCommand(vscodeInstallCmd)
	VscodeCmd.AddCommand(vscodeServerRunCmd)
}

var VscodeCmd = &cobra.Command{
	Use:   "vscode",
	Short: "Vscode desc",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Vscode server tools")
	},
}
