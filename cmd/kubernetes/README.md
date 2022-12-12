# Kubernetes tools

## Usage

- Install kubernetes

``` bash
# Single ip
./dev k8s install --masters x.x.x.x --nodes x.x.x.x -p {password of all machines}

# Ip array
./dev k8s install --masters x.x.x.x,x.x.x.x --nodes x.x.x.x,x.x.x.x -p {password of all machines}

# Ip range
./dev k8s install --masters x.x.x.x-x.x.x.x --nodes x.x.x.x-x.x.x.x -p {password of all machines}
```
