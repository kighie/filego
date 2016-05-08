package protocol

import (
	"io"
	"bufio"
	"fmt"
	"bytes"
//	"strings"
//	"reflect"
//	"errors"
	"encoding/binary"
)

const (
	MSG_BEFORE_SEND	= byte(0x01)	// sender to receiver
	MSG_PREPARE		= byte(0x02)	// receiver to sender
	MSG_SEND		= byte(0x03)	// sender to receiver
	MSG_RECEIVED	= byte(0x04)	// receiver to sender
	MSG_RELAY		= byte(0x11)	// sender to receiver
	MSG_QUIT 		= byte(0x91)
	MSG_ERROR 		= byte(0x92)
	MSG_COMPLETE 	= byte(0x99)
)

type Communicator struct {
	reader *bufio.Reader
	writer *bufio.Writer
}

func NewCommunicator(r io.Reader, w io.Writer) (Communicator, error) {
	comm := Communicator { reader:bufio.NewReader(r), writer:bufio.NewWriter(w) }
	
	return comm, nil
}

type Header struct {
	Type byte
	Length	int64
	Ext	[]byte
}

func (h *Header) String() string {
	return fmt.Sprint("Header{type=",h.Type,", len=", h.Length, "ext=", h.Ext, "}")
}

// 
func (c *Communicator)ReadHeader() (Header, error) {
	var header Header
	
	err := binary.Read(c.reader, binary.BigEndian, header)
	
	fmt.Println(header)
	
	return header, err
} 


func (c *Communicator)WriteHeader(msgType byte, data ...interface{}) (Header, error) {
	var length int64
	var buffer bytes.Buffer
	
	for _,v := range data {
		binary.Write(&buffer, binary.BigEndian, v)
	}
	
	ext := buffer.Bytes();
	length = int64(buffer.Len())
	header := Header{Type:msgType, Length:length, Ext:ext}
	
	c.writer.WriteByte(msgType)
	binary.Write(c.writer, binary.BigEndian, length)
	c.writer.Write(ext)
	
	rbuf := bufio.NewReader(bytes.NewReader(ext))
	
	var pi int64
	var ps string
	
	err := binary.Read(rbuf, binary.BigEndian, &pi)
	if err != nil {
		fmt.Println("binary.Read failed:", err)
	}
	
	ps, err = rbuf.ReadString(0)
	
	if err != nil {
		fmt.Println("binary.Read ps failed:", err)
	}
	fmt.Println(pi)
	fmt.Println(ps)
	
//	return header, binary.Write(c.writer, binary.BigEndian, header)
	return header, nil
}

func (c *Communicator)Write(data []byte) (int,error) {
	return c.writer.Write(data)
}

func (c *Communicator)WriteString(data string) (int,error) {
	return c.writer.WriteString(data)
}

func (c *Communicator)WriteInt64(data int64) error {
	return binary.Write(c.writer, binary.BigEndian, data)
}





func (c *Communicator) Flush() error {
	return c.writer.Flush()
}
