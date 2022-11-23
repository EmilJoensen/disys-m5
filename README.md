# disys-m5
### By Anders & Emil
## Install
No installation required other than golang

## Running the code
The code has two programs. One for the clients and one for the servers.

### Client
To start the client from the root of the folder:

```
go run client/client.go <ID>
```

For example start two clients: 

```
go run client/client.go 1
go run client/client.go 2
```

### Server
To start the servers from the root of the folder: 

```
go run server/server.go
```

You can start multiple servers. If one server already is running, then new servers will start in a standby mode.
Standby servers are ready to take over if the primary server crashes. Multiple standby servers can be ready to take over at once.
Standby servers listen to the primary server to know both the remaining time and the current highest bid. 

To start one primary server with 2 replicas.

```
go run server/server.go
go run server/server.go
go run server/server.go
```

## Output of running system - The logs
explain...

# Report
## Introduction 
A short introduction to what you have done.

##
Architecture
A description of the architecture of the system and the protocol (behaviour), including any protocols used internally between nodes of the system.

## Correctness 1
Argue whether your implementation satisfies linearisability or sequential consistency (after precisely defining  the property it satisfies).

## Correctness 2
An argument that your protocol is correct in the absence and the presence of failures.
