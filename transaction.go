package getho

import "fmt"

type Transaction struct {
	From     *string `json:"from"`
	Gas      *string `json:"gas"`
	GasPrice *string `json:"gasPrice"`
	To       *string `json:"to"`
	Value    *string `json:"value"`
	Data     *string `json:"data"`
}

func (tx *Transaction) toMap() map[string]interface{} {
	txMap := make(map[string]interface{})

	if tx.From != nil {
		txMap["from"] = tx.From
	}

	if tx.Gas != nil {
		txMap["gas"] = tx.Gas
	}

	if tx.GasPrice != nil {
		txMap["gasPrice"] = tx.GasPrice
	}

	if tx.To != nil {
		txMap["to"] = *tx.To
	}

	if tx.Value != nil {
		txMap["value"] = tx.Value
	}

	if tx.Data != nil {
		txMap["data"] = tx.Data
	}

	return txMap
}

func (tx *Transaction) String() string {
	return fmt.Sprintf("from: %v\nto: %v\ngas: %v\ngasPrice: %v\nvalue: %v\ndata: %v\n", tx.From, tx.To, tx.Gas, tx.GasPrice, tx.Value, tx.Data)
}

type RawTransaction struct {
	Transaction
	BlockHash        string  `json:"blockHash"`
	BlockNumber      *string `json:"blockNumber"`
	Hash             string  `json:"hash"`
	Input            string  `json:"input"`
	Nonce            string  `json:"nonce"`
	R                string  `json:"r"`
	S                string  `json:"s"`
	TransactionIndex string  `json:"transactionIndex"`
	V                string  `json:"v"`
}

func (tx *RawTransaction) String() string {
	return fmt.Sprintf("from: %v\nto: %v\ngas: %v\ngasPrice: %v\nvalue: %v\ndata: %v\nblockHash: %v\nblockNumber: %v\nhash: %v\ninput: %v\nnounce: %v\nr: %v\ns: %v\nv: %v\ntransactionIndex: %v\n",
		tx.From,
		tx.To,
		tx.Gas,
		tx.GasPrice,
		tx.Value,
		tx.Data,
		tx.BlockHash,
		tx.BlockNumber,
		tx.Hash,
		tx.Input,
		tx.Nonce,
		tx.R,
		tx.S,
		tx.V,
		tx.TransactionIndex,
	)
}
