package main 

import (
	"fmt"
	"log"
	"net"
	"google.golang.org/grpc"
	"poc_team1/pb"
	"poc_team1/grpcgen"
)


type server struct{}

func (s server) DataStreamer(in *pb.Request, srv grpcgen.Streamit_DataStreamerServer) error{
	log.Printf("DataStreamer is invoked")
	for i:=0;i<10;i++ {
		resp := pb.Response{Result:fmt.Sprintf("Response id : %v",i)}
		srv.Send(&resp)
	}
	return nil
}

func main() {
	lis,_ := net.Listen("tcp",":50005")
	s:=grpc.NewServer()
	grpcgen.RegisterStreamitServer(s,server{})
	s.Serve(lis)
}