package convert

import (
	"io"
//	"bytes"
	"encoding/binary"
)


const (
	INT8 	= 1
	INT16 	= 2
	INT32 	= 4
	INT64	= 8
)


func WriteInt64(w io.Writer, i int64) error  {
	bs := Int64ToBytes(i)
	_,err := w.Write(bs)
	return err
}


func ReadInt64(r io.Reader) (int64, error) {
	var b = make([]byte, INT64)
	_,err := r.Read(b)
	if err != nil {
		return 0, err
	}
	return BytesToInt64(b), nil
}

func Int64ToBytes(i int64) []byte {
	var bs = make([]byte, INT64)
	binary.BigEndian.PutUint64(bs,uint64(i))
	
	return bs
}

func BytesToInt64(bs []byte) int64 {
	return int64(binary.BigEndian.Uint64(bs))
}


func WriteInt16(w io.Writer, i int16) error  {
	bs := Int16ToBytes(i)
	_,err := w.Write(bs)
	return err
}


func ReadInt16(r io.Reader) (int16, error) {
	var b = make([]byte, INT16)
	_,err := r.Read(b)
	if err != nil {
		return 0, err
	}
	return BytesToInt16(b), nil
}

func Int16ToBytes(i int16) []byte {
	var bs = make([]byte, INT16)
	binary.BigEndian.PutUint16(bs,uint16(i))
	
	return bs
}

func BytesToInt16(bs []byte) int16 {
	return int16(binary.BigEndian.Uint16(bs))
}

func ReadByte(r io.Reader) (byte, error) {
	var b = make([]byte, 1)
	_,err := r.Read(b)
	if err != nil {
		return 0, err
	}
	return b[0], nil
}

