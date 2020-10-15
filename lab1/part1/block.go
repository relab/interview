package main

import (
	"time"
)

// Block keeps block information
type Block struct {
	Timestamp     int64          // the block creation timestamp
	Transactions  []*Transaction // The block transactions
	PrevBlockHash []byte         // the hash of the previous block
	Hash          []byte         // the hash of the block
}

// NewBlock creates and returns Block
func NewBlock(transactions []*Transaction, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), transactions, prevBlockHash, []byte{}}
	block.SetHash()
	return block
}

// NewGenesisBlock creates and returns genesis Block
func NewGenesisBlock(coinbase *Transaction) *Block {
	return NewBlock([]*Transaction{coinbase}, []byte{})
}

// SetHash calculates and sets the block hash
func (b *Block) SetHash() {
	var headers []byte
	// TODO(student)
	// For now, just use the strconv go library to converts the timestamp to a byte array. First format the int to string and than to bytes. We will improve this in future labs.
	// Reference: https://golang.org/pkg/strconv/#FormatInt
	//  - b.HashTransactions() need to be used here when combining the block header data.
	//  - You should set the block hash to be the hash of the header, so the line below should be changed.
	b.Hash = headers[:]
}

// HashTransactions returns a hash of the transactions in the block
func (b *Block) HashTransactions() []byte {
	var txHash [32]byte
	// TODO(student)
	// You should compute the hash of all transactions
	return txHash[:]
}
