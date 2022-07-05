package chainclient

import (
	"context"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	"time"
)

type EthContext struct {
	chainUrl      string
	privateKey    *ecdsa.PrivateKey
	walletAddress common.Address
	client        *ethclient.Client
	chainID       *big.Int
}

func NewEthClientContext(chainUrl string, priKey *ecdsa.PrivateKey, addr common.Address) *EthContext {
	ctx := new(EthContext)
	ctx.chainUrl = chainUrl
	ctx.privateKey = priKey
	ctx.walletAddress = addr

	client, err := ethclient.Dial(chainUrl)
	if err != nil {
		log.Fatal(err)
	}
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	ctx.client = client
	ctx.chainID = chainID
	return ctx
}

func (ctx *EthContext) GetPrivateKey() *ecdsa.PrivateKey {
	return ctx.privateKey
}

func (ctx *EthContext) GetWalletAddress() common.Address {
	return ctx.walletAddress
}

func (ctx *EthContext) GetClient() *ethclient.Client {
	return ctx.client
}
func (ctx *EthContext) EstimateGas(from common.Address, to *common.Address, input []byte, gas uint64, gasPrice *big.Int) (uint64, error) {
	msg := ethereum.CallMsg{From: from, To: to, Data: input, Gas: gas, GasPrice: gasPrice}
	timeout := time.Duration(500) * time.Millisecond
	backendCtx, cancelFn := context.WithTimeout(context.Background(), timeout)
	defer cancelFn()

	estimatedGas, err := ctx.client.EstimateGas(backendCtx, msg)
	if err != nil {
		return 0, err
	}
	return estimatedGas, nil
}

func (ctx *EthContext) BuildTxOpts(value, gasLimit uint64) (*bind.TransactOpts, error) {
	nonce, err := ctx.client.PendingNonceAt(context.Background(), ctx.walletAddress)
	if err != nil {
		return nil, err
	}

	gasPrice, err := ctx.client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}

	//txOpts := bind.NewKeyedTransactor(m.Config.privateKey)
	txOpts, err := bind.NewKeyedTransactorWithChainID(ctx.privateKey, ctx.chainID)
	if err != nil {
		return nil, err
	}

	txOpts.Nonce = new(big.Int).SetUint64(nonce)
	txOpts.Value = new(big.Int).SetUint64(value) // in wei
	txOpts.GasLimit = gasLimit                   // in units
	txOpts.GasPrice = gasPrice
	return txOpts, nil
}
