package docker

import (
	"dev/species"
	"fmt"
	"github.com/spf13/cobra"
)



var dockerInstallCmd = &cobra.Command{
	Use:   "install",
	Short: "Install docker",
	Run: func(cmd *cobra.Command, args []string) {

		species.Call(docker_install, script, func(msg string) string {
			return fmt.Sprintf(`%s - %s`, msg, "Install success, more info please refer https://docs.docker.com/engine/install/centos/")

		})

	},
}

