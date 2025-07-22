package utils

import (
	"errors"
	"fmt"

	"github.com/ethereum/go-ethereum/crypto"
)

func RecoverAddressFromSignature(message string, signature string) (string, error) {
	msg := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(message), message)
	hash := crypto.Keccak256Hash([]byte(msg))

	sig := []byte(signature)
	if len(sig) != 65 {
		return "", errors.New("invalid signature length")
	}
	if sig[64] != 27 && sig[64] != 28 {
		sig[64] += 27
	}

	pubKey, err := crypto.SigToPub(hash.Bytes(), sig)
	if err != nil {
		return "", err
	}

	addr := crypto.PubkeyToAddress(*pubKey)
	return addr.Hex(), nil
}
