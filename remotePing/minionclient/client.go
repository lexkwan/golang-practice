package main

import (
	"context"
	"flag"
	"log"

	minionpb "github/lexkwan/golang-practice/remotePing/minionpb"

	"google.golang.org/grpc"
)

func main() {

	serip := flag.String("server", "localhost:10086", "Minion server ip")
	target := flag.String("target", "", "Ping Target ip")
	flag.Parse()
	cc, err := grpc.Dial(*serip, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error while trying to connect to the server: %v", err)
	}
	defer cc.Close()

	c := minionpb.NewPingServiceClient(cc)

	req := &minionpb.PingRequest{
		Ip: *target,
	}

	resStream, err := c.Ping(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling ping rpc %v ", err)

	}
	for {
		msg, err := resStream.Recv()
		if err != nil {
			log.Printf("%v", err)
			return
		}
		log.Println(msg.GetResult())
	}

}
