package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/EmilJoensen/disys-m5/auction"
	"google.golang.org/grpc"
)

func main() {
	arg1, _ := strconv.ParseInt(os.Args[1], 10, 32)
	ownPort := int32(arg1) + 8000

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	au := &AuctionServer{
		id:         ownPort,
		auction:    true,
		highestBid: 0,
		bidders:    make(map[int32]int32),
		ctx:        ctx,
	}

	list, err := net.Listen("tcp", fmt.Sprintf(":%v", ownPort))
	if err != nil {
		log.Fatalf("Failed to listen on port: %v", err)
	}
	grpcServer := grpc.NewServer()
	auction.RegisterAuctionServer(grpcServer, au)

	starttime := int(time.Now().Unix())
	runtime := rand.Intn(100)

	log.Printf("Auction has started and will run for %v seconds", runtime)
	// Close auction after given time
	go func() {
		for {
			if runtime+starttime == int(time.Now().Unix()) {
				au.mu.Lock()
				defer au.mu.Unlock()
				au.auction = false
				log.Printf("Auction is closed!")
				log.Printf("Winner is...")
				break
			}
		}
	}()

	if err := grpcServer.Serve(list); err != nil {
		log.Fatalf("failed to server %v", err)
	}
}

type AuctionServer struct {
	auction.UnimplementedAuctionServer
	id         int32
	auction    bool            // Indicate whether auction is on or not
	highestBid int32           // Stores highest overall bid
	bidders    map[int32]int32 // Stores highest bid for each bidder
	ctx        context.Context
	mu         sync.Mutex
}

func (au *AuctionServer) Bid(ctx context.Context, in *auction.BidAmount) (*auction.BidAck, error) {

	if _, ok := au.bidders[in.Id]; !ok {
		au.bidders[in.Id] = 0
		log.Printf("Received bid from: %v", in.Id)
	}

	au.mu.Lock()
	defer au.mu.Unlock()

	rep := &auction.BidAck{}
	au.bidders[in.Id] = in.Amount

	if in.Amount < au.highestBid || !au.auction {
		rep = &auction.BidAck{Ack: "Failed"}
		log.Printf("Bid %v from %v too low", in.Amount, in.Id)
	} else if in.Amount > au.highestBid {
		au.highestBid = in.Amount
		rep = &auction.BidAck{Ack: "Success"}
		log.Printf("Highest bid is now %v from %v", au.highestBid, in.Id)
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

	return rep, nil
}
