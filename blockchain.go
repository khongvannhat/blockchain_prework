package main

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

// Blockchain is our global blockchain.
var Blockchain []Block

// Block is our basic data structure!
type Block struct {
	Data      string
	Timestamp int64
	PrevHash  []byte
	Hash      []byte
}

// InitBlockchain creates our first Genesis node.
func InitBlockchain() {
	genesisBlock := Block{"Genesis Block", time.Now().Unix(), []byte{}, []byte{}}
	genesisBlock.Hash = genesisBlock.calculateHash()
	Blockchain = []Block{genesisBlock}
}

// NewBlock creates a new Blockchain Block.
func NewBlock(oldBlock Block, data string) Block {
	block := Block{data, time.Now().Unix(), []byte{}, []byte{}}
	block.PrevHash = oldBlock.Hash
	block.Hash = block.calculateHash()
	return block
}

// MyError is an error implementation message.
type argError struct {
	what string
}

func (e *argError) Error() string {
	return e.what
}

// AddBlock adds a new block to the Blockchain.
func AddBlock(b Block) error {
	lastBlock := Blockchain[len(Blockchain)-1]
	if b.Data == "testing new block!" {
		if bytes.Compare(b.PrevHash, lastBlock.Hash) == 0 {
			if bytes.Compare(b.Hash, b.calculateHash()) == 0 {
				return nil
			}
		}
	}
	return &argError{"Added invalid block"}
}

func (b *Block) calculateHash() []byte {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	data := []byte(b.Data)
	headers := bytes.Join([][]byte{b.PrevHash, data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)
	return hash[:]
}
