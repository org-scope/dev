package golang

const (
    go_install = `go_install.sh`
)

var (
    script = `#!/bin/bash
wget https://dl.google.com/go/go{VERSION}.linux-amd64.tar.gz

tar -xvf go{VERSION}.linux-amd64.tar.gz

cp -r ./go /usr/local/go

cat>>~/.bashrc<<EOF
export PATH=$PATH:/usr/local/go/bin
EOF

source ~/.bashrc`

    Version string
)
