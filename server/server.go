package main

import (
	proto "activeReplication/proto"
	"flag"
	"context"
	grpc "google.golang.org/grpc"
)

type Server struct {
	proto.UnimplementedIncrementServiceServer
	name string
	port int32
	amount int32
}

var port = flag.Int("port", 0, "server port number")

func startServer(server Server){
	//create a grpc server
	grpcServer := grpc.NewServer()

	// Make the server listen at the given port (convert int port to string)
	listener, err := net.Listen("tcp", ":"+strconv.Itoa(server.port))

	if err != nil {
		log.Fatalf("Could not create the server %v", err)
	}
	log.Printf("Started server at port: %d\n", server.port)

	// Register the grpc server and serve its listener
	proto.RegisterIncrementServiceServer(grpcServer, server)
	serveError := grpcServer.Serve(listener)
	if serveError != nil {
		log.Fatalf("Could not serve listener")
	}

}

func (c *Server) Increment(ctx context.Context, in *proto.IncRequest) (*proto.IncResponse, error) {
	return nil, nil;
}