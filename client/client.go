package main 

import (
	"context"
	// "io"
	"poc_team1/grpcgen"
	"poc_team1/pb"
	"google.golang.org/grpc"
	"log"
)

func main() {
	conn,_ := grpc.Dial("localhost:50005",grpc.WithInsecure())
	client := grpcgen.NewStreamitClient(conn)
	in := &pb.Request{Message:"hi"}
	stream,_:= client.DataStreamer(context.Background(),in)
	for i:=0;i<10;i++{
		resp,_:= stream.Recv()
		log.Printf(resp.Result)
	}
}