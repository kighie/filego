package common

import (
	"net"
	"fmt"
	"io"
	"errors"
	"math/rand"
	"crypto/sha1"
	"encoding/base64"
)


func MakeUid(conn net.Conn) string {
	connStr := fmt.Sprint(conn.LocalAddr() , ":" , conn.RemoteAddr(), ":", rand.Float64())
	h := sha1.New()
	io.WriteString(h, connStr)
	bytes := h.Sum(nil)
	encoded := base64.StdEncoding.EncodeToString(bytes)
	return encoded
}

func NewError(v ...interface{}) error {
	return errors.New(fmt.Sprint(v))
}