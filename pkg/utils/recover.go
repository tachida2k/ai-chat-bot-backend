package utils

import (
	"encoding/hex"
	"errors"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/crypto"
)

// / RecoverAddressFromSignature recovers Ethereum address from a signed message
func RecoverAddressFromSignature(message string, signature string) (string, error) {
	signature = strings.TrimPrefix(signature, "0x")

	sig, err := hex.DecodeString(signature)
	if err != nil {
		return "", fmt.Errorf("failed to decode signature: %w", err)
	}

	if len(sig) != 65 {
		return "", errors.New("invalid signature length")
	}

	prefixedMsg := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(message), message)
	msgHash := crypto.Keccak256Hash([]byte(prefixedMsg))

	if sig[64] != 27 && sig[64] != 28 {
		return "", errors.New("invalid v value in signature")
	}
	sig[64] -= 27

	pubKeyBytes, err := crypto.Ecrecover(msgHash.Bytes(), sig)
	if err != nil {
		return "", fmt.Errorf("ecrecover failed: %w", err)
	}

	pubKey, err := crypto.UnmarshalPubkey(pubKeyBytes)
	if err != nil {
		return "", fmt.Errorf("unmarshal pubkey failed: %w", err)
	}

	address := crypto.PubkeyToAddress(*pubKey)
	return address.Hex(), nil
}
