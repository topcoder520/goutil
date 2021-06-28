package goutil

import (
	"errors"
	"io"
	"sync"
)

var ErrPoolClosed = errors.New("资源池已经被关闭。")

//Pool 一个安全的资源池，被管理的资源必须都实现io.Close接口
type Pool struct {
	m       sync.Mutex
	res     chan io.Closer
	Factory func() (io.Closer, error)
	closed  bool
}

//New 创建资源池
func New(fn func() (io.Closer, error), size uint) (*Pool, error) {
	if size <= 0 {
		return nil, errors.New("size的值太小了。")
	}
	return &Pool{
		Factory: fn,
		res:     make(chan io.Closer, size),
	}, nil
}

//Get 从资源池里获取一个资源
func (p *Pool) Get() (io.Closer, error) {
	select {
	case r, ok := <-p.res: //获取共享资源
		if !ok {
			return nil, ErrPoolClosed
		}
		return r, nil
	default:
		return p.Factory() //新生成资源
	}
}

//Close 关闭资源池，释放资源
func (p *Pool) Close() {
	p.m.Lock()
	defer p.m.Unlock()
	if p.closed {
		return
	}
	p.closed = true
	close(p.res) //关闭通道，不让写入了
	for r := range p.res {
		r.Close() //关闭通道里的资源
	}
}

//Release 释放资源
func (p *Pool) Release(r io.Closer) {
	//保证该操作和Close方法的操作是安全的
	p.m.Lock()
	defer p.m.Unlock()
	//资源池都关闭了，就省这一个没有释放的资源了，释放即可
	if p.closed {
		r.Close()
		return
	}
	select {
	case p.res <- r:
		return
	default:
		//资源池满了
		r.Close()
	}
}
