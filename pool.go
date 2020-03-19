package poolx

import (
	"errors"
	"gitee.com/JMArch/micro/transport"
)

var (
	//ErrClosed 连接池已经关闭Error
	ErrClosed = errors.New("pool is closed")
)

// Pool 基本方法
type Pool interface {
	Get(addr string, tr transport.Transport, opts ...transport.DialOption) (interface{}, error)

	Put(interface{}) error

	Close(interface{}) error

	Release()

	Len() int
}
