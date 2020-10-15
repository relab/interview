package main

// https://en.bitcoin.it/wiki/File:Jonny1000thetimes.png
const genesisCoinbaseData = "The Times 03/Jan/2009 Chancellor on brink of second bailout for banks"

// Blockchain keeps a sequence of Blocks
type Blockchain struct {
	blocks []*Block
}

// CreateBlockchain creates a new blockchain with genesis Block
func CreateBlockchain() *Blockchain {
	// TODO(student)
	return nil
}

// NewBlockchain creates a Blockchain
func NewBlockchain() *Blockchain {
	return CreateBlockchain()
}

// AddBlock saves the block into the blockchain
func (bc *Blockchain) AddBlock(transactions []*Transaction) *Block {
	// TODO(student)
	return nil
}

// GetGenesisBlock returns the Genesis Block
func (bc *Blockchain) GetGenesisBlock() *Block {
	// TODO(student)
	return nil
}

// CurrentBlock returns the last block
func (bc *Blockchain) CurrentBlock() *Block {
	// TODO(student)
	return nil
}

// GetBlock returns the block of a given hash
func (bc *Blockchain) GetBlock(hash []byte) (*Block, error) {
	// TODO(student)
	return nil, nil
}
