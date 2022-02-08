package pow

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math"
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

func (pow *ProofOfWork) Run() int {
	var intHash big.Int
	var hash [32]byte

	counter := 0

	for counter < math.MaxInt64 {
		data := pow.InitData(pow.hc, counter)
		hash = sha256.Sum256(data)

		fmt.Printf("\r%x", hash)
		intHash.SetBytes(hash[:])

		if intHash.Cmp(pow.Target) == -1 {
			break
		} else {
			counter++
		}
	}
	fmt.Println()
	fmt.Printf("counter = %v (amount of work)", counter)
	fmt.Println()

	return counter
}

func (pow *ProofOfWork) InitData(hc *HashCash, counter int) []byte {
	data := bytes.Join([][]byte{
		[]byte(hc.IP),
		[]byte(hc.Date),
		{byte(counter)},
		[]byte(hc.RandomStr),
		ToHex(int64(Difficulty)),
	},
		[]byte{},
	)
	return data
}
