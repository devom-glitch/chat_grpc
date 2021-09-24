package main 

import (
	"github.com/satori/go.uuid"
	"context"
	// "io"
	"poc_team1/grpcgen"
	"poc_team1/pb"
	"google.golang.org/grpc"
	"log"
)

func main() {
	myuuid := uuid.NewV4()
	conn,_ := grpc.Dial(":8080",grpc.WithInsecure())
	client := grpcgen.NewStreamitClient(conn)
	in := &pb.Request{Message:myuuid.String()}
	stream,_:= client.DataStreamer(context.Background(),in)
	for {
		resp,err:= stream.Recv()
		if err != nil {
			// log.Printf("%v",err)
			break
		}
		log.Printf(resp.Result)
	}
}