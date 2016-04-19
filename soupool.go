// soupool project main.go
package soupool

import (
	"sync"
	"time"
)

//Dail:func()(soupool.Conn,error){
//	c,err:=soupool.Dial("tcp","127.0.0.1:8888")
//	if err!=nil{
//		return nil,err
//	}
//	return c,nil
//}
type Pool struct {
	Dial    func() (Conn, error)
	Timeout time.Duration
	Max     uint8
	m       sync.Mutex
}

func Newsoupool() *Pool {
	return &Pool{}
}

func (p *Pool) Get() (Conn, error) {
	p.m.Lock()
	defer p.m.Unlock()
	c, err := p.Dial()
	if err != nil {
		return nil, err
	}
	c.SetTimeout(p.Timeout)
	return c, nil
}
