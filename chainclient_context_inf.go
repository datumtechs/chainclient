package chainclient

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
)

type Context interface {
	GetPrivateKey() *ecdsa.PrivateKey
	GetWalletAddress() common.Address
	GetClient() *ethclient.Client
	EstimateGas(from common.Address, to *common.Address, data []byte, gas uint64, gasPrice *big.Int) (uint64, error)
	BuildTxOpts(value, gasLimit uint64) (*bind.TransactOpts, error)
}
