package test

import (
	"context"
	"math/big"
	"testing"
	"time"

	"github.com/ArminRmt/fundme-go/pkg/contracts/generated/fundme"
	"github.com/ArminRmt/fundme-go/pkg/contracts/generated/mockaggregator"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

func TestFundMeContract(t *testing.T) {
	// Setup simulated backend
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		t.Fatal(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(1337))
	if err != nil {
		t.Fatal(err)
	}

	// Set reasonable gas fees
	auth.GasFeeCap = big.NewInt(10000000000) // 10 gwei
	auth.GasTipCap = big.NewInt(2000000000)  // 2 gwei

	balance := new(big.Int)
	balance.SetString("10000000000000000000", 10) // 10 ETH

	address := auth.From
	genesisAlloc := map[common.Address]core.GenesisAccount{
		address: {
			Balance: balance,
		},
	}

	backend := backends.NewSimulatedBackend(genesisAlloc, 10000000)

	// Deploy MockV3Aggregator
	mockAddress, _, _, err := mockaggregator.DeployMockV3Aggregator(
		auth,
		backend,
		uint8(8),
		big.NewInt(200000000000), // $2000
	)
	if err != nil {
		t.Fatal(err)
	}
	backend.Commit()

	// Deploy FundMe
	fundMeAddress, _, fundMeInstance, err := fundme.DeployFundMe(
		auth,
		backend,
		mockAddress,
	)
	if err != nil {
		t.Fatal(err)
	}
	backend.Commit()

	t.Run("Funding with sufficient ETH", func(t *testing.T) {
		// Fund with 0.1 ETH ($200 at $2000/ETH)
		auth.Value = big.NewInt(100000000000000000) // 0.1 ETH
		_, err := fundMeInstance.Fund(auth)
		if err != nil {
			t.Fatal("Failed to fund contract:", err)
		}
		backend.Commit()

		// Verify contract balance
		contractBalance, err := backend.BalanceAt(nil, fundMeAddress, nil)
		if err != nil {
			t.Fatal("Failed to get contract balance:", err)
		}
		if contractBalance.Cmp(auth.Value) != 0 {
			t.Fatalf("Contract balance incorrect: expected %v, got %v", auth.Value, contractBalance)
		}

		// Verify funder tracking
		contribution, err := fundMeInstance.GetAddressToAmountFunded(&bind.CallOpts{}, auth.From)
		if err != nil {
			t.Fatal(err)
		}
		if contribution.Cmp(auth.Value) != 0 {
			t.Fatalf("Contribution not recorded: expected %v, got %v", auth.Value, contribution)
		}

		// Verify funder in array
		funder0, err := fundMeInstance.GetFunder(&bind.CallOpts{}, big.NewInt(0))
		if err != nil {
			t.Fatal(err)
		}
		if funder0 != auth.From {
			t.Fatalf("Funder not in array: expected %v, got %v", auth.From, funder0)
		}
	})

	t.Run("Funding with insufficient ETH", func(t *testing.T) {
		// Reset value from previous test
		auth.Value = big.NewInt(0)
		backend.Commit()

		// Fund with 0.001 ETH ($2 at $2000/ETH)
		auth.Value = big.NewInt(1000000000000000) // 0.001 ETH
		_, err := fundMeInstance.Fund(auth)
		if err == nil {
			t.Fatal("Expected funding to revert with insufficient ETH, but succeeded")
		}
		backend.Commit()

		// Verify contract state unchanged
		contractBalance, err := backend.BalanceAt(nil, fundMeAddress, nil)
		if err != nil {
			t.Fatal("Failed to get contract balance:", err)
		}
		expectedBalance := big.NewInt(100000000000000000) // Previous 0.1 ETH
		if contractBalance.Cmp(expectedBalance) != 0 {
			t.Fatalf("Contract balance changed: expected %v, got %v", expectedBalance, contractBalance)
		}
	})

	t.Run("Owner withdrawal with single funder", func(t *testing.T) {
		// Reset value
		auth.Value = big.NewInt(0)
		backend.Commit()

		// Get initial owner balance
		initialBalance, err := backend.BalanceAt(nil, auth.From, nil)
		if err != nil {
			t.Fatal("Failed to get initial owner balance:", err)
		}

		// Advance time beyond funding period
		backend.AdjustTime(7 * 24 * time.Hour + time.Second)
		backend.Commit()

		// Withdraw funds
		withdrawAuth := *auth
		withdrawAuth.GasLimit = 300000
		tx, err := fundMeInstance.Withdraw(&withdrawAuth)
		if err != nil {
			t.Fatal("Failed to withdraw from contract:", err)
		}
		backend.Commit()

		// Verify transaction success
		receipt, err := backend.TransactionReceipt(context.Background(), tx.Hash())
		if err != nil {
			t.Fatal("Failed to get receipt:", err)
		}
		if receipt.Status == types.ReceiptStatusFailed {
			t.Fatal("Withdrawal transaction failed")
		}

		// Verify contract balance
		contractBalance, err := backend.BalanceAt(nil, fundMeAddress, nil)
		if err != nil {
			t.Fatal("Failed to get contract balance:", err)
		}
		if contractBalance.Cmp(big.NewInt(0)) != 0 {
			t.Fatalf("Contract balance not zero: %v", contractBalance)
		}

		// Verify owner received funds
		finalBalance, err := backend.BalanceAt(nil, auth.From, nil)
		if err != nil {
			t.Fatal("Failed to get final owner balance:", err)
		}
		if finalBalance.Cmp(initialBalance) <= 0 {
			t.Fatalf("Owner balance not increased: initial %v, final %v", initialBalance, finalBalance)
		}

		// Verify state reset
		contribution, err := fundMeInstance.GetAddressToAmountFunded(&bind.CallOpts{}, auth.From)
		if err != nil {
			t.Fatal(err)
		}
		if contribution.Cmp(big.NewInt(0)) != 0 {
			t.Fatalf("Funder contribution not reset: %v", contribution)
		}
	})

	t.Run("Non-owner withdrawal attempt", func(t *testing.T) {
		// Create non-owner account
		nonOwnerKey, err := crypto.GenerateKey()
		if err != nil {
			t.Fatal(err)
		}
		nonOwnerAuth, err := bind.NewKeyedTransactorWithChainID(nonOwnerKey, big.NewInt(1337))
		if err != nil {
			t.Fatal(err)
		}

		// Set reasonable gas fees for non-owner
		nonOwnerAuth.GasFeeCap = big.NewInt(10000000000)
		nonOwnerAuth.GasTipCap = big.NewInt(2000000000)

		// Fund non-owner account by sending ETH from owner
		nonce, err := backend.PendingNonceAt(nil, auth.From)
		if err != nil {
			t.Fatal(err)
		}
		tx := types.NewTransaction(
			nonce,
			nonOwnerAuth.From,
			big.NewInt(1000000000000000000), // 1 ETH
			21000,
			big.NewInt(1000000000), // 1 gwei
			nil,
		)
		signedTx, err := auth.Signer(auth.From, tx)
		if err != nil {
			t.Fatal(err)
		}
		err = backend.SendTransaction(nil, signedTx)
		if err != nil {
			t.Fatal(err)
		}
		backend.Commit()

		// Advance time beyond funding period
		backend.AdjustTime(7 * 24 * time.Hour + time.Second)
		backend.Commit()

		// Attempt withdrawal
		withdrawAuth := *nonOwnerAuth
		withdrawAuth.GasLimit = 300000
		tx, err = fundMeInstance.Withdraw(&withdrawAuth)
		if err != nil {
			t.Fatal("Failed to send withdrawal transaction:", err)
		}
		backend.Commit()

		// Verify transaction failed
		receipt, err := backend.TransactionReceipt(context.Background(), tx.Hash())
		if err != nil {
			t.Fatal("Failed to get receipt:", err)
		}
		if receipt.Status == types.ReceiptStatusSuccessful {
			t.Fatal("Expected withdrawal to revert for non-owner, but succeeded")
		}
	})

	t.Run("Multiple funders and withdrawal", func(t *testing.T) {
		// Reset contract state by redeploying
		fundMeAddress, _, fundMeInstance, err = fundme.DeployFundMe(
			auth,
			backend,
			mockAddress,
		)
		if err != nil {
			t.Fatal(err)
		}
		backend.Commit()

		// Create second funder
		funder2Key, err := crypto.GenerateKey()
		if err != nil {
			t.Fatal(err)
		}
		funder2Auth, err := bind.NewKeyedTransactorWithChainID(funder2Key, big.NewInt(1337))
		if err != nil {
			t.Fatal(err)
		}
		
		// Fund funder2
		nonce, err := backend.PendingNonceAt(nil, auth.From)
		if err != nil {
			t.Fatal(err)
		}
		tx := types.NewTransaction(
			nonce,
			funder2Auth.From,
			big.NewInt(1000000000000000000), // 1 ETH
			21000,
			big.NewInt(1000000000), // 1 gwei
			nil,
		)
		signedTx, err := auth.Signer(auth.From, tx)
		if err != nil {
			t.Fatal(err)
		}
		err = backend.SendTransaction(nil, signedTx)
		if err != nil {
			t.Fatal(err)
		}
		backend.Commit()

		// First funder (owner)
		auth.Value = big.NewInt(50000000000000000) // 0.05 ETH
		_, err = fundMeInstance.Fund(auth)
		if err != nil {
			t.Fatal(err)
		}
		backend.Commit()

		// Second funder
		funder2Auth.Value = big.NewInt(75000000000000000) // 0.075 ETH
		_, err = fundMeInstance.Fund(funder2Auth)
		if err != nil {
			t.Fatal(err)
		}
		backend.Commit()

		// Verify contract balance
		expectedBalance := big.NewInt(0).Add(auth.Value, funder2Auth.Value)
		contractBalance, err := backend.BalanceAt(nil, fundMeAddress, nil)
		if err != nil {
			t.Fatal("Failed to get contract balance:", err)
		}
		if contractBalance.Cmp(expectedBalance) != 0 {
			t.Fatalf("Contract balance incorrect: expected %v, got %v", expectedBalance, contractBalance)
		}

		// Verify funders array
		funder0, err := fundMeInstance.GetFunder(&bind.CallOpts{}, big.NewInt(0))
		if err != nil {
			t.Fatal(err)
		}
		if funder0 != auth.From {
			t.Fatalf("First funder incorrect: expected %v, got %v", auth.From, funder0)
		}

		funder1, err := fundMeInstance.GetFunder(&bind.CallOpts{}, big.NewInt(1))
		if err != nil {
			t.Fatal(err)
		}
		if funder1 != funder2Auth.From {
			t.Fatalf("Second funder incorrect: expected %v, got %v", funder2Auth.From, funder1)
		}

		// Advance time beyond funding period
		backend.AdjustTime(7 * 24 * time.Hour + time.Second)
		backend.Commit()

		// Withdraw funds
		auth.Value = big.NewInt(0)
		_, err = fundMeInstance.Withdraw(auth)
		if err != nil {
			t.Fatal("Failed to withdraw from contract:", err)
		}
		backend.Commit()

		// Verify state reset
		contribution, err := fundMeInstance.GetAddressToAmountFunded(&bind.CallOpts{}, auth.From)
		if err != nil {
			t.Fatal(err)
		}
		if contribution.Cmp(big.NewInt(0)) != 0 {
			t.Fatalf("Owner contribution not reset: %v", contribution)
		}

		contribution2, err := fundMeInstance.GetAddressToAmountFunded(&bind.CallOpts{}, funder2Auth.From)
		if err != nil {
			t.Fatal(err)
		}
		if contribution2.Cmp(big.NewInt(0)) != 0 {
			t.Fatalf("Funder2 contribution not reset: %v", contribution2)
		}
	})

	t.Run("Refund functionality", func(t *testing.T) {
		// Redeploy fresh contract
		fundMeAddress, _, fundMeInstance, err = fundme.DeployFundMe(
			auth,
			backend,
			mockAddress,
		)
		if err != nil {
			t.Fatal(err)
		}
		backend.Commit()

		// Fund with 0.1 ETH
		auth.Value = big.NewInt(100000000000000000) // 0.1 ETH
		_, err = fundMeInstance.Fund(auth)
		if err != nil {
			t.Fatal("Failed to fund contract:", err)
		}
		backend.Commit()

		// Get initial balance
		initialBalance, err := backend.BalanceAt(nil, auth.From, nil)
		if err != nil {
			t.Fatal("Failed to get initial balance:", err)
		}

		// Process refund with sufficient gas
		refundAuth := *auth
		refundAuth.GasLimit = 300000  // Set sufficient gas limit
		refundAuth.Value = big.NewInt(0)
		tx, err := fundMeInstance.Refund(&refundAuth)
		if err != nil {
			t.Fatal("Failed to process refund:", err)
		}
		backend.Commit()

		// Verify transaction success
		receipt, err := backend.TransactionReceipt(context.Background(), tx.Hash())
		if err != nil {
			t.Fatal("Failed to get receipt:", err)
		}
		if receipt.Status == types.ReceiptStatusFailed {
			t.Fatal("Refund transaction failed")
		}

		// Verify balance increased
		finalBalance, err := backend.BalanceAt(nil, auth.From, nil)
		if err != nil {
			t.Fatal("Failed to get final balance:", err)
		}
		
		// Calculate expected balance: initial + refunded amount - gas costs
		// Note: Gas costs will reduce the balance, so we only check that it's greater than initial
		if finalBalance.Cmp(initialBalance) <= 0 {
			t.Fatalf("Balance not increased after refund: initial %v, final %v", initialBalance, finalBalance)
		}

		// Verify state reset
		contribution, err := fundMeInstance.GetAddressToAmountFunded(&bind.CallOpts{}, auth.From)
		if err != nil {
			t.Fatal(err)
		}
		if contribution.Cmp(big.NewInt(0)) != 0 {
			t.Fatalf("Contribution not reset: %v", contribution)
		}

		funderCount := 0
		for {
			_, err := fundMeInstance.GetFunder(&bind.CallOpts{}, big.NewInt(int64(funderCount)))
			if err != nil {
				break
			}
			funderCount++
		}
		if funderCount != 0 {
			t.Fatalf("Funder array not reset: %v funders remaining", funderCount)
		}
	})
}