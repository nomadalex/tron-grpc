package main

import (
	"context"
	"encoding/hex"
	"github.com/fullstackwang/tron-grpc/address"
	"github.com/fullstackwang/tron-grpc/client"
	"github.com/fullstackwang/tron-grpc/contract"
	"github.com/fullstackwang/tron-grpc/trx"
	"github.com/fullstackwang/tron-grpc/tx"
	"github.com/fullstackwang/tron-grpc/wallet"
	"google.golang.org/grpc"
	"log"
	"math/big"
)

const erc20Abi = "[ { \"inputs\": [ { \"internalType\": \"string\", \"name\": \"name_\", \"type\": \"string\" }, { \"internalType\": \"string\", \"name\": \"symbol_\", \"type\": \"string\" } ], \"stateMutability\": \"nonpayable\", \"type\": \"constructor\" }, { \"anonymous\": false, \"inputs\": [ { \"indexed\": true, \"internalType\": \"address\", \"name\": \"owner\", \"type\": \"address\" }, { \"indexed\": true, \"internalType\": \"address\", \"name\": \"spender\", \"type\": \"address\" }, { \"indexed\": false, \"internalType\": \"uint256\", \"name\": \"value\", \"type\": \"uint256\" } ], \"name\": \"Approval\", \"type\": \"event\" }, { \"anonymous\": false, \"inputs\": [ { \"indexed\": true, \"internalType\": \"address\", \"name\": \"from\", \"type\": \"address\" }, { \"indexed\": true, \"internalType\": \"address\", \"name\": \"to\", \"type\": \"address\" }, { \"indexed\": false, \"internalType\": \"uint256\", \"name\": \"value\", \"type\": \"uint256\" } ], \"name\": \"Transfer\", \"type\": \"event\" }, { \"inputs\": [ { \"internalType\": \"address\", \"name\": \"owner\", \"type\": \"address\" }, { \"internalType\": \"address\", \"name\": \"spender\", \"type\": \"address\" } ], \"name\": \"allowance\", \"outputs\": [ { \"internalType\": \"uint256\", \"name\": \"\", \"type\": \"uint256\" } ], \"stateMutability\": \"view\", \"type\": \"function\" }, { \"inputs\": [ { \"internalType\": \"address\", \"name\": \"spender\", \"type\": \"address\" }, { \"internalType\": \"uint256\", \"name\": \"amount\", \"type\": \"uint256\" } ], \"name\": \"approve\", \"outputs\": [ { \"internalType\": \"bool\", \"name\": \"\", \"type\": \"bool\" } ], \"stateMutability\": \"nonpayable\", \"type\": \"function\" }, { \"inputs\": [ { \"internalType\": \"address\", \"name\": \"account\", \"type\": \"address\" } ], \"name\": \"balanceOf\", \"outputs\": [ { \"internalType\": \"uint256\", \"name\": \"\", \"type\": \"uint256\" } ], \"stateMutability\": \"view\", \"type\": \"function\" }, { \"inputs\": [], \"name\": \"decimals\", \"outputs\": [ { \"internalType\": \"uint8\", \"name\": \"\", \"type\": \"uint8\" } ], \"stateMutability\": \"view\", \"type\": \"function\" }, { \"inputs\": [ { \"internalType\": \"address\", \"name\": \"spender\", \"type\": \"address\" }, { \"internalType\": \"uint256\", \"name\": \"subtractedValue\", \"type\": \"uint256\" } ], \"name\": \"decreaseAllowance\", \"outputs\": [ { \"internalType\": \"bool\", \"name\": \"\", \"type\": \"bool\" } ], \"stateMutability\": \"nonpayable\", \"type\": \"function\" }, { \"inputs\": [ { \"internalType\": \"address\", \"name\": \"spender\", \"type\": \"address\" }, { \"internalType\": \"uint256\", \"name\": \"addedValue\", \"type\": \"uint256\" } ], \"name\": \"increaseAllowance\", \"outputs\": [ { \"internalType\": \"bool\", \"name\": \"\", \"type\": \"bool\" } ], \"stateMutability\": \"nonpayable\", \"type\": \"function\" }, { \"inputs\": [], \"name\": \"name\", \"outputs\": [ { \"internalType\": \"string\", \"name\": \"\", \"type\": \"string\" } ], \"stateMutability\": \"view\", \"type\": \"function\" }, { \"inputs\": [], \"name\": \"symbol\", \"outputs\": [ { \"internalType\": \"string\", \"name\": \"\", \"type\": \"string\" } ], \"stateMutability\": \"view\", \"type\": \"function\" }, { \"inputs\": [], \"name\": \"totalSupply\", \"outputs\": [ { \"internalType\": \"uint256\", \"name\": \"\", \"type\": \"uint256\" } ], \"stateMutability\": \"view\", \"type\": \"function\" }, { \"inputs\": [ { \"internalType\": \"address\", \"name\": \"recipient\", \"type\": \"address\" }, { \"internalType\": \"uint256\", \"name\": \"amount\", \"type\": \"uint256\" } ], \"name\": \"transfer\", \"outputs\": [ { \"internalType\": \"bool\", \"name\": \"\", \"type\": \"bool\" } ], \"stateMutability\": \"nonpayable\", \"type\": \"function\" }, { \"inputs\": [ { \"internalType\": \"address\", \"name\": \"sender\", \"type\": \"address\" }, { \"internalType\": \"address\", \"name\": \"recipient\", \"type\": \"address\" }, { \"internalType\": \"uint256\", \"name\": \"amount\", \"type\": \"uint256\" } ], \"name\": \"transferFrom\", \"outputs\": [ { \"internalType\": \"bool\", \"name\": \"\", \"type\": \"bool\" } ], \"stateMutability\": \"nonpayable\", \"type\": \"function\" } ]"

func main() {
	client := client.New("grpc.shasta.trongrid.io:50051", "")
	client.SetPrivateKey("8e812436a0e3323166e1f0e8ba79e19e217b2c4a53c970d4cca0cfb1078979df")
	log.Println(client.Signer.Address().String())

	client.Start(grpc.WithInsecure())

	c := createContract(client)
	checkConstantCall(client, c)
	checkSend(client, c)
	checkEvents(client, c)

	checkErc20(client)
}

func createContract(client *client.Client) *contract.Contract {
	contractAddr, _ := address.FromBase58("TMLTMoPQFXyof68S2C6X8Nsjso3cqajUMz")
	log.Println(contractAddr.Hex(), contractAddr.String())
	c := contract.New(client, contractAddr)
	_ = c.LoadABI([]byte(erc20Abi))
	return c
}

func checkConstantCall(client *client.Client, c *contract.Contract) {
	ret, err := c.Call(context.Background(), "balanceOf", client.Signer.Address())
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(ret)
}

func checkSend(client *client.Client, c *contract.Contract) {
	toAddr, _ := address.FromBase58("TSdRuWzGTce2f9fzNWw122fmBMfDdBoxcx")
	tx, err := c.Send(context.Background(), "transfer", toAddr, 1000000)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(hex.EncodeToString(tx.Txid))
	err = tx.WaitConfirmation()
	if err != nil {
		log.Fatalln(err)
	}

	ret, err := tx.GetResult()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(ret)
}

func checkEvents(client *client.Client, c *contract.Contract) {
	tt, err := tx.GetFromID(context.Background(), client, "dce46f40a26ae385b3d9770a92666f050d4a612dd1e5ba73aeb6158de1549342", true)
	if err != nil {
		log.Fatalln(err)
	}

	events, err := c.GetEvents(tt)
	if err != nil {
		log.Fatalln(err)
	}

	for i, event := range events {
		log.Println(i, "event", event)
		for i, input := range event.Inputs {
			log.Println(i, "input", input)
		}
	}
}

func checkErc20(client *client.Client) {
	c := contract.NewErc20(client, address.FromBase58Unsafe("TMLTMoPQFXyof68S2C6X8Nsjso3cqajUMz"))

	name, err := c.Name(context.Background())
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("name", name)
	symbol, err := c.Symbol(context.Background())
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("symbol", symbol)
	decimals, err := c.Decimals(context.Background())
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("decimals", decimals)
	totalSupply, err := c.TotalSupply(context.Background())
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("totalSupply", totalSupply)
	balance, err := c.BalanceOf(context.Background(), client.Signer.Address())
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("balance", balance)

	signer, _ := wallet.FromPrivateKey("8e812436a0e3323166e1f0e8ba79e19d217b2c4a53c970d4cca0cfb1078979df")
	log.Println("new signer", signer.Address().String())

	trxClient := trx.New(client)
	tx, err := trxClient.Transfer(context.Background(), signer.Address().String(), 20000000)
	if err != nil {
		log.Fatalln(err)
	}
	tx.WaitConfirmation()

	tx, err = c.Approve(context.Background(), signer.Address(), big.NewInt(100000))
	if err != nil {
		log.Fatalln(err)
	}
	tx.WaitConfirmation()

	allowance, err := c.Allowance(context.Background(), client.Signer.Address(), signer.Address())
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("allowance", allowance)

	oldSigner := client.Signer
	client.Signer = signer

	tx, err = c.TransferFrom(context.Background(), oldSigner.Address(), signer.Address(), big.NewInt(100000))
	if err != nil {
		log.Fatalln(err)
	}
	tx.WaitConfirmation()

	tx, err = c.Transfer(context.Background(), oldSigner.Address(), big.NewInt(100000))
	if err != nil {
		log.Fatalln(err)
	}
	tx.WaitConfirmation()
}
