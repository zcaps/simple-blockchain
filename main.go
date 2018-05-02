package main
import (
 "fmt"
 "github.com/zcaps/simple-blockchain/blockchain"
)

func main() {
	fmt.Println("Blockchain in Go")
	bc := blockchain.NewBlockchain()
	bc.AddBlock("Send 1 BTC to User")
	bc.AddBlock("Send 3 BTC to User")

	for _, block := range bc.Blocks {
		fmt.Printf("Prev Hash: %x\n", block.PrevHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n\n", block.Hash)
	}
	
}
