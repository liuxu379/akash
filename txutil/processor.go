package txutil

import (
	"fmt"

	"github.com/gogo/protobuf/proto"
	"github.com/ovrclk/photon/types"
	crypto "github.com/tendermint/go-crypto"
)

type TxProcessor interface {
	Validate() error
	GetTx() *types.Tx
}

func ProcessTx(buf []byte) (*types.Tx, error) {
	txp, err := NewTxProcessor(buf)
	if err != nil {
		return nil, err
	}
	if err := txp.Validate(); err != nil {
		return nil, err
	}
	return txp.GetTx(), nil
}

func NewTxProcessor(buf []byte) (TxProcessor, error) {
	tx := new(types.Tx)
	if err := tx.Unmarshal(buf); err != nil {
		return nil, err
	}
	return &txProcessor{tx}, nil
}

type txProcessor struct {
	tx *types.Tx
}

func (txp *txProcessor) Validate() error {
	if txp.tx.Key == nil {
		return fmt.Errorf("missing key")
	}
	if txp.tx.Signature == nil {
		return fmt.Errorf("missing signature")
	}

	pbytes, err := proto.Marshal(&txp.tx.Payload)
	if err != nil {
		return err
	}

	if !txp.tx.Key.VerifyBytes(pbytes, crypto.Signature(*txp.tx.Signature)) {
		return fmt.Errorf("invalid signature")
	}

	return nil
}

func (txp *txProcessor) GetTx() *types.Tx {
	return txp.tx
}