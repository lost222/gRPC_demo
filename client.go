package main

import (
	"context"
	"fmt"
	pb "./pb"
	"google.golang.org/grpc"
)

func main() {
	serviceAddress := "127.0.0.1:1234"
	conn,err := grpc.Dial(serviceAddress, grpc.WithInsecure())
	if err != nil {
		panic("connect error")
	}
	// @todo close的err没有处理
	defer conn.Close()
	stringClient := pb.NewStringServiceClient(conn)
	stringReq := &pb.StringRequest{A:"A", B:"B"}
	reply, _ := stringClient.Concat(context.Background(), stringReq)
	fmt.Printf("StringService Concat : %s concat %s = %s", stringReq.A, stringReq.B, reply.Ret)
}
