# Goob
Goob(Go + noob) is a simple and straightforward project for noobs (newbies) in Go which covers some important aspects of a standard Go project like: [standard layout](https://github.com/golang-standards/project-layout), [configuration solution](https://github.com/spf13/viper), [testing](https://github.com/stretchr/testify) and more.


The project is a [gRPC-Protobuf](https://grpc.io/docs/what-is-grpc/introduction/) Client-Server [Command-line](https://github.com/spf13/cobra) application with the context of Blockchain, the client requests a [nonce](https://www.tutorialspoint.com/what-is-a-nonce-in-block-chain) for a sample record and the server, after performing an intensive task will responds it.
Note that the purpose of this project is not about the Blockchain topics, just to perform an intensive task for the sake of practicing [concurrent patterns](https://github.com/lotusirous/go-concurrency-patterns) in Go.


# How to use

Run server
```
make run-server
```


Run client
```
make run-client
```

See the Makefile for the other commands
