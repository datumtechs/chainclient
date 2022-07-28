package chainclient

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/common"
)

type Wallet interface {
	GetPrivateKey() *ecdsa.PrivateKey
	GetPublicKey() ecdsa.PublicKey
	GetAddress() common.Address
}
