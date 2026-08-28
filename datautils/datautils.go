package datautils

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"github.com/google/uuid"
)

type RawData struct {
	id        [16]byte
	email     [100]byte
	name      [50]byte
	createdAt [8]byte
	age       [1]byte
}

type Data struct {
	id        uuid.UUID
	email     string
	name      string
	createdAt uint64
	age       uint8
}

type BTree struct {
	pageSize        uint16
	minDegree       uint8
	filepath        string
	rootNodePageNum uint16
}

type Node struct {
	pageType uint8 // pagetype can be root, internal, leaf. need to find encoding
	numKeys  uint16
	data     []Data
	children []uint16
}

type PageHeader struct {
	pageType uint8
	numKeys  uint8
}
type Page []byte

func EncodeKey(d Data) ([]byte, error) {
	var s RawData
	dataBytes := make([]byte, 175)

	emailBytes := []byte(d.email)
	if len(emailBytes) > len(s.email) {
		return nil, fmt.Errorf("email too long: got %d bytes, %d max", len(emailBytes), len(s.email))
	}

	nameBytes := []byte(d.name)
	if len(nameBytes) > len(s.name) {
		return nil, fmt.Errorf("name too long: got %d bytes, %d max", len(nameBytes), len(s.name))
	}

	copy(dataBytes[0:16], d.id[:])
	copy(dataBytes[16:116], emailBytes[:])
	copy(dataBytes[116:166], nameBytes[:])
	binary.BigEndian.PutUint64(dataBytes[166:174], d.createdAt)
	dataBytes[174] = d.age

	return dataBytes, nil
}

func Decode(r []byte) (Data, error) {
	id, err := uuid.FromBytes(r[0:16])
	if err != nil {
		return Data{}, fmt.Errorf("error generating uuid from bytes: %w", err)
	}
	data := Data{
		id:        id,
		email:     string(bytes.TrimRight(r[16:116], "\x00")),
		name:      string(bytes.TrimRight(r[116:166], "\x00")),
		createdAt: binary.BigEndian.Uint64(r[166:174]),
		age:       r[174],
	}

	return data, nil
}
