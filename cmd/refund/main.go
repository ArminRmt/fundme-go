package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ArminRmt/fundme-go/internal/client"
	"github.com/ArminRmt/fundme-go/internal/config"
	"github.com/ArminRmt/fundme-go/pkg/contracts/generated/fundme"
	"github.com/ethereum/go-ethereum/common"
)

func main() {
    cfg, err := config.LoadConfig()
    if err != nil {
        log.Fatal("Failed to load config:", err)
    }

    if cfg.ContractAddress == "" {
        log.Fatal("CONTRACT_ADDRESS not set in environment")
    }

    ethClient, err := client.NewEthereumClient(cfg.EthereumRPC, cfg.PrivateKey, cfg.ChainID)
    if err != nil {
        log.Fatal("Failed to create Ethereum client:", err)
    }

    ctx := context.Background()
    err = ethClient.GetNonce(ctx)
    if err != nil {
        log.Fatal("Failed to get nonce:", err)
    }

    ethClient.UpdateGasSettings(cfg.GasLimit, cfg.GasPrice)

    contractAddress := common.HexToAddress(cfg.ContractAddress)
    fundMe, err := fundme.NewFundMe(contractAddress, ethClient.Client)
    if err != nil {
        log.Fatal("Failed to instantiate FundMe contract:", err)
    }

    tx, err := fundMe.Refund(ethClient.Auth)
    if err != nil {
        log.Fatal("Failed to process refund:", err)
    }

    fmt.Printf("Refund transaction hash: %s\n", tx.Hash().Hex())
    fmt.Println("🔄 Refund processed successfully!")
}