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
	"fmt"
	"log"
	"math/rand"
	"net"
	"sync"
	"time"
	"github.com/EmilJoensen/disys-m5/auction"
	"google.golang.org/grpc"
)

type Spinner int

var spinBytes = []byte{'-', '/', '|', '\\'}

func (s *Spinner) Tick() {
	*s = (*s + 1) % 4
	fmt.Printf("\b%s", spinBytes[*s:*s+1])
}

func main() {
    rand.Seed(time.Now().Unix())
    handleAuction(0, time.Now().Unix())
}

func handleAuction(highestBid int32, starttime int64) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	au := &AuctionServer{
		id:         8000,
		auction:    true,
		highestBid: highestBid,
        startTime: starttime,
		bidders:    make(map[int32]int32),
		ctx:        ctx,
	}

	list, err := net.Listen("tcp", fmt.Sprintf(":%v", 8000))
	if err != nil {
		log.Printf("Failed to listen on port: %v", err)
        log.Printf("It seems the allready is a server running on port 8000")
        log.Printf("I will go to standby mode")
        standby()
        return
	}
    log.Println("Primary server")
    log.Println("If this server is killed, a secondary server will take over")
	grpcServer := grpc.NewServer()
	auction.RegisterAuctionServer(grpcServer, au)

    var runtime int64 = 45

    elapsed := time.Since(time.Unix(starttime,0))

	log.Printf("Auction has started and will run for %v seconds", runtime - int64(elapsed / time.Second))
	// Close auction after given time
	go func() {
		for {
			if runtime+starttime < time.Now().Unix() {
				au.mu.Lock()
				defer au.mu.Unlock()
				au.auction = false
				log.Printf("Auction is closed!")
				log.Printf("Winner of auction had bid of %v", au.highestBid)
				break
			}
		}
	}()

	if err := grpcServer.Serve(list); err != nil {
		log.Fatalf("failed to server %v", err)
	}
}

func standby (){
    log.Println("Secondary backup server")
    log.Println("If the primary server is killed, this server will take over")
    log.Println("Waiting until primary server is killed...")
    var s Spinner
    var conn *grpc.ClientConn
    var opts []grpc.DialOption
    opts = append(opts, grpc.WithBlock(), grpc.WithInsecure())
	conn, err := grpc.Dial("localhost:8000", opts...)
	if err != nil {
		log.Fatalf("Could not connect: %s", err)
	}
	defer conn.Close()
    var starttime int64 = time.Now().Unix()
	var highestBid int32
	c := auction.NewAuctionClient(conn)
	for {

		req := auction.ResultVoid{}

		response, err := c.Result(context.Background(), &req)
		if err != nil {
            fmt.Print("\b")
            log.Printf("It seems the primary server died")
            log.Printf("Taking over...")
            time.Sleep(2 * time.Second)
            handleAuction(highestBid, starttime)
            return
		}

		s.Tick()

		if response.Status == "Auction finished" {
			break
		} else if response.Status == "Auction running" {
			highestBid = response.Outcome
		}
        
        starttime = response.Starttime

		time.Sleep(time.Duration(100) * time.Millisecond)
	}
}   

type AuctionServer struct {
	auction.UnimplementedAuctionServer
	id         int32
	auction    bool            // Indicate whether auction is on or not
	highestBid int32           // Stores highest overall bid
    startTime  int64           // Starttime in int64
	bidders    map[int32]int32 // Stores highest bid for each bidder
	ctx        context.Context
	mu         sync.Mutex
}

func (au *AuctionServer) Bid(ctx context.Context, in *auction.BidAmount) (*auction.BidAck, error) {

	if _, ok := au.bidders[in.Id]; !ok {
		au.bidders[in.Id] = 0
	}

	au.mu.Lock()
	defer au.mu.Unlock()

	rep := &auction.BidAck{}
	au.bidders[in.Id] = in.Amount

	if in.Amount < au.highestBid || !au.auction {
		rep = &auction.BidAck{Ack: "Failed"}
		log.Printf("Bid %v too low", in.Amount)
	} else if in.Amount > au.highestBid {
		au.highestBid = in.Amount
		rep = &auction.BidAck{Ack: "Success"}
		log.Printf("Highest bid is now %v", au.highestBid)
	} else {
		rep = &auction.BidAck{Ack: "Error"}
	}

	log.Printf("Highest bid is %v", au.highestBid)

	return rep, nil
}

func (au *AuctionServer) Result(ctx context.Context, in *auction.ResultVoid) (*auction.ResultOutcome, error) {
	au.mu.Lock()
	defer au.mu.Unlock()

	rep := &auction.ResultOutcome{Outcome: au.highestBid}
	if au.auction {
		rep.Status = "Auction running"
	} else {
		rep.Status = "Auction finished"
	}
    rep.Starttime = au.startTime

	return rep, nil
}
