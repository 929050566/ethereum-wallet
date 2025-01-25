package utils

import (
	"io"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
)

type RLPHeader types.Header

func (h *RLPHeader) EncodeRLP(w io.Writer) error {
	return types.NewBlockWithHeader((*types.Header)(h)).EncodeRLP(w)
}

func (h *RLPHeader) DecodeRLP(s *rlp.Stream) error {
	block := new(types.Block)
	err := block.DecodeRLP(s)
	if err != nil {
		return err
	}

	header := block.Header()
	*h = (RLPHeader)(*header)
	return nil
}

func (h *RLPHeader) Header() *types.Header {
	return (*types.Header)(h)
}

func (h *RLPHeader) Hash() common.Hash {
	return h.Header().Hash()
}

type Bytes []byte

func (b Bytes) Bytes() []byte {
	return b[:]
}
func (b *Bytes) SetBytes(bytes []byte) {
	*b = bytes
}

type TokenBalance struct {
	Address      common.Address `json:"address"`
	TokenAddress common.Address `json:"to_ken_address"`
	Balance      *big.Int       `json:"balance"`
	LockBalance  *big.Int       `json:"lock_balance"`
	TxType       uint8          `json:"tx_type"` // 0:充值；1:提现；2:归集；3:热转冷；4:冷转热
}
