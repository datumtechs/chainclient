package chainclient

import (
	"context"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
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

// EstimateGas uses context's wallet address as the caller (from)
func (ctx *EthContext) EstimateGas(to *common.Address, input []byte, gas uint64, gasPrice *big.Int) (uint64, error) {
	msg := ethereum.CallMsg{From: ctx.walletAddress, To: to, Data: input, Gas: gas, GasPrice: gasPrice}
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

func (ctx *EthContext) GetReceipt(timeoutCtx context.Context, txHash common.Hash, interval time.Duration) *types.Receipt {
	fetchReceipt := func(txHash common.Hash) (*types.Receipt, error) {
		receipt, err := ctx.client.TransactionReceipt(context.Background(), txHash)
		if nil != err {
			//including NotFound
			log.Printf("failed to get transaction receipt, txHash: %s", txHash.Hex())
			return nil, err
		} else {
			log.Printf("transaction receipt: %#v", receipt)
			return receipt, nil
		}
	}

	if interval < 0 { // do once only
		receipt, err := fetchReceipt(txHash)
		if nil != err {
			return nil
		}
		return receipt

	} else {
		ticker := time.NewTicker(interval)
		defer ticker.Stop()
		for {
			select {
			case <-timeoutCtx.Done():
				log.Printf("get transaction receipt timeout, txHash: %s", txHash.Hex())
				return nil
			case <-ticker.C:
				if receipt, err := fetchReceipt(txHash); nil == err {
					return receipt
				}
			}
		}
	}
}
