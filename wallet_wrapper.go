package chainclient

import (
	"crypto/ecdsa"
	ethcommon "github.com/ethereum/go-ethereum/common"
)

type WalletWrapper interface {
	SetPrivateKey(privateKey *ecdsa.PrivateKey)
	GetPrivateKey() *ecdsa.PrivateKey
	GetPublicKey() *ecdsa.PublicKey
	GetAddress() ethcommon.Address
}
