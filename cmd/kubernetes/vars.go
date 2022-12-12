package kubernetes

const (
    k8s_install = `k8s_install.sh`
)

var (
    script = `#!/bin/bash
wget https://github.com/labring/sealos/releases/download/v4.1.3/sealos_4.1.3_linux_amd64.tar.gz
tar -zxvf sealos_4.1.3_linux_amd64.tar.gz sealos
chmod +x sealos
mv sealos /usr/bin

sealos run labring/kubernetes:v1.24.0 labring/calico:v3.22.1 \
    --masters {MASTERS} \
    --nodes {NODES} \
	-p {PASSWORD}`

    Masters  string
    Nodes    string
    Password string
)
