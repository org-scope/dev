package golang

import (
	"dev/species"
	"fmt"
	"github.com/spf13/cobra"
	"strings"
)

func init() {
	golangInstallCmd.Flags().StringVarP(&Version, "version", "v", "1.19.1", "Golang Version")
}




var golangInstallCmd = &cobra.Command{
	Use:   "install",
	Short: "Install golang",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(Version)
		script := strings.Replace(script, "{VERSION}", Version, -1)
		species.Call(go_install, script, func(msg string) string {
			return fmt.Sprintf(`%s %s`, msg, `Please run 'source ~/.bashrc' to set the env"`)
		})
	},
}
