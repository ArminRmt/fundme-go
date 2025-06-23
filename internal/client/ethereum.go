package client

import (
    "context"
    "crypto/ecdsa"
    "math/big"

    "github.com/ethereum/go-ethereum/accounts/abi/bind"
    "github.com/ethereum/go-ethereum/crypto"
    "github.com/ethereum/go-ethereum/ethclient"
)

type EthereumClient struct {
    Client     *ethclient.Client
    PrivateKey *ecdsa.PrivateKey
    Auth       *bind.TransactOpts
}

func NewEthereumClient(rpcURL, privateKeyHex string, chainID int64) (*EthereumClient, error) {
    client, err := ethclient.Dial(rpcURL)
    if err != nil {
        return nil, err
    }

    privateKey, err := crypto.HexToECDSA(privateKeyHex)
    if err != nil {
        return nil, err
    }

    auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(chainID))
    if err != nil {
        return nil, err
    }

    return &EthereumClient{
        Client:     client,
        PrivateKey: privateKey,
        Auth:       auth,
    }, nil
}

func (ec *EthereumClient) UpdateGasSettings(gasLimit uint64, gasPrice int64) {
    ec.Auth.GasLimit = gasLimit
    ec.Auth.GasPrice = big.NewInt(gasPrice)
}

// ordering, prevents fraud (double-spending), integrity
// timeout for Network requests (context)
func (ec *EthereumClient) GetNonce(ctx context.Context) error {
    publicKey := ec.PrivateKey.Public()
    publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
    if !ok {
        return fmt.Errorf("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
    }

    fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA) // sender's address
    nonce, err := ec.Client.PendingNonceAt(ctx, fromAddress)
    if err != nil {
        return err
    }

    ec.Auth.Nonce = big.NewInt(int64(nonce))
    return nil
}
