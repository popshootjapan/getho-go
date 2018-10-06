package getho

import (
	"encoding/json"
	"fmt"
	"math/big"
	"strconv"
)

type Transaction struct {
	From     *string  `json:"from"`
	Gas      *uint64  `json:"gas"`
	GasPrice *big.Int `json:"gasPrice"`
	To       *string  `json:"to"`
	Value    *big.Int `json:"value"`
	Data     *string  `json:"data"`
}

func (tx *Transaction) toMap() map[string]interface{} {
	txMap := make(map[string]interface{})

	if tx.From != nil {
		txMap["from"] = *tx.From
	}

	if tx.Gas != nil {
		txMap["gas"] = fmt.Sprintf("0x%s", toHexString(int(*tx.Gas)))
	}

	if tx.GasPrice != nil {
		txMap["gasPrice"] = fmt.Sprintf("0x%0x", tx.GasPrice)
	}

	if tx.To != nil {
		txMap["to"] = *tx.To
	}

	if tx.Value != nil {
		txMap["value"] = fmt.Sprintf("0x%0x", tx.Value)
	}

	if tx.Data != nil {
		if len(*tx.Data) == 0 {
			txMap["data"] = "0x"
		} else {
			txMap["data"] = *tx.Data
		}
	}

	return txMap
}

func (tx *Transaction) String() string {
	return fmt.Sprintf("from: %v\nto: %v\ngas: %v\ngasPrice: %v\nvalue: %v\ndata: %v\n",
		getString(tx.From),
		getString(tx.To),
		getUint64(tx.Gas),
		tx.GasPrice.String(),
		tx.Value.String(),
		getString(tx.Data),
	)
}

type RawTransaction struct {
	Transaction
	BlockHash        string   `json:"blockHash"`
	BlockNumber      *big.Int `json:"blockNumber"`
	Hash             string   `json:"hash"`
	Input            string   `json:"input"`
	Nonce            uint64   `json:"nonce"`
	R                *big.Int `json:"r"`
	S                *big.Int `json:"s"`
	TransactionIndex *uint64  `json:"transactionIndex"`
	V                *big.Int `json:"v"`
}

func (tx *RawTransaction) String() string {
	return fmt.Sprintf("from: %v\nto: %v\ngas: %v\ngasPrice: %v\nvalue: %v\ndata: %v\nblockHash: %v\nblockNumber: %v\nhash: %v\ninput: %v\nnounce: %v\nr: %v\ns: %v\nv: %v\ntransactionIndex: %v\n",
		getString(tx.From),
		getString(tx.To),
		getUint64(tx.Gas),
		tx.GasPrice.String(),
		tx.Value.String(),
		getString(tx.Data),
		tx.BlockHash,
		tx.BlockNumber.String(),
		tx.Hash,
		tx.Input,
		tx.Nonce,
		tx.R.String(),
		tx.S.String(),
		tx.V.String(),
		getUint64(tx.TransactionIndex),
	)
}

func (t *RawTransaction) UnmarshalJSON(data []byte) error {
	type Alias RawTransaction
	a := &struct {
		Gas              *string `json:"gas"`
		GasPrice         *string `json:"gasPrice"`
		Value            *string `json:"value"`
		BlockNumber      *string `json:"blockNumber"`
		Nonce            *string `json:"nonce"`
		R                *string `json:"r"`
		S                *string `json:"s"`
		TransactionIndex *string `json:"transactionIndex"`
		V                *string `json:"v"`
		*Alias
	}{
		Alias: (*Alias)(t),
	}
	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	if a.Gas != nil {
		gas, _ := strconv.ParseUint(remove0x(*a.Gas), 16, 64)
		t.Gas = &gas
	}

	if a.GasPrice != nil {
		gasPrice, _ := new(big.Int).SetString(remove0x(*a.GasPrice), 16)
		t.GasPrice = gasPrice
	}

	if a.Value != nil {
		value, _ := new(big.Int).SetString(remove0x(*a.Value), 16)
		t.Value = value
	}

	if a.BlockNumber != nil {
		blkNumber, _ := new(big.Int).SetString(remove0x(*a.BlockNumber), 16)
		t.BlockNumber = blkNumber
	}

	if a.Nonce != nil {
		nonce, _ := strconv.ParseUint(remove0x(*a.Nonce), 16, 64)
		t.Nonce = nonce
	}

	if a.R != nil {
		r, _ := new(big.Int).SetString(remove0x(*a.R), 16)
		t.R = r
	}

	if a.S != nil {
		s, _ := new(big.Int).SetString(remove0x(*a.S), 16)
		t.S = s
	}

	if a.V != nil {
		v, _ := new(big.Int).SetString(remove0x(*a.V), 16)
		t.V = v
	}

	if a.TransactionIndex != nil {
		txIndex, _ := strconv.ParseUint(remove0x(*a.TransactionIndex), 16, 64)
		t.TransactionIndex = &txIndex
	}

	return nil
}
