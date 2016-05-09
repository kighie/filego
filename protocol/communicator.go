package protocol

import (
//	"io"
	"bufio"
	"fmt"
	"net"
//	"bytes"
//	"strings"
//	"log"
	"errors"
//	"encoding/binary"
	"github.com/kighie/filego/common/convert"
)

const (
	MSG_START		= byte(0x01)	// sender to receiver : start of heading
	MSG_TEXT_DEL	= byte(0x1E)	// sender to receiver : record separator
	
	MSG_BEFORE_SEND	= byte(0xA1)	// sender to receiver
	MSG_PREPARE		= byte(0xA2)	// receiver to sender
	MSG_SEND		= byte(0xA3)	// sender to receiver
	MSG_RECEIVED	= byte(0xA4)	// receiver to sender
	MSG_RELAY		= byte(0xB1)	// sender to receiver
	MSG_QUIT 		= byte(0xF1)
	MSG_ERROR 		= byte(0xF2)
	MSG_COMPLETE 	= byte(0xF9)
)

const MSG_LENGTH_UNDEFINED = -1


type Communicator struct {
	reader *bufio.Reader
	writer *bufio.Writer
	conn net.Conn
}

func NewCommunicator(c net.Conn) (Communicator, error) {
	comm := Communicator { reader:bufio.NewReader(c), writer:bufio.NewWriter(c), conn:c }
	
	return comm, nil
}

//func NewCommunicator(r io.Reader, w io.Writer) (Communicator, error) {
//	comm := Communicator { reader:bufio.NewReader(r), writer:bufio.NewWriter(w) }
//	
//	return comm, nil
//}

type Header struct {
	Type 		byte
	Length		int64
	SessionId 	string
}

func (h Header) String() string {
	return fmt.Sprint("HEADER {type=",h.Type,", len=", h.Length, ", session=", h.SessionId, "}")
}


func (c Communicator)GetWriter() *bufio.Writer {
	return c.writer
}

func (c Communicator)GetReader() *bufio.Reader {
	return c.reader
}


func (c *Communicator)WriteHeader(msgType byte, length int64, sessionId string) (Header, error) {
	header := Header{Type:msgType, Length:length, SessionId:sessionId}
	
	c.writer.WriteByte(MSG_START)
	c.writer.WriteByte(msgType)
	c.writer.Write( convert.Int64ToBytes(length) )
	c.writer.WriteString(sessionId)
	c.writer.WriteByte(MSG_TEXT_DEL)
	
	return header, nil
}

// 
func (c *Communicator)ReadHeader() (Header, error) {
	var header Header
	var b byte
	var err error

	if b, err = c.reader.ReadByte() ; err != nil {
		return header, err
	}
	
	if b != MSG_START {
		n := c.reader.Buffered()
		c.reader.Discard(n)
		return header, errors.New("No Start Flag[" + string(b) + "][]") 
	}
	
	if header.Type, err = c.reader.ReadByte() ; err != nil {
		return header, err
	}
	if header.Length, err = convert.ReadInt64(c.reader) ; err != nil {
		return header, err
	}
	if header.SessionId, err = c.reader.ReadString(MSG_TEXT_DEL) ; err != nil {
		return header, err
	}
	
	return header, err
} 


func (c *Communicator)Write(data []byte) (int,error) {
	return c.writer.Write(data)
}


func (c *Communicator)WriteInt64(data int64) error {
	_,err := c.writer.Write(convert.Int64ToBytes(data))
	
	return err
}

func (c *Communicator)WriteInt16(data int16) error {
	_,err := c.writer.Write(convert.Int16ToBytes(data))
	
	return err
}

func (c *Communicator) Flush() error {
	return c.writer.Flush()
}


func (c *Communicator)ReadInt64() (int64, error) {
	return convert.ReadInt64(c.reader)
}

func (c *Communicator)ReadInt16() (int16, error) {
	return convert.ReadInt16(c.reader)
}

/**
 * wrtire string followed by MSG_TEXT_DEL flag
 */
func (c *Communicator)WriteTextField(data string) (int,error) {
	n, err := c.writer.WriteString(data)
	if err != nil {
		return n, err
	}
	err = c.writer.WriteByte(MSG_TEXT_DEL)
	return n+1, err
}

func (c *Communicator)ReadTextField() (string, error) {
	text, err := c.reader.ReadString(MSG_TEXT_DEL)

	if err == nil {
		text = text[:len(text)-1]
	}
		
	return text, err
}
