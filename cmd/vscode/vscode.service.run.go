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



var vscodeServerRunCmd = &cobra.Command{
	Use:   "run",
	Short: "Run vscode",
	Run: func(cmd *cobra.Command, args []string) {

		homeDir, _ := os.UserHomeDir()
		_ = os.MkdirAll(path.Join(homeDir, species.Workspace), os.ModePerm)

		filepath := path.Join(homeDir, species.Workspace, vscode_run)
		fp, _ := os.Create(filepath)

		_, _ = fp.WriteString(scriptRun)

		command := exec.Command("bash", filepath)
		f, err := pty.Start(command)
		if err != nil {
			fmt.Println(`error:`, err)
		}
		_, _ = io.Copy(os.Stdout, f)
	},
}
