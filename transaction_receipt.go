package getho

import (
	"encoding/json"
	"math/big"
	"strconv"
)

type TransactionReceipt struct {
	BlockHash         *string     `json:"blockHash"`
	BlockNumber       *big.Int    `json:"blockNumber"`
	ContractAddress   *string     `json:"contractAddress"`
	CumulativeGasUsed *uint64     `json:"cumulativeGasUsed"`
	From              string      `json:"from"`
	GasUsed           uint64      `json:"gasUsed"`
	Logs              interface{} `json:"logs"`
	LogsBloom         string      `json:"logsBloom"`
	To                *string     `json:"to"`
	TransactionHash   string      `json:"transactionHash"`
	TransactionIndex  *uint64     `json:"transactionIndex"`
	Root              *string     `json:"root"`
	Status            string      `json:"status"`
}

func (r *TransactionReceipt) UnmarshalJSON(data []byte) error {
	type Alias TransactionReceipt
	a := &struct {
		BlockNumber       *string `json:"blockNumber"`
		CumulativeGasUsed *string `json:"cumulativeGasUsed"`
		GasUsed           *string `json:"gasUsed"`
		TransactionIndex  *string `json:"transactionIndex"`
		*Alias
	}{
		Alias: (*Alias)(r),
	}
	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	if a.BlockNumber != nil {
		blkNumber, _ := new(big.Int).SetString(remove0x(*a.BlockNumber), 16)
		r.BlockNumber = blkNumber
	}

	if a.CumulativeGasUsed != nil {
		cumulativeGasUsed, _ := strconv.ParseUint(remove0x(*a.CumulativeGasUsed), 16, 64)
		r.CumulativeGasUsed = &cumulativeGasUsed
	}

	if a.GasUsed != nil {
		gasUsed, _ := strconv.ParseUint(remove0x(*a.GasUsed), 16, 64)
		r.GasUsed = gasUsed
	}

	if a.TransactionIndex != nil {
		txIndex, _ := strconv.ParseUint(remove0x(*a.TransactionIndex), 16, 64)
		r.TransactionIndex = &txIndex
	}

	return nil
}
