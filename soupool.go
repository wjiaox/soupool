// soupool project main.go
package soupool

import (
	"fmt"
)

func test() {
	fmt.Println("Hello World!")
}

type Pool struct {
}

func Newsoupool() *Pool {
	return &Pool{}
}

func (p *Pool) Get() {

}
