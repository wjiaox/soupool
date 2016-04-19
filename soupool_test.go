package soupool

import "testing"

func Test4Conn(t *testing.T) {
	pool := &Pool{
		Dial: func() (Conn, error) {
			c, err := Dial("tcp", "127.0.0.1:12345")
			if err != nil {
				return nil, err
			}
			return c, nil
		},
	}
	c, err := pool.Get()
	if err != nil {
		t.Log("pool get():", err.Error())
	} else {
		c.Write([]byte("test"))
		//buf := make([]byte, 1024)
		//c.Read(buf)
		//t.Log("rcv:", string(buf))
	}
	c.Close()

}
