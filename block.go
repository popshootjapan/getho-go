package getho

import (
	"encoding/json"
	"fmt"
	"math/big"
	"strconv"
)

type Block struct {
	Difficulty       *big.Int         `json:"difficulty"`
	ExtraData        string           `json:"extraData"`
	GasLimit         uint64           `json:"gasLimit"`
	GasUsed          uint64           `json:"gasUsed"`
	Hash             string           `json:"hash"`
	LogsBloom        string           `json:"logsBloom"`
	Miner            string           `json:"miner"`
	MixHash          string           `json:"mixHash"`
	Nonce            *big.Int         `json:"nonce"`
	Number           *big.Int         `json:"number"`
	ParentHash       string           `json:"parentHash"`
	ReceiptsRoot     string           `json:"receiptsRoot"`
	SHA3Uncles       string           `json:"sha3Uncles"`
	Size             *big.Int         `json:"size"`
	StateRoot        string           `json:"stateRoot"`
	Timestamp        *big.Int         `json:"timestamp"`
	TotalDifficulty  *big.Int         `json:"totalDifficulty"`
	Transactions     []RawTransaction `json:"transactions"`
	TransactionsRoot string           `json:"transactionsRoot"`
	Uncles           []string         `json:"uncles"`
}

func (b *Block) String() string {
	return fmt.Sprintf("\ndifficulty: %v\nextradata: %v\ngasLimit: %v\ngasUsed: %v\nhash: %v\nlogsBloom: %v\nminer: %v\nmixHash: %v\nnonce: %v\nnumber: %v\nparentHash: %v\nreceiptsRoot: %v\nsha3Uncles: %v\nsize: %v\nstateRoot: %v\ntimestamp: %v\ntotalDifficulty: %v\ntransactions: %v\ntransactionRoot: %v\nuncles: %v\n",
		b.Difficulty.String(),
		b.ExtraData,
		b.GasLimit,
		b.GasUsed,
		b.Hash,
		b.LogsBloom,
		b.Miner,
		b.MixHash,
		b.Nonce.String(),
		b.Number.String(),
		b.ParentHash,
		b.ReceiptsRoot,
		b.SHA3Uncles,
		b.Size,
		b.StateRoot,
		b.Timestamp.String(),
		b.TotalDifficulty.String(),
		b.Transactions,
		b.TransactionsRoot,
		b.Uncles,
	)
}

func (b *Block) UnmarshalJSON(data []byte) error {
	type Alias Block
	a := &struct {
		Difficulty      *string `json:"difficulty"`
		GasLimit        *string `json:"gasLimit"`
		GasUsed         *string `json:"gasUsed"`
		Nonce           *string `json:"nonce"`
		Number          *string `json:"number"`
		Size            *string `json:"size"`
		Timestamp       *string `json:"timestamp"`
		TotalDifficulty *string `json:"totalDifficulty"`
		*Alias
	}{
		Alias: (*Alias)(b),
	}
	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	if a.Difficulty != nil {
		difficulty, _ := new(big.Int).SetString(remove0x(*a.Difficulty), 16)
		b.Difficulty = difficulty
	}

	if a.GasLimit != nil {
		gasLimit, _ := strconv.ParseUint(remove0x(*a.GasLimit), 16, 64)
		b.GasLimit = gasLimit
	}

	if a.GasUsed != nil {
		gasUsed, _ := strconv.ParseUint(remove0x(*a.GasUsed), 16, 64)
		b.GasUsed = gasUsed
	}

	if a.Nonce != nil {
		nonce, _ := new(big.Int).SetString(remove0x(*a.Nonce), 16)
		b.Nonce = nonce
	}

	if a.Number != nil {
		number, _ := new(big.Int).SetString(remove0x(*a.Number), 16)
		b.Number = number
	}

	if a.Size != nil {
		size, _ := new(big.Int).SetString(remove0x(*a.Size), 16)
		b.Size = size
	}

	if a.Timestamp != nil {
		timestamp, _ := new(big.Int).SetString(remove0x(*a.Timestamp), 16)
		b.Timestamp = timestamp
	}

	if a.TotalDifficulty != nil {
		totalDifficulty, _ := new(big.Int).SetString(remove0x(*a.TotalDifficulty), 16)
		b.TotalDifficulty = totalDifficulty
	}

	return nil
}
