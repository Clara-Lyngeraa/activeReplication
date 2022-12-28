package main

import (
	proto "activeReplication/proto"
	"flag"
)

type Server struct {
	proto.UnimplementedIncrementServiceServer
	name string
	port int32
	amount int32
}

var port = flag.Int("port", 0, "server port number")

