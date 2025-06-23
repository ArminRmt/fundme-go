# commands run regardless of whether a file with that name exists
.PHONY: compile-contracts install-deps generate-bindings deploy fund withdraw test clean

install-deps:
	go mod tidy
	go mod download

# Compile Solidity contracts
# abi: JSON-formatted description of a smart contract for frontend, or another smart contracts
# bin: bytecode which deployed onto the blockchain
compile-contracts:
	solc --abi contracts/FundMe.sol -o build --overwrite
	solc --bin contracts/FundMe.sol -o build --overwrite
	solc --abi contracts/MockV3Aggregator.sol -o ``build --overwrite
	solc --bin contracts/MockV3Aggregator.sol -o build --overwrite

# Generate Go bindings
generate-bindings:
	mkdir -p pkg/contracts/generated
	abigen --bin=build/FundMe.bin --abi=build/FundMe.abi --pkg=generated --out=pkg/contracts/generated/FundMe.go
	abigen --bin=build/MockV3Aggregator.bin --abi=build/MockV3Aggregator.abi --pkg=generated --out=pkg/contracts/generated/MockV3Aggregator.go

# Deploy contracts
deploy:
	go run cmd/deploy/main.go

# Fund the contract
fund:
	go run cmd/fund/main.go

# Withdraw from contract
withdraw:
	go run cmd/withdraw/main.go

# Run tests
test:
	go test ./test/... -v

# Start Ganache (local Ethereum blockchain simulator)
# deterministic: same accounts and private key each time
# accounts: 10 test accounts with 100 ETH each
# --host 0.0.0.0: allows RPC server accessible across local network
start-ganache:
	ganache-cli --deterministic --accounts 10 --host 0.0.0.0 --port 8545

# Clean build artifacts
clean:
	rm -rf build/
	rm -rf pkg/contracts/generated/

# Full setup
setup: install-deps compile-contracts generate-bindings
