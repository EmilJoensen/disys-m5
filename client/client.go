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
	conn, err := grpc.Dial(":8000", grpc.WithInsecure())
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
			log.Fatalf("Error when calling GetTime: %s", err)
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
