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
In this assignment we are creating an auction system. We have three types of nodes. We have clients, primary server and standby servers.

Clients and servers communicate through a gRPC protocol. A client can ask for the status of the auction, and it can bid on the auction. A client uses this information to make decisions on what its next bid should be. If an auction is finished, it knows whether it won the auction or not.

When a client sends a bid request it waits for an ACK response. The ACK response is either success or a failure. If the bid was higher than the highest bit it will be a success. If it was not higher, then it will be a failure.

There is a delay server side for all bids of 100 milliseconds. In these 100 milliseconds the client does not know if the bid went through. After getting a succes the client will still have to ask for the Status of the auction to know if it still holds the highest bid. 

The primary server handles the auction. When an auction is started it runs for 45 seconds. It starts with a highest bid of 0. It also stores the Unix timestamp where it started. This is bundled with every request for the status of the auction.

While the auction is running the server is open for connections on port 8000. The server will stay open for Status requests after it is done with the auction. So clients can see the status of the auction. 

Standby servers are multiple server that are ready to take over for the primary server. To be up to date with the primary server, a standby server sends Status request every 10 milliseconds. This means that a standby server will have the highest bid, the starttime of the auction updated more quickly than clients get ACK messages.

If the status request from the standby server results in an error. Say connection refused. Then the Standby server will try to start listening on port 8000 for connections and if this is successful it will become the primary server. if this port is still in use it is assumed that the primary server has recovered or another standby server got there first. The standby server will go back to standby mode if this is the case.

Because the standby server has both the starttime and the highest bid, it can continue to serve when it turns into the primary server. From the clients point of view they can be sure that their bids go through if they get the ack message. If a crash happens before the ack message, then their request might go through. But then trying the same bid again on the new primary server should be consistent for the client.

If you implemented this over the internet, then you would have to rely on a loadbalancer. In this assignment we assume that acquiring the port is the same as acquiring a lock with timeout in a distributed environment. We assume only crash tolerance and not byzantine fault tolerance.

We only need one alive primary server / standby server to complete the auction.
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
