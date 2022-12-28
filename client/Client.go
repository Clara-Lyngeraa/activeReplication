package main

import (
	proto "activeReplication/proto" 
	 "google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

)

type Client struct {
	id int32
	portNumber int32
}

var (
	clientPort   = flag.Int("cPort", 0, "client port number")
	frontendPort = flag.Int("fPort", 0, "frontend port number")
)