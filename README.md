# disys-m5
### By Anders & Emil
## Install
No installation required other than golang

## Running the code
The code has two programs. One for the clients and one for the servers.

### client.go
To start the client from the root of the folder:

```
go run client/client.go <ID>
```

For example start two clients: 

```
go run client/client.go 1
go run client/client.go 2
```

### server.go
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


```
┌─────────────────────────────┐
│                             │
│  client                     │
│  ------                     │
│                             │
│  Bid:                       │
│    Bid on auction           │
│                             │
│  Result:                    │
│    Check status of auction  │
│                             │
└─────────────────────────────┘

┌────────────────────────────────────────────────┐
│                                                │
│  server                                        │
│  ------                                        │
│                                                │
│  handleAuction:                                │
│    Check if auction is running if not start    │
│    auction and start taking bids. Else standby │
│                                                │
│  standby:                                      │
│    If auction primary server is running read   │
│    result and keep track of time, and highest  │
│    bid. When the primary server stops for any  │
│    reason try to become primary server by      │
│    calling handleAuction.                      │
│                                                │
│  Bid:                                          │
│    Check if bid is higher than highestBid. If  │
│    bid is too low return Ack Failed. If bid    │
│    is higher than highestBid update it and send│
│    Ack Succes back.                            │
│                                                │
│  Result:                                       │
│    Send back the status of the auction.        │
│    Status of the current highest bid and       │
│    wheter the auction is running or finished.  │
│                                                │
│                                                │
└────────────────────────────────────────────────┘

┌────────────────┐
│                │                                        ┌────────────────────┐
│ Client 1       │ Result                                 │                    │
│ --------       ├────────►┌────────────────────┐  Result │  Standby server    │
│                │ Status  │                    │ ◄───────┤  --------------    │
└────────────────┘◄────────┤  Primary server    │         │                    │
                           │  --------------    │ Status  │  Port: ?           │
┌────────────────┐ Bid     │                    ├───────► │                    │
│                ├────────►│  Port: 8000        │         └────────────────────┘
│ Client 2       │ Ack F/S │                    │   Result
│ --------       │◄─────── └────────────┬───────┘◄────────────┐
│                │                      │                     │
└────────────────┘                      │                 ┌───┴────────────────┐
                                        │                 │                    │
                                        │                 │  Standby server    │
                                        │    Status       │  --------------    │
                                        └───────────────► │                    │
                                                          │  Port: ?           │
                                                          │                    │
                                                          └────────────────────┘
```

## Correctness 1
Argue whether your implementation satisfies linearisability or sequential consistency (after precisely defining  the property it satisfies).

## Correctness 2
An argument that your protocol is correct in the absence and the presence of failures.
