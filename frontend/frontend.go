package main

import (
	proto "activeReplication/proto"
)

type Frontend struct {
	proto.UnimplementedIncrementServiceServer
	name            string
	port            int32
	servers         []proto.incrementServiceClient
}

var port = flag.Int("port", 0, "server port number") // create the port that recieves the port that the client wants to access to

func startFrontend(frontend *Frontend) {
	// Create a new grpc server
	grpcServer := grpc.NewServer()

	// Make the server listen at the given port (convert int port to string)
	listener, err := net.Listen("tcp", "localhost:"+strconv.Itoa(frontend.port))

	if err != nil {
		log.Fatalf("Could not create the frontend %v", err)
	}
	log.Printf("Started frontend at port: %d\n", frontend.port)

	// Register the grpc server and serve its listener
	proto.RegisterAuctionServer(grpcServer, frontend)

	serveError := grpcServer.Serve(listener)
	fmt.Printf("nedern")
	if serveError != nil {
		log.Fatalf("Could not serve listener frontend")

	}
}

func (c *Server) Increment(ctx context.Context, in *proto.IncRequest) (*proto.IncResponse, error) {
	return nil, nil;
}