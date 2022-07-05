package chainclient

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Context interface {
	GetPrivateKey() *ecdsa.PrivateKey
	GetWalletAddress() common.Address
	GetClient() *ethclient.Client
	BuildTxOpts(uint64, gasLimit uint64) (*bind.TransactOpts, error)
}
