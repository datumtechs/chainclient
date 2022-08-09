package chainclient

import (
	"context"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"time"
)

type Context interface {
	SetPrivateKey(*ecdsa.PrivateKey)
	GetPrivateKey() *ecdsa.PrivateKey
	GetPublicKey() *ecdsa.PublicKey
	GetAddress() ethcommon.Address
	GetClient() *ethclient.Client
	BlockNumber(timeoutCtx context.Context) (uint64, error)
	PendingNonceAt(timeoutCtx context.Context) (uint64, error)
	SuggestGasPrice(timeoutCtx context.Context) (*big.Int, error)

	EstimateGas(timeoutCtx context.Context, to ethcommon.Address, data []byte) (uint64, error)
	CallContract(timeoutCtx context.Context, to ethcommon.Address, data []byte) ([]byte, error)

	BuildTxOpts(value, gasLimit uint64) (*bind.TransactOpts, error)
	WaitReceipt(timeoutCtx context.Context, txHash ethcommon.Hash, interval time.Duration) *ethtypes.Receipt
	GetLog(timeoutCtx context.Context, toAddr ethcommon.Address, blockNo *big.Int) []ethtypes.Log
	//GetLogs(timeoutCtx context.Context, toAddr common.Address, blockNo *big.Int) []*types.Log
}
