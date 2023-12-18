package TimeLib

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

type Blockchain struct {
	Chain []Block
}

type Block struct {
	Index     int
	Timestamp string
	Data      string
	PrevHash  string
	Hash      string
}

func CreateBlock(index int, data string, prevHash string) Block {
	block := Block{
		Index:     index,
		Timestamp: time.Now().String(),
		Data:      data,
		PrevHash:  prevHash,
	}
	block.Hash = CalculateHash(block)
	return block
}

func CalculateHash(block Block) string {
	record := string(block.Index) + block.Timestamp + block.Data + block.PrevHash
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.Chain[len(bc.Chain)-1]
	newBlock := CreateBlock(prevBlock.Index+1, data, prevBlock.Hash)
	bc.Chain = append(bc.Chain, newBlock)
}

func Running() {
	genesisBlock := CreateBlock(0, "Genesis Block", "")
	blockchain := Blockchain{Chain: []Block{genesisBlock}}

	fmt.Println("Genesis Block Oluşturuldu : ")
	fmt.Printf("Index: %d\nTimestamp: %s\nData: %s\nPrevHash: %s\nHash: %s\n",
		genesisBlock.Index, genesisBlock.Timestamp, genesisBlock.Data, genesisBlock.PrevHash, genesisBlock.Hash)

	fmt.Println("\nBir sonraki bloğu eklemek için veri girin:")
	var data string
	fmt.Scanln(&data)
	blockchain.AddBlock(data)

	newBlock := blockchain.Chain[len(blockchain.Chain)-1]
	fmt.Printf("\nYeni Block Oluşturuldu:\nIndex: %d\nTimestamp: %s\nData: %s\nPrevHash: %s\nHash: %s\n", newBlock.Index, newBlock.Timestamp, newBlock.Data, newBlock.PrevHash, newBlock.Hash)

}
