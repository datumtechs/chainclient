package chainclient

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"testing"
	"time"
)

func Test_GetLog(t *testing.T) {
	ethcontext := NewEthClientContext("ws://8.219.126.197:6790", "lat", MockWalletInstance())
	logs, _ := ethcontext.GetLog(context.Background(), common.HexToAddress("0xfd67957F61F9cC7A85da7657ED0B54b0A5867223"), new(big.Int).SetUint64(26301099))
	for _, log := range logs {
		t.Logf("log:%+v", log)
	}
}

func Test_GetLogTimeout(t *testing.T) {
	// 遍历区块日志，查询document其它数据
	timeout := time.Duration(2000) * time.Millisecond
	timeoutCtx, cancelFn := context.WithTimeout(context.Background(), timeout)
	defer cancelFn()
	ethcontext := NewEthClientContext("ws://8.219.126.197:6790", "lat", MockWalletInstance())
	prevBlock := new(big.Int).SetUint64(26301099)

	breakFor := false
	go func(ctx context.Context) {
		for prevBlock.Uint64() > 0 && !breakFor {
			logs, _ := ethcontext.GetLog(timeoutCtx, common.HexToAddress("0xfd67957F61F9cC7A85da7657ED0B54b0A5867223"), new(big.Int).SetUint64(26301099))
			for _, log := range logs {
				t.Logf("log:%+v", log)
			}
		}
	}(timeoutCtx)

	select {
	case <-timeoutCtx.Done():
		breakFor = true
	}

}
