package test

import (
	"math/big"
	"testing"

	"github.com/ArminRmt/fundme-go/pkg/contracts/generated/fundme"
	"github.com/ArminRmt/fundme-go/pkg/contracts/generated/mockaggregator"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
)

func TestFundMeContract(t *testing.T) {
    // Setup simulated backend
    privateKey, err := crypto.GenerateKey()
    if err != nil {
        t.Fatal(err)
    }

	// ethClient.Auth
    auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(1337))
    if err != nil {
        t.Fatal(err)
    }

    balance := new(big.Int)
    balance.SetString("10000000000000000000", 10) // 10 ETH

	// initial state of the simulated blockchain.
    address := auth.From
    genesisAlloc := map[common.Address]core.GenesisAccount{
        address: {
            Balance: balance,
        },
    }

	// in-memory blockchain
    backend := backends.NewSimulatedBackend(genesisAlloc, 10000000)

    // Deploy MockV3Aggregator	
    mockAddress, _, _, err := mockaggregator.DeployMockV3Aggregator(
        auth,
        backend,
        uint8(8),
        big.NewInt(200000000000),
    )
    if err != nil {
        t.Fatal(err)
    }
    backend.Commit()

     // Deploy FundMe
    _, _, fundMeInstance, err := fundme.DeployFundMe(
        auth,
        backend,
        mockAddress,
    )
    if err != nil {
        t.Fatal(err)
    }
    backend.Commit()

    // Test funding
    auth.Value = big.NewInt(100000000000000000) // 0.1 ETH
    _, err = fundMeInstance.Fund(auth)
    if err != nil {
        t.Fatal("Failed to fund contract:", err)
    }
    backend.Commit()

    // Test withdrawal (should succeed as owner
    auth.Value = big.NewInt(0)
    _, err = fundMeInstance.Withdraw(auth)
    if err != nil {
        t.Fatal("Failed to withdraw from contract:", err)
    }
    backend.Commit()

    t.Log("All tests passed!")
}