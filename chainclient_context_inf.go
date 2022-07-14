package chainclient

import (
	"context"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"time"
)

type Context interface {
	GetPrivateKey() *ecdsa.PrivateKey
	GetWalletAddress() common.Address
	GetClient() *ethclient.Client
	BlockNumber(timeoutCtx context.Context) (uint64, error)
	PendingNonceAt(timeoutCtx context.Context) (uint64, error)
	SuggestGasPrice(timeoutCtx context.Context) (*big.Int, error)

	EstimateGas(timeoutCtx context.Context, to common.Address, data []byte) (uint64, error)
	CallContract(timeoutCtx context.Context, to common.Address, data []byte) ([]byte, error)

	BuildTxOpts(value, gasLimit uint64) (*bind.TransactOpts, error)
	WaitReceipt(timeoutCtx context.Context, txHash common.Hash, interval time.Duration) *types.Receipt
	GetLog(timeoutCtx context.Context, toAddr common.Address, blockNo *big.Int) []*types.Log
}
