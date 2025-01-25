package address

import (
	"encoding/hex"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

type EthAddress struct {
	Private string `json:"private"`
	Public  string `json:"public"`
	Address string `json:"address"`
}

func CreateEthAddress() (*EthAddress, error) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		return nil, err
	}

	address := &EthAddress{
		Private: hex.EncodeToString(crypto.FromECDSA(privateKey)),
		Public:  hex.EncodeToString(crypto.FromECDSAPub(&privateKey.PublicKey)),
		Address: crypto.PubkeyToAddress(privateKey.PublicKey).String(),
	}
	return address, nil
}

func PublicKeytoAddress(publicKey string) (string, error) {
	pub, err := hex.DecodeString(publicKey)
	if err != nil {
		return "", err
	}
	address := common.BytesToAddress(crypto.Keccak256(pub[1:])[:12]).String()
	return address, nil
}
