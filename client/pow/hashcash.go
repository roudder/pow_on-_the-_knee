package pow

import (
	"bytes"
	"encoding/binary"
	"log"
)

type HashCash struct {
	//and etc. from https://ru.wikipedia.org/wiki/Hashcash data
	IP        string
	Date      string
	Counter   int
	RandomStr string
}

func ToHex(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}
	return buff.Bytes()
}
