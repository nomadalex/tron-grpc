package main

import (
	"context"
	"encoding/json"
	"github.com/fullstackwang/tron-grpc"
	"github.com/fullstackwang/tron-grpc/api"
	"google.golang.org/grpc"
	"log"
)

func logData(v any) {
	s, err := json.Marshal(v)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(string(s))
}

func main() {
	client := tron_grpc.NewGrpcClient("", "")
	err := client.Start(grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}

	var in api.EmptyMessage
	block, err := client.GetNowBlock2(context.Background(), &in)
	if err != nil {
		log.Fatalln(err)
	}

	logData(block)
}
