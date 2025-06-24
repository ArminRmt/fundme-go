package main

import (
	"context"
	"fmt"
	"log"
	"math/big"

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

    // Fund with 0.1 ETH
    ethClient.Auth.Value = big.NewInt(100000000000000000)

    tx, err := fundMe.Fund(ethClient.Auth)
    if err != nil {
        log.Fatal("Failed to fund contract:", err)
    }

    fmt.Printf("Funding transaction hash: %s\n", tx.Hash().Hex())
    fmt.Println("💸 Successfully funded the contract with 0.1 ETH!")

    // Reset value for future transactions
    ethClient.Auth.Value = big.NewInt(0)
}
