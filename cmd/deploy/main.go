package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"os"
	"strings"

	"github.com/ArminRmt/fundme-go/internal/client"
	"github.com/ArminRmt/fundme-go/internal/config"
	"github.com/ArminRmt/fundme-go/pkg/contracts/generated/fundme"
	"github.com/ArminRmt/fundme-go/pkg/contracts/generated/mockaggregator"
)

func main() {
    cfg, err := config.LoadConfig()
    if err != nil {
        log.Fatal("Failed to load config:", err)
    }

    ethClient, err := client.NewEthereumClient(cfg.EthereumRPC, cfg.PrivateKey, cfg.ChainID)
    if err != nil {
        log.Fatal("Failed to create Ethereum client:", err)
    }
    // No defer ethClient.Client.Close() here, as this program exits shortly

    ctx := context.Background()
    err = ethClient.GetNonce(ctx)
    if err != nil {
        log.Fatal("Failed to get nonce:", err)
    }

    ethClient.UpdateGasSettings(cfg.GasLimit, cfg.GasPrice)

    //  simplified, dummy version of the real Chainlink V3Aggregator contract
    // get price from an external, on-chain source called a price feed oracle
    mockPriceFeedAddress, mockTx, _, err := mockaggregator.DeployMockV3Aggregator(
        ethClient.Auth,
        ethClient.Client,
        uint8(8),              // decimals
        big.NewInt(200000000000), // $2000 initial price
    )
    if err != nil {
        log.Fatal("Failed to deploy MockV3Aggregator:", err)
    }

    fmt.Printf("MockV3Aggregator deployed at: %s\n", mockPriceFeedAddress.Hex())
    fmt.Printf("MockV3Aggregator transaction hash: %s\n", mockTx.Hash().Hex())

    // Update nonce for next transaction
    ethClient.Auth.Nonce.Add(ethClient.Auth.Nonce, big.NewInt(1))

    // Deploy the FundMe contract with the mock price feed address
    fundMeAddress, fundMeTx, fundMeInstance, err := fundme.DeployFundMe(
        ethClient.Auth,
        ethClient.Client,
        mockPriceFeedAddress,
    )
    if err != nil {
        log.Fatal("Failed to deploy FundMe:", err)
    }

    fmt.Printf("FundMe contract deployed at: %s\n", fundMeAddress.Hex())
    fmt.Printf("FundMe transaction hash: %s\n", fundMeTx.Hash().Hex())
    
    
    err = saveContractAddress(".env", fundMeAddress.Hex())
	if err != nil {
		log.Fatalf("Failed to save contract address to .env file: %v", err)
	}

	fmt.Println("✅ Contracts deployed successfully!")

    // Go compiler complain about an unused variable (Store for future use)
    _ = fundMeInstance
}


func saveContractAddress(envFilePath, address string) error {
	const keyToUpdate = "CONTRACT_ADDRESS"

	content, err := os.ReadFile(envFilePath)
	if err != nil {
		if os.IsNotExist(err) {
			newContent := fmt.Sprintf("%s=%s\n", keyToUpdate, address)
			return os.WriteFile(envFilePath, []byte(newContent), 0644)
		}
		return fmt.Errorf("failed to read .env file: %w", err)
	}

	lines := strings.Split(string(content), "\n")
	keyFound := false

	for i, line := range lines {
		if strings.HasPrefix(line, keyToUpdate+"=") {
			lines[i] = fmt.Sprintf("%s=%s", keyToUpdate, address)
			keyFound = true
			break 
		}
	}

	// If the key was not found, append it to the end
	if !keyFound {
		lastLine := ""
		if len(lines) > 0 {
			lastLine = lines[len(lines)-1]
		}
		if lastLine != "" {
			lines = append(lines, "")
		}
		lines = append(lines, fmt.Sprintf("%s=%s", keyToUpdate, address))
	}

	output := strings.Join(lines, "\n")
	err = os.WriteFile(envFilePath, []byte(output), 0644)
	if err != nil {
		return fmt.Errorf("failed to write updated content to .env file: %w", err)
	}

	return nil
}