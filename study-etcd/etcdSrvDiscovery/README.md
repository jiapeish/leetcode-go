git@github.com:jiapeish/etcdSrvDiscovery.git

## Build
#### Prepare Go Environment
1. Install Golang: sudo yum install golang
2. `mkdir -p /opt/go-work/`
3. add environment variable to ~/.profile
- `export GOROOT=/usr/local/go`
- `export GOPATH=/opt/go-work/`
- `export PATH=$GOPATH/bin:$GOROOT/bin:$PATH`

#### Install govendor and Get vendor packages
1. put the project inside `/opt/go-work/src/github.com/jiapeish/`
2. `go get -u github.com/kardianos/govendor`
3. `cd jiapeish/etcdSrvDiscovery/`
4. `govendor init`
5. `govendor list` or `govendor list -v fmt`
6. `govendor fetch +out` or `govendor add +e`
7. `cd etcdSrvDiscovery/exmaple`
8. `go build`

## Run
1. run etcd server on the localhost.
2. `./example -role master`
3. `./example -role worker`