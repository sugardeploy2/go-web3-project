package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	HelloWorld "github.com/sugrdeploy2/go-web3-project/build"
)

func main() {
	client, err := ethclient.Dial("http://127.0.0.1:8545")
	if err != nil {
		fmt.Println("Failed to connect to the Ethereum client:", err)
	}

	privateKey, err := crypto.HexToECDSA("your-private-key-here")
	if err != nil {
		fmt.Println("Failed to get private key:", err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		fmt.Println("Failed to convert public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		fmt.Println("Failed to get nonce:", err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		fmt.Println("Failed to get suggested gas price:", err)
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	address := common.HexToAddress("0x38fDE154dCfE43A069e04e8A2F2E0C8247991D3f")
	instance, err := HelloWorld.NewHelloWorld(address, client)
	if err != nil {
		fmt.Println("Failed to instantiate contract:", err)
	}

	fmt.Println("Contract is loaded")
	message, err := instance.Message(nil)
	if err != nil {
		fmt.Println("Failed to retrieve message from contract:", err)
	}
	fmt.Println("Message from contract:", message)

	newMessage := "Hello, world!"
	tx, err := instance.SetMessage(auth, newMessage)
	if err != nil {
		fmt.Println("Failed to set message in contract:", err)
	}
	fmt.Printf("Transaction sent: %s", tx.Hash().Hex())

	updatedMessage, err := instance.Message(nil)
	if err != nil {
		fmt.Println("Failed to retrieve message from contract:", err)
	}
	fmt.Println("Updated message from contract:", updatedMessage)
}
