# Concurrent TCP Server with TCP Client

Server to check whether a number is prime

## Usage

First run server in a new terminal

```
cd ./server/
go run server.go -port=9091
```

Then open new terminal for each client

```
cd ./client/
go run client.go -port=9091
```

If port is not specified, it defaults to 8080
