package chainclient

import (
	"crypto/ecdsa"
	"github.com/PlatONnetwork/PlatON-Go/common"
)

type WalletWrapper interface {
	SetPrivateKey(privateKey *ecdsa.PrivateKey)
	GetPrivateKey() *ecdsa.PrivateKey
	GetPublicKey() *ecdsa.PublicKey
	GetAddress() common.Address
}
