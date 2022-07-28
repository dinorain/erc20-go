package main

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/dinorain/erc20-go/erc20"
)

type Config struct {
	Network         string `yaml:"network"`
	ContractAddress string `yaml:"contract_address"`
}

func main() {
	conf := Config{
		Network:         "https://goerli.infura.io",
		ContractAddress: "",
	}
	client, err := ethclient.Dial(conf.Network)
	if err != nil {
		fmt.Printf("Failed to connect to eth: %v", err)
		return
	}
	token, err := erc20.NewToken(common.HexToAddress(conf.ContractAddress), client)
	if err != nil {
		fmt.Printf("Failed to instantiate a Token contract: %v", err)
		return
	}

	totalSupply, err := token.TotalSupply(nil)
	if err != nil {
		fmt.Printf("Failed to get name: %v", err)
		return
	}
	fmt.Printf("totalSupply: %v\n", totalSupply.String())
}