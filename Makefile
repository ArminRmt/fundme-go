# commands run regardless of whether a file with that name exists
.PHONY: compile-contracts install-deps generate-bindings deploy fund refund withdraw status test clean setup

install-deps:
	go mod download

# Compile Solidity contracts
compile-contracts:
	solc --abi --evm-version london contracts/FundMe.sol -o build --overwrite
	solc --bin --evm-version london contracts/FundMe.sol -o build --overwrite
	solc --abi --evm-version london contracts/MockV3Aggregator.sol -o build --overwrite
	solc --bin --evm-version london contracts/MockV3Aggregator.sol -o build --overwrite

# Generate Go bindings
generate-bindings:
	mkdir -p pkg/contracts/generated/fundme
	go run github.com/ethereum/go-ethereum/cmd/abigen@v1.13.0 --bin=build/FundMe.bin --abi=build/FundMe.abi --pkg=fundme --type FundMe --out=pkg/contracts/generated/fundme/fundme.go
	mkdir -p pkg/contracts/generated/mockaggregator
	go run github.com/ethereum/go-ethereum/cmd/abigen@v1.13.0 --bin=build/MockV3Aggregator.bin --abi=build/MockV3Aggregator.abi --pkg=mockaggregator --type MockV3Aggregator --out=pkg/contracts/generated/mockaggregator/mockaggregator.go


deploy:
	go run cmd/deploy/main.go

fund:
	go run cmd/fund/main.go

refund:
	go run cmd/refund/main.go  

withdraw:
	go run cmd/withdraw/main.go

status:
	go run cmd/status/main.go

test:
	go test ./test/... -v

# Start Ganache (local Ethereum blockchain simulator)
# deterministic: same accounts and private key each time
# accounts: 10 test accounts with 100 ETH each
# --host 0.0.0.0: allows RPC server accessible across local network
start-ganache:
	ganache-cli --deterministic --accounts 10 --host 0.0.0.0 --port 8545

clean:
	rm -rf build/
	rm -rf pkg/contracts/generated/

setup: install-deps compile-contracts generate-bindings
	go mod tidy
