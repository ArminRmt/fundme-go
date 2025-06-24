package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/ArminRmt/fundme-go/internal/config"
	"github.com/ArminRmt/fundme-go/pkg/contracts/generated/fundme"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Failed to load config:", err)
	}

	if cfg.ContractAddress == "" {
		log.Fatal("CONTRACT_ADDRESS not set in environment")
	}

	client, err := ethclient.Dial(cfg.EthereumRPC)
	if err != nil {
		log.Fatal("Failed to connect to Ethereum client:", err)
	}

	ctx := context.Background()
	contractAddress := common.HexToAddress(cfg.ContractAddress)
	fundMe, err := fundme.NewFundMe(contractAddress, client)
	if err != nil {
		log.Fatal("Failed to instantiate FundMe contract:", err)
	}

	balance, err := client.BalanceAt(ctx, contractAddress, nil)
	if err != nil {
		log.Fatal("Failed to get contract balance:", err)
	}

	owner, err := fundMe.GetOwner(nil)
	if err != nil {
		log.Fatal("Failed to get owner:", err)
	}

	startTime, err := fundMe.SStartTime(nil)
	if err != nil {
		log.Fatal("Failed to get start time:", err)
	}

	fundingPeriod, err := fundMe.FUNDINGPERIOD(nil)
	if err != nil {
		log.Fatal("Failed to get funding period:", err)
	}

	deadline := new(big.Int).Add(startTime, fundingPeriod)
	deadlineTime := time.Unix(deadline.Int64(), 0)
	timeRemaining := deadlineTime.Sub(time.Now())

	funderCount := 0
	if balance.Cmp(big.NewInt(0)) > 0 {
		_, err := fundMe.GetFunder(nil, big.NewInt(0))
		if err == nil {
			funderCount = 1
		}
		
		funderCountNote := " (at least 1, full count disabled for Ganache stability)"
		if funderCount > 0 {
			fmt.Printf("👥 Funders: %d%s\n", funderCount, funderCountNote)
		} else {
			fmt.Printf("👥 Funders: 0\n")
		}
	} else {
		fmt.Printf("👥 Funders: 0\n")
	}

	fmt.Println("\n📊 FundMe Contract Status")
	fmt.Println("=========================")
	fmt.Printf("👤 Owner: %s\n", owner.Hex())
	fmt.Printf("💰 Balance: %s ETH\n", weiToEther(balance))
	fmt.Printf("⏱️  Funding Deadline: %s\n", deadlineTime.Format(time.RFC1123))

	if timeRemaining > 0 {
		fmt.Printf("🟢 Time Remaining: %s\n", formatDuration(timeRemaining))
	} else {
		fmt.Printf("🔴 Funding Closed: %s ago\n", formatDuration(-timeRemaining))
	}

	fmt.Println("=========================")
}

func weiToEther(wei *big.Int) string {
	ether := new(big.Float).Quo(new(big.Float).SetInt(wei), big.NewFloat(1e18))
	return ether.Text('f', 6)
}

func formatDuration(d time.Duration) string {
	d = d.Round(time.Minute)
	days := d / (24 * time.Hour)
	d -= days * 24 * time.Hour
	hours := d / time.Hour
	d -= hours * time.Hour
	minutes := d / time.Minute
	return fmt.Sprintf("%dd %dh %dm", days, hours, minutes)
}