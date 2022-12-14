package cmd

import (
	"dev/cmd/docker"
	golang "dev/cmd/go"
	"dev/cmd/kubernetes"
	"dev/cmd/vscode"
	"dev/species"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"path"
)

func init() {
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(vscode.VscodeCmd)
	rootCmd.AddCommand(docker.DockerCmd)
	rootCmd.AddCommand(golang.GolangCmd)
	rootCmd.AddCommand(kubernetes.KubernetesCmd)

}

var rootCmd = &cobra.Command{
	Use:   "dev",
	Short: "dev tools for devops",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(`./dev help`)

		homeDir, _ := os.UserHomeDir()
		workspace := path.Join(homeDir, species.Workspace)
		_ = os.MkdirAll(workspace, os.ModePerm)

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
