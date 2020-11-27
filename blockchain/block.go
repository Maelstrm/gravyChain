package blockchain

type BlockChain struct {
	// array of pointers to Blocks
	Blocks []*Block
}

type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
	Nonce    int
}

// Outputs a pointer to a block
func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash, 0}
	pow := NewProof(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}

func (chain *BlockChain) AddBlock(data string) {
	prevBlock := chain.Blocks[len(chain.Blocks)-1]
	new := CreateBlock(data, prevBlock.Hash)
	chain.Blocks = append(chain.Blocks, new)
}

func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

func InitBlockChain() *BlockChain {
	// Pointer to BockChain in memory with an array of blocks that makes a call to Genesis
	return &BlockChain{[]*Block{Genesis()}}
}
