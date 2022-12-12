package vscode

import (
	"dev/species"
	"fmt"
	"github.com/spf13/cobra"
)



var vscodeInstallCmd = &cobra.Command{
	Use:   "install",
	Short: "Install vscode server",
	Run: func(cmd *cobra.Command, args []string) {

		species.Call(vscode_install, scriptInstall, func(msg string) string {
			return fmt.Sprintf(`%s - %s`, msg, "Install success, more info please refer https://github.com/coder/code-server")
		})
	},
}
