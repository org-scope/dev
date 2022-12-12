package vscode

const (
    vscode_install = `vscode_install.sh`
    vscode_run = `vscode-run.sh`
)

var (
    scriptInstall = `#!/bin/bash
curl -fsSL https://code-server.dev/install.sh | sh`


    scriptRun = `#!/bin/bash
PASSWORD=$(uuidgen)

mkdir -p ~/.config/vscode
cat>~/.config/vscode/config.yaml<<EOF
bind-addr: 0.0.0.0:80
auth: password
password: ${PASSWORD}
cert: false
EOF

echo "Please use ${PASSWORD} to login"
code-server`
)
