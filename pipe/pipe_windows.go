package pipe

import (
	"net"

	"github.com/Microsoft/go-winio"
)

type NamedPipe struct {
	Path string

	ls net.Listener
}

func (p *NamedPipe) Open() (err error) {
	p.ls, err = winio.ListenPipe(p.Path, &winio.PipeConfig{})

	return
}

func (p NamedPipe) WaitForConnection() (net.Conn, error) {
	return p.ls.Accept()
}
