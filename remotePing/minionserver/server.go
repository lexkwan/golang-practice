package main

import (
	"bufio"
	"log"
	"net"
	"os/exec"

	minionpb "github.com/lexkwan/golang-practice/remotePing/minionpb"

	"google.golang.org/grpc"
)

type server struct{}

func (*server) Ping(req *minionpb.PingRequest, ps minionpb.PingService_PingServer) error {
	ip := req.GetIp()
	log.Printf("Receiving ping request for %v", ip)
	cmd := exec.Command("ping", ip)
	stdout, _ := cmd.StdoutPipe()
	cmd.Start()

	scanner := bufio.NewScanner(stdout)

	for scanner.Scan() {
		m := scanner.Text()
		log.Println(m)
		res := &minionpb.PingResponse{
			Result: m,
		}
		ps.Send(res)
	}
	cmd.Wait()
	return nil
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:10086")
	if err != nil {
		log.Fatalf("failed to listen:%sv", err)
	}
	s := grpc.NewServer()

	minionpb.RegisterPingServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to server: %v", err)
	}
}
