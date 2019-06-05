package ecc

import (
	"fmt"
	"github.com/codexnetwork/codexio-go/btcsuite/btcd/btcec"
//	"github.com/codexnetwork/codexio-go/btcsuite/btcutil/base58"

)

type innerK1PublicKey struct {
}

func (p *innerK1PublicKey) key(content []byte) (*btcec.PublicKey, error) {
	key, err := btcec.ParsePubKey(content, btcec.S256())
	if err != nil {
		return nil, fmt.Errorf("parsePubKey: %s", err)
	}

	return key, nil
}

func (p *innerK1PublicKey) prefix() string {
	return PublicKeyPrefixCompat
}