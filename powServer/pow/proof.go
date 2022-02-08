package pow

import (
	"crypto/sha256"
	"math/big"
)

const Difficulty = 6

type ProofOfWork struct {
	hc     *HashCash
	Target *big.Int
}

func NewProofOfWork(hc HashCash) *ProofOfWork {
	//TODO make target configurable and out of this block (the same story with difficulty)
	target := big.NewInt(1)
	target.Lsh(target, uint(256-Difficulty))

	proof := &ProofOfWork{&HashCash{
		IP:        hc.IP,
		Date:      hc.Date,
		Counter:   hc.Counter,
		RandomStr: hc.RandomStr,
	}, target}
	return proof
}

func (pow *ProofOfWork) Validate() bool {
	var intHash big.Int

	hc := InitData(pow.hc)
	hash := sha256.Sum256(hc)
	intHash.SetBytes(hash[:])

	return intHash.Cmp(pow.Target) == -1
}
