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

Output of server 1 until crash

```
$ go run server/server.go 
2022/11/28 10:33:48 Primary server
2022/11/28 10:33:48 If this server is killed, a secondary server will take over
2022/11/28 10:33:48 Auction has started and will run for 45 seconds
2022/11/28 10:33:50 Highest bid is now 1
2022/11/28 10:33:50 Highest bid is 1
2022/11/28 10:33:54 Highest bid is now 2
2022/11/28 10:33:54 Highest bid is 2
2022/11/28 10:33:57 Highest bid is now 9
2022/11/28 10:33:57 Highest bid is 9
signal: interrupt (Crash)
```

Output of server 2

```
$go run server/server.go 
2022/11/28 10:33:49 Failed to listen on port: listen tcp :8000: bind: address already in use
2022/11/28 10:33:49 It seems the allready is a server running on port 8000
2022/11/28 10:33:49 I will go to standby mode
2022/11/28 10:33:49 Secondary backup server
2022/11/28 10:33:49 If the primary server is killed, this server will take over
2022/11/28 10:33:49 Waiting until primary server is killed...
2022/11/28 10:33:59 It seems the primary server died
2022/11/28 10:33:59 Taking over...
2022/11/28 10:34:01 Primary server
2022/11/28 10:34:01 If this server is killed, a secondary server will take over
2022/11/28 10:34:01 Auction has started and will run for 32 seconds
2022/11/28 10:34:01 Highest bid is now 16
2022/11/28 10:34:01 Highest bid is 16
2022/11/28 10:34:07 Highest bid is now 17
2022/11/28 10:34:07 Highest bid is 17
2022/11/28 10:34:10 Highest bid is now 18
2022/11/28 10:34:10 Highest bid is 18
2022/11/28 10:34:15 Highest bid is now 23
2022/11/28 10:34:15 Highest bid is 23
2022/11/28 10:34:15 Highest bid is now 29
2022/11/28 10:34:15 Highest bid is 29
2022/11/28 10:34:15 Highest bid is now 33
2022/11/28 10:34:15 Highest bid is 33
2022/11/28 10:34:16 Highest bid is now 35
2022/11/28 10:34:16 Highest bid is 35
signal: interrupt (Crash)
```

Output of server 1 after restart

```
$ go run server/server.go 
2022/11/28 10:34:04 Failed to listen on port: listen tcp :8000: bind: address already in use
2022/11/28 10:34:04 It seems the allready is a server running on port 8000
2022/11/28 10:34:04 I will go to standby mode
2022/11/28 10:34:04 Secondary backup server
2022/11/28 10:34:04 If the primary server is killed, this server will take over
2022/11/28 10:34:04 Waiting until primary server is killed...
2022/11/28 10:34:18 It seems the primary server died
2022/11/28 10:34:18 Taking over...
2022/11/28 10:34:20 Primary server
2022/11/28 10:34:20 If this server is killed, a secondary server will take over
2022/11/28 10:34:20 Auction has started and will run for 13 seconds
2022/11/28 10:34:21 Highest bid is now 40
2022/11/28 10:34:21 Highest bid is 40
2022/11/28 10:34:21 Highest bid is now 46
2022/11/28 10:34:21 Highest bid is 46
2022/11/28 10:34:21 Highest bid is now 50
2022/11/28 10:34:21 Highest bid is 50
2022/11/28 10:34:22 Highest bid is now 52
2022/11/28 10:34:22 Highest bid is 52
2022/11/28 10:34:25 Highest bid is now 60
2022/11/28 10:34:25 Highest bid is 60
2022/11/28 10:34:29 Highest bid is now 61
2022/11/28 10:34:29 Highest bid is 61
2022/11/28 10:34:31 Highest bid is now 69
2022/11/28 10:34:31 Highest bid is 69
2022/11/28 10:34:34 Auction is closed!
2022/11/28 10:34:34 Winner of auction had bid of 69
```

Output of client 1

```
$go run client/client.go 
2022/11/28 10:33:54 Highest bid is currently 1
2022/11/28 10:33:54 Bidding 2
2022/11/28 10:34:01 Highest bid is currently 9
2022/11/28 10:34:01 Bidding 16
2022/11/28 10:34:10 Highest bid is currently 17
2022/11/28 10:34:10 Bidding 18
2022/11/28 10:34:19 Error when trying to connect: rpc error: code = Unavailable desc = connection error: desc = "transport: Error while dialing dial tcp [::1]:8000: connect: connection refused"
2022/11/28 10:34:19 Will retry in 1 second
2022/11/28 10:34:20 Error when trying to connect: rpc error: code = Unavailable desc = connection error: desc = "transport: Error while dialing dial tcp [::1]:8000: connect: connection refused"
2022/11/28 10:34:20 Will retry in 1 second
2022/11/28 10:34:21 Highest bid is currently 35
2022/11/28 10:34:21 Bidding 40
2022/11/28 10:34:21 Highest bid is currently 40
2022/11/28 10:34:21 Bidding 46
2022/11/28 10:34:21 Highest bid is currently 46
2022/11/28 10:34:21 Bidding 50
2022/11/28 10:34:22 Highest bid is currently 50
2022/11/28 10:34:22 Bidding 52
2022/11/28 10:34:31 Highest bid is currently 61
2022/11/28 10:34:31 Bidding 69
2022/11/28 10:34:35 Auction finished
2022/11/28 10:34:35 You won the auction with a bid of 69
```

Output of client 2

```
$go run client/client.go 
2022/11/28 10:33:50 Highest bid is currently 0
2022/11/28 10:33:50 Bidding 1
2022/11/28 10:33:57 Highest bid is currently 2
2022/11/28 10:33:57 Bidding 9
2022/11/28 10:34:07 Highest bid is currently 16
2022/11/28 10:34:07 Bidding 17
2022/11/28 10:34:15 Highest bid is currently 18
2022/11/28 10:34:15 Bidding 23
2022/11/28 10:34:15 Highest bid is currently 23
2022/11/28 10:34:15 Bidding 29
2022/11/28 10:34:15 Highest bid is currently 29
2022/11/28 10:34:15 Bidding 33
2022/11/28 10:34:16 Highest bid is currently 33
2022/11/28 10:34:16 Bidding 35
2022/11/28 10:34:25 Highest bid is currently 52
2022/11/28 10:34:25 Bidding 60
2022/11/28 10:34:29 Highest bid is currently 60
2022/11/28 10:34:29 Bidding 61
2022/11/28 10:34:34 Auction finished
2022/11/28 10:34:34 You lost the auction
2022/11/28 10:34:34 Highest bid is was 69
```
# Report
## Introduction 
In this assignment we are creating an auction system. We have three types of nodes. We have clients, primary server and standby servers.

Clients and servers communicate through a gRPC protocol. A client can ask for the status of the auction, and it can bid on the auction. A client uses this information to make decisions on what its next bid should be. If an auction is finished, it knows whether it won the auction or not.

When a client sends a bid request it waits for an ACK response. The ACK response is either success or a failure. If the bid was higher than the highest bit it will be a success. If it was not higher, then it will be a failure.

There is a delay server side for all bids of 100 milliseconds. In these 100 milliseconds the client does not know if the bid went through. After getting a succes the client will still have to ask for the Status of the auction to know if it still holds the highest bid.

The primary server handles the auction. When an auction is started it runs for 45 seconds. It starts with a highest bid of 0. It also stores the Unix timestamp where it started. This is bundled with every request for the status of the auction.

While the auction is running the server is open for connections on port 8000. The server will stay open for Status requests after it is done with the auction. So clients can see the status of the auction. 

Standby servers are multiple server that are ready to take over for the primary server. To be up to date with the primary server, a standby server sends Status request every 10 milliseconds. This means that a standby server will have the highest bid, the starttime of the auction updated more quickly than clients get ACK messages.

This can cause a small bug were a crash results in a client not getting an ACK message, but the bid goes through anyway for the standby server. This edgecase is not causing confusion for the client. As the client can just retry the bid. 

If we also saved the username of the bidding client we could store that along side the highest bid. This would mean that a new bid to the server could be verifyed serverside as a retry which would eliemnate the otherwise shadow bid that the client would have to bid higher. Though for a shadey auction house this is not really a bug and more like a "feature".

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
Argue whether your implementation satisfies linearisability or sequential consistency (after precisely defining the property it satisfies).

This exercise satisfies sequential consistency. For example, if all messages go through but a server crashes, the 100 millisecond delay of the ACK message will enable this. Linearizability does not hold for this exercise because incoming messages might be held captive serverside.

## Correctness 2
An argument that your protocol is correct in the absence and the presence of failures.

Passive replication allows for more than one server to crash and the system as a whole can still recover. Standby servers are getting the same information as the primary server through the status call. As long as the standby server has a similar local clock as the primary server it can take over and serve exactly the same information. If a backup server's clock has drifted, the auction time might vary.

Because we store all the information of the auction in the status and we do not use the past information. We can argue that a new standby server that comes online and receives one status request through to the primary server can seconds later stand in for a failing primary server without the client finding out. This holds as long as network delay is not longer than 90 milliseconds. As the backup server requests the primary server every 10 micro seconds.
