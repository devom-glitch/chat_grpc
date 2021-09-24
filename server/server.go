package main 

import (
	"math/rand"
	"fmt"
	"log"
	"net"
	"google.golang.org/grpc"
	"poc_team1/pb"
	"poc_team1/grpcgen"
	"time"
)


type server struct{}

func (s server) DataStreamer(in *pb.Request, srv grpcgen.Streamit_DataStreamerServer) error{
	log.Printf("DataStreamer is invoked by %v",in.Message)
	i:=0
	for _ = range time.Tick(1*time.Second) {
		resp := pb.Response{Result:fmt.Sprintf("Response id : %v",rand.Intn(1000))}
		if (i<20){
			srv.Send(&resp)
		} else{
			srv.Send(nil)
			break
		}
		i++
	}
	log.Printf("Streaming off by server for %v",in.Message)
	return nil
}

func main() {
	lis,_ := net.Listen("tcp",":50005")
	s:=grpc.NewServer()
	grpcgen.RegisterStreamitServer(s,server{})
	s.Serve(lis)
}