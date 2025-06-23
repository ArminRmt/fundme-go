package config

import (
    "os"
    "strconv"
    "github.com/joho/godotenv"
)

type Config struct {
    EthereumRPC     string
    PrivateKey      string
    ContractAddress string
    ChainID         int64
    GasLimit        uint64
    GasPrice        int64
}

func LoadConfig() (*Config, error) {
    err := godotenv.Load()
    if err != nil {
        return nil, err
    }

    chainID, _ := strconv.ParseInt(getEnv("CHAIN_ID", "1337"), 10, 64)
    gasLimit, _ := strconv.ParseUint(getEnv("GAS_LIMIT", "300000"), 10, 64)
    gasPrice, _ := strconv.ParseInt(getEnv("GAS_PRICE", "20000000000"), 10, 64)

    return &Config{
        EthereumRPC:     getEnv("ETHEREUM_RPC", "http://localhost:8545"),
        PrivateKey:      getEnv("PRIVATE_KEY", ""),
        ContractAddress: getEnv("CONTRACT_ADDRESS", ""),
        ChainID:         chainID,
        GasLimit:        gasLimit,
        GasPrice:        gasPrice,
    }, nil
}

func getEnv(key, defaultValue string) string {
    if value := os.Getenv(key); value != "" {
        return value
    }
    return defaultValue
}
