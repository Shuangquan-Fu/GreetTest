package main

import (
	"GreetTest/greet/greetpb"
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
)



type  server struct {
	greetpb.UnimplementedGreeterServer
}


func (s *server) SayHello(ctx context.Context, in *greetpb.HelloRequest) (*greetpb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &greetpb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func main()  {
	lis, err :=net.Listen("tcp","localhost:50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	greetpb.RegisterGreeterServer(s, &server{})
	if err:= s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}