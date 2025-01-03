package gochain

type Blockchain struct {
	Chain []*Block
}

type Block struct {
	Block_number int64
	Prev_hash    []byte
	Hash         []byte
	Data         []byte
	Nonce        int64
	Difficulty   int64
}

// func (b *Block) Derive_hash() {
// 	digest := bytes.Join([][]byte{b.Data, b.Prev_hash}, []byte{})
// 	hash := sha256.Sum256(digest)
// 	b.Hash = hash[:]
// }

func create_block(_block_number int64, _prev_hash []byte, _data []byte) *Block {
	block := &Block{_block_number, _prev_hash, []byte{}, _data, 0, 0}
	pow := New_proof(block)
	pow.Run()
	return block
}

func (c *Blockchain) Add_block(_data string) {
	_prev_hash := c.Chain[len(c.Chain)-1].Hash
	b := create_block(int64(len(c.Chain)), _prev_hash, []byte(_data))
	c.Chain = append(c.Chain, b)
}

func generate_genesis() *Block {
	genesis_data := []byte("GENESIS")
	b := create_block(0, []byte{}, genesis_data)
	pow := New_proof(b)
	pow.Run()
	return b
}

func Init_chain() *Blockchain {
	return &Blockchain{
		[]*Block{generate_genesis()},
	}
}
