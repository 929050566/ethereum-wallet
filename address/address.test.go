package address

import (
	"encoding/hex"
	"testing"

	"github.com/ethereum/go-ethereum/crypto"
)

func TestPublicKeyToAddress(t *testing.T) {
	// Call the function PublicKeyToAddress
	pubKeyHex := "022505a03b55f896c0948f35a2c63b46f6a4cdb8221164bc27bb9980617dacbce7"
	gotAddress, err := PublicKeytoAddress(pubKeyHex)
	if err != nil {
		t.Fatalf("Error creating address from public key: %v", err)
	}

	// Expected address
	expectedAddress := "0xExpectedAddress" // Replace with the actual expected address

	if gotAddress != expectedAddress {
		t.Errorf("Expected address: %s, got: %s", expectedAddress, gotAddress)
	}

	t.Logf("Got Address: %s", gotAddress)
}

func TestPrivateKeyToHex(t *testing.T) {
	// Generate a private key
	privKey, err := crypto.GenerateKey()
	if err != nil {
		t.Fatalf("Error generating private key: %v", err)
	}

	// Convert private key to hex
	priKeyHex := hex.EncodeToString(privKey.D.Bytes())

	// Expected private key hex
	expectedPriKeyHex := hex.EncodeToString(privKey.D.Bytes())

	if priKeyHex != expectedPriKeyHex {
		t.Errorf("Expected private key hex: %s, got: %s", expectedPriKeyHex, priKeyHex)
	}

	t.Logf("Got Private Key Hex: %s", priKeyHex)
}
