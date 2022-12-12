package kubernetes

import (
    "dev/species"
    "fmt"
    "strings"

    "github.com/spf13/cobra"
)

func init() {
    kubernetesInstallCmd.Flags().StringVarP(&Masters, "masters", "m", "", "Master IP or IP Range")
    _ = kubernetesInstallCmd.MarkFlagRequired("masters")
    kubernetesInstallCmd.Flags().StringVarP(&Nodes, "nodes", "n", "", "Node IP or IP Range")
    _ = kubernetesInstallCmd.MarkFlagRequired("nodes")
    kubernetesInstallCmd.Flags().StringVarP(&Password, "password", "p", "", "SSH Password")
    _ = kubernetesInstallCmd.MarkFlagRequired("password")
}

var kubernetesInstallCmd = &cobra.Command{
    Use:   "install",
    Short: "Install kubernetes",
    Run: func(cmd *cobra.Command, args []string) {
        script := strings.Replace(script, "{MASTERS}", Masters, -1)
        script = strings.Replace(script, "{NODES}", Nodes, -1)
        script = strings.Replace(script, "{PASSWORD}", Password, -1)
		fmt.Println(Masters, Nodes, Password)
        species.Call(k8s_install, script, func(msg string) string {
            return fmt.Sprintf(`%s %s`, msg, "Install success, more info please refer https://www.sealos.io/docs/getting-started/installation")
        })

    },
}
