package gochain

import (
	"bytes"
	"crypto/sha256"
	"math"
	"math/big"
)

var DIFFICULTY = int64(18) // Default Difficulty 0x000___

type Pow struct {
	Block  *Block
	Target *big.Int
}

func New_proof(_block *Block) *Pow {
	_block.Difficulty = DIFFICULTY
	_target := big.NewInt(1)
	_target.Lsh(_target, uint(256-_block.Difficulty))
	pow := &Pow{_block, _target}
	return pow
}

func (pow *Pow) Generate_hash(nonce int64) []byte {
	digest := bytes.Join(
		[][]byte{
			big.NewInt(pow.Block.Block_number).Bytes(),
			pow.Block.Prev_hash,
			pow.Block.Data,
			big.NewInt(nonce).Bytes(),
			big.NewInt(pow.Block.Difficulty).Bytes(),
		},
		[]byte{},
	)
	_hash := sha256.Sum256(digest)
	return _hash[:]
}

func (pow *Pow) Run() {
	var _hash_in_bigint big.Int
	var nonce int64

	for nonce < math.MaxInt64 {
		_hash := pow.Generate_hash(nonce)

		_hash_in_bigint.SetBytes(_hash)
		if _hash_in_bigint.Cmp(pow.Target) == -1 {
			pow.Block.Nonce = nonce
			pow.Block.Hash = _hash
			return
		} else {
			nonce++
		}
	}
}

func (pow *Pow) Validate() bool {
	var _hash_bigint big.Int
	_hash := pow.Generate_hash(pow.Block.Nonce)
	_hash_bigint.SetBytes(_hash)
	return _hash_bigint.Cmp(pow.Target) == -1
}
