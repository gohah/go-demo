package main

import (
	"fmt"
	"hash/crc32"
	"math/rand"
	"work/eg1/balance"
)

type HashBalance struct {
}

func init() {
	balance.RegisterBalancer("hash", &HashBalance{})
}

func (p *HashBalance) DoBalance(insts []*balance.Instance, key ...string) (inst *balance.Instance, err error) {
	var defKey string = fmt.Sprintf("%d", rand.Int())
	if len(key) > 0 {
		defKey = key[0]
	}

	lens := len(insts)
	if lens == 0 {
		err = fmt.Errorf("No backend instance")
		return
	}

	crcTable := crc32.MakeTable(crc32.IEEE)
	hashVal := crc32.Checksum([]byte(defKey), crcTable)
	index := int(hashVal) % lens
	inst = insts[index]

	return
}
