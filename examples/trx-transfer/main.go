package main

import (
	"context"
	"encoding/hex"
	"github.com/fullstackwang/tron-grpc/client"
	"github.com/fullstackwang/tron-grpc/trx"
	"google.golang.org/grpc"
	"log"
)

func main() {
	client := client.New("grpc.shasta.trongrid.io:50051", "")
	client.SetPrivateKey("8e812436a0e3323166e1f0e8ba79e19e217b2c4a53c970d4cca0cfb1078979df")
	log.Println(client.Signer.Address().String())

	client.Start(grpc.WithInsecure())

	trxClient := trx.New(client)
	balance, err := trxClient.GetBalance(context.Background(), client.Signer.Address().String())
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(balance)

	tx, err := trxClient.Transfer(context.Background(), "TSdRuWzGTce2f9fzNWw122fmBMfDdBoxcx", 100000)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(hex.EncodeToString(tx.Txid))
}
