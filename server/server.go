package main

import (
	"context"

	"github.com/EmilJoensen/disys-m5/auction"
)

type AuctionServer struct {
	auction.UnimplementedAuctionServer
	id      int32
	bidders map[int32]auction.AuctionClient
	ctx     context.Context
}

func (au *AuctionServer) Bid(ctx context.Context, in *auction.BidAmount) (*auction.BidAck, error) {
	c := auction.NewAuctionClient()
	au.bidders[in.Id] = c
}
