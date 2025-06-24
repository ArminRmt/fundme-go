# FundMe Smart Contract Project

![Solidity](https://img.shields.io/badge/Solidity-%23363636.svg?style=for-the-badge&logo=solidity&logoColor=white)
![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
![Foundry](https://img.shields.io/badge/Foundry-%23FF5E1A.svg?style=for-the-badge)

FundMe is a Solidity smart contract that allows users to send Ether to a decentralized funding pool with enforced minimum contributions. Only the contract owner can withdraw funds after the funding period ends. This project includes comprehensive tests, deployment scripts, and interaction tools.

## Features

- 💰 Accept ETH contributions with minimum USD value enforcement
- 🔒 Owner-only withdrawal functionality
- 📊 Tracks individual funders and their contributions
- ⏱️ Time-limited funding period (bonus feature)
- 🔁 User refund capability before deadline (bonus feature)
- 🧪 Comprehensive test suite with Foundry
- 🤖 Automated deployment and interaction scripts

## Prerequisites

- [Go](https://golang.org/dl/) (v1.19+)
- [Foundry](https://getfoundry.sh/)
- [Ganache CLI](https://www.npmjs.com/package/ganache)
- [solc](https://docs.soliditylang.org/en/latest/installing-solidity.html)
- [abigen](https://geth.ethereum.org/docs/tools/abigen)

## Project Structure

```
fundme-project/
├── contracts/              # Solidity smart contracts
│   ├── FundMe.sol
│   └── MockV3Aggregator.sol
├── pkg/                    # Generated Go bindings
│   └── contracts/generated/
├── cmd/                    # Interaction scripts
│   ├── deploy/
│   │   └── main.go
│   ├── fund/
│   │   └── main.go
│   └── withdraw/
│   │   └── main.go
│   └── refund/
│       └── main.go
├── test/                   # Go tests
│   └── fundme_test.go
├── Makefile                # Build automation
└── .env.example            # Environment template
```

## Getting Started

### 1. Clone the repository
```bash
git clone https://github.com/your-username/fundme-project.git
cd fundme-project
```

### 2. Set up environment
```bash
cp .env.example .env
# Edit .env with your private key from ganache
```

### 3. Install dependencies,compile-contracts,generate-bindings
```bash
make setup
```

### 4. Start local blockchain
```bash
make start-ganache
```

## Usage

### Deploy the contract
```bash
make deploy
```
Example output:
```
MockV3Aggregator deployed at: 0x5FbDB2315678afecb367f032d93F642f64180aa3
MockV3Aggregator transaction hash: 0x7a9e8e1d6a7e5d4c3b2a1f0e9d8c7b6a5a4b3c2d1e0f9a8b7c6d5e4f3a2b1c0d
FundMe contract deployed at: 0xe7f1725E7734CE288F8367e1Bb143E90bb3F0512
FundMe transaction hash: 0x6a5a4b3c2d1e0f9a8b7c6d5e4f3a2b1c0d9e8f7a6b5c4d3e2f1a0b9c8d7e6f5d4
Save this contract address to your .env file:
CONTRACT_ADDRESS=0xe7f1725E7734CE288F8367e1Bb143E90bb3F0512
```


### Fund the contract (send ETH)
```bash
make fund
```
Example output:
```
Funding transaction hash: 0x3a2b1c0d9e8f7a6b5c4d3e2f1a0b9c8d7e6f5d4c3b2a1f0e9d8c7b6a5a4b3c2d1
Successfully funded the contract with 0.1 ETH!
```    

### Withdraw funds (owner only)
```bash
make withdraw
```
Example output:
```
Withdrawal transaction hash: 0xf1a0b9c8d7e6f5d4c3b2a1f0e9d8c7b6a5a4b3c2d1e0f9a8b7c6d5e4f3a2b1c0d
Successfully withdrew all funds from the contract!
```   

### Refund contributions
```
make refund
```
Example output:
```
Refund transaction hash: 0xa74dbc595d22d99d341425ce39ddade6440433270ce84c11250a1f2bf20a2e11
Successfully processed refund!
```    


1. **Funding Validation**:
   - ✅ Funding with sufficient ETH (meets $5 minimum)
   - ✅ Funding with insufficient ETH (properly reverts)

2. **Withdrawal Functionality**:
   - ✅ Owner withdrawal after funding period ends
   - ❌ Non-owner withdrawal attempts (properly reverts)
   - ✅ Multiple funders withdrawal with state reset

3. **Advanced Features**:
   - ✅ Refund functionality during funding period
   - ✅ Time-based restrictions enforcement
   - ✅ Funder tracking and state management

### Test Execution

```bash
# Run all tests
make test

# Sample output
=== RUN   TestFundMeContract
=== RUN   TestFundMeContract/Funding_with_sufficient_ETH
=== RUN   TestFundMeContract/Funding_with_insufficient_ETH
=== RUN   TestFundMeContract/Owner_withdrawal_with_single_funder
=== RUN   TestFundMeContract/Non-owner_withdrawal_attempt
=== RUN   TestFundMeContract/Multiple_funders_and_withdrawal
=== RUN   TestFundMeContract/Refund_functionality
--- PASS: TestFundMeContract (0.04s)
    --- PASS: TestFundMeContract/Funding_with_sufficient_ETH (0.00s)
    --- PASS: TestFundMeContract/Funding_with_insufficient_ETH (0.00s)
    --- PASS: TestFundMeContract/Owner_withdrawal_with_single_funder (0.00s)
    --- PASS: TestFundMeContract/Non-owner_withdrawal_attempt (0.00s)
    --- PASS: TestFundMeContract/Multiple_funders_and_withdrawal (0.01s)
    --- PASS: TestFundMeContract/Refund_functionality (0.00s)
PASS
```

### Test Features

- ⏱️ Time manipulation for funding period validation
- 🔄 State reset verification after operations
- ⛽ Gas optimization checks
- 🔐 Permission testing (owner vs non-owner)
- 💰 Balance verification pre/post transactions
- 🧪 Isolated test environments with contract redeployment

```   

### Clean build artifacts    
```bash
make clean
```    
     
## Key Commands

| Command              | Description                               |
| -------------------- | ----------------------------------------- |
| `make setup`         | Install dependencies & generate bindings  |
| `make start-ganache` | Start local Ethereum blockchain           |
| `make deploy`        | Deploy contract to local network          |
| `make fund`          | Send ETH to contract                      |
| `make refund`        | Refund contributions                      |
| `make withdraw`      | Withdraw funds from contract (owner only) |
| `make test`          | Run all tests                             |
| `make clean`         | Remove build artifacts                    |
```   

## Smart Contract Features

### Core Functionality
- **Payable `fund()` function**: Accepts ETH contributions
- **Owner-restricted `withdraw()`**: Only contract deployer can withdraw
- **Minimum USD contribution**: Enforced using price feeds
- **Funder tracking**: Records all contributors and amounts

### Advanced Features
- **Mock price feeds**: For local development and testing
- **Funding deadline**: Time-limited contribution period
- **Refund mechanism**: Allows contributors to reclaim funds before deadline

## Testing

Run the comprehensive test suite:
```bash
make test
```

Test coverage includes:
- Successful funding transactions
- Minimum contribution enforcement
- Owner withdrawal functionality
- Non-owner withdrawal prevention
- Funder tracking and state reset
- Time-based restrictions (bonus features)
- Refund functionality (bonus features)

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.