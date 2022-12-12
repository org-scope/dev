package vscode

import (
	"dev/species"
	"fmt"
	"github.com/creack/pty"
	"github.com/spf13/cobra"
	"io"
	"os"
	"os/exec"
	"path"
)



var vscodeInstallCmd = &cobra.Command{
	Use:   "install",
	Short: "Install vscode server",
	Run: func(cmd *cobra.Command, args []string) {

		homeDir, _ := os.UserHomeDir()
		_ = os.MkdirAll(path.Join(homeDir, species.Workspace), os.ModePerm)
		fp := path.Join(homeDir, species.Workspace, vscode_install)

		create, _ := os.Create(fp)
		_, _ = create.WriteString(scriptInstall)

		command := exec.Command(`bash`, fp)
		start, err := pty.Start(command)
		if err != nil {
			fmt.Println(`error : `, err)
		}
		_, _ = io.Copy(os.Stdout, start)
	},
}
