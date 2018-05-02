package blockchain
import(
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)
type Block struct {
	Timestamp int64
	PrevHash []byte
	Data []byte
	Hash []byte
}

func (b* Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{b.PrevHash, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)
	b.Hash = hash[:]
}

func NewBlock(data string, prevHash []byte) *Block {
	block := &Block{ time.Now().Unix(), prevHash, []byte(data), []byte{} }
	block.SetHash()
	return block
}

func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}
