package chainclient

import (
	"context"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"math/big"
	"testing"
)

func Test_GetLog(t *testing.T) {
	ethcontext := NewPlatonClientContext("ws://8.219.126.197:6790", "lat", MockWalletInstance())
	logs := ethcontext.GetLog(context.Background(), common.HexToAddress("0xfd67957F61F9cC7A85da7657ED0B54b0A5867223"), new(big.Int).SetUint64(26301099))
	t.Log(logs)
}
