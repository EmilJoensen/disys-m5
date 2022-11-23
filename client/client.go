/*
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
 */
package main

import (
	"context"
	"log"
	"math/rand"
	"time"

	"github.com/EmilJoensen/disys-m5/auction"
	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn
    var opts []grpc.DialOption
    opts = append(opts, grpc.WithBlock(), grpc.WithInsecure())
	conn, err := grpc.Dial("localhost:8000", opts...)
	if err != nil {
		log.Fatalf("Could not connect: %s", err)
	}
	defer conn.Close()

	var highestBid int32
	c := auction.NewAuctionClient(conn)
	for {

		req := auction.ResultVoid{}

		response, err := c.Result(context.Background(), &req)
		if err != nil {
			log.Printf("Error when trying to connect: %s", err)
			log.Printf("Will retry in 1 second")
            time.Sleep(1 * time.Second)
            continue
		}

		log.Printf(response.Status)

		if response.Status == "Auction finished" {
			break
		} else if response.Status == "Auction running" {
			highestBid = response.Outcome
			log.Printf("Highest bid is currently %v", highestBid)
		}

		var bid int32 = int32(highestBid) + int32(rand.Intn(10))

		log.Printf("Bidding %v", bid)
		bidreq := auction.BidAmount{Amount: bid}
		c.Bid(context.Background(), &bidreq)

		time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
	}
}
