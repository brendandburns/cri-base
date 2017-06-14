# cri-base
A simple base application for implementing Kubernetes Container Runtime Interface Servers

# Building
`go build cmd/criserver.go`

# Running
`criserver --alsologtostderr`

# Testing
```sh
go build cmd/client.go
./client
```
