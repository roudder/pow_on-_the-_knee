package pow

import (
	"bytes"
	"encoding/binary"
	"log"
	"strconv"
	"strings"
)

type HashCash struct {
	//and etc. from https://ru.wikipedia.org/wiki/Hashcash data
	IP        string
	Date      string
	Counter   int
	RandomStr string
}

func NewHashCash(hcStr string) (*HashCash, error) {
	hcSplitted := strings.Split(hcStr, ":")
	//TODO impl. validation of headers (like ip should be correct, date should be is actual(not early than 2 day before))
	counter, err := strconv.Atoi(hcSplitted[2])
	if err != nil {
		return nil, err
	}
	hc := &HashCash{
		IP:        hcSplitted[0],
		Date:      hcSplitted[1],
		Counter:   counter,
		RandomStr: hcSplitted[3],
	}
	return hc, nil
}

func InitData(hc *HashCash) []byte {
	data := bytes.Join([][]byte{
		[]byte(hc.IP),
		[]byte(hc.Date),
		{byte(hc.Counter)},
		[]byte(hc.RandomStr),
		ToHex(int64(Difficulty)),
	},
		[]byte{},
	)
	return data
}

func ToHex(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}
	return buff.Bytes()
}
