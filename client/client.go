package main

import (
	"context"
	"fmt"
	"log"

	pb "learn-go/gen/proto"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:8089", grpc.WithInsecure())
	if err != nil {
		log.Println(err)
	}
	client := pb.NewTestApiClient(conn)
	resp, err := client.Echo(context.Background(), &pb.ResponseRequest{Msg: "HELOO EVERYONE"})
	if err != nil {
		log.Println(err)
	}
	fmt.Println(resp)
	fmt.Println(resp.Msg)
}
