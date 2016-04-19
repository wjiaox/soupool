package soupool

import (
	"bufio"
	"net"
	"sync"
	"time"
)

type Conn interface {
	Write(buf []byte) (int, error)
	Read(buf []byte) (int, error)
	SetTimeout(timeout time.Duration)
	Close()
}

type conn struct {
	m    sync.Mutex
	conn net.Conn
	bw   *bufio.Writer
	br   *bufio.Reader
}

func Dial(network, address string) (Conn, error) {
	c, err := net.Dial(network, address)
	if err != nil {
		return nil, err
	}
	return &conn{
		conn: c,
		bw:   bufio.NewWriter(c),
		br:   bufio.NewReader(c),
	}, nil
}

func (c *conn) SetTimeout(timeout time.Duration) {
	if timeout == 0 {
		//default 5s
		timeout = 5
	}
	c.conn.SetReadDeadline(time.Now().Add(timeout))
	c.conn.SetWriteDeadline(time.Now().Add(timeout))
}

func (c *conn) Write(buf []byte) (int, error) {
	//if c.conn != nil {
	return c.bw.Write(buf)

}
func (c *conn) Read(buf []byte) (int, error) {
	//if c.conn != nil {
	return c.br.Read(buf)

}

func (c *conn) Close() {
	c.m.Lock()
	c.conn.Close()
	c.conn = nil
	c.m.Unlock()
}
