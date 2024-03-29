package server

import (
	"bytes"

	"github.com/iobrother/ztimer"
	"github.com/panjf2000/gnet"
	"github.com/zchat-team/zim/app/conn/protocol"
)

type Client struct {
	Status    int
	TimerTask *ztimer.TimerTask
	DeviceId  string
	Conn      gnet.Conn
	Version   int
	Uin       string
	Platform  string
	Server    string
}

func (c *Client) Write(data []byte) error {
	return c.Conn.AsyncWrite(data)
}

func (c *Client) WritePacket(p *protocol.Packet) error {
	buf := &bytes.Buffer{}
	if err := p.Write(buf); err != nil {
		return err
	}
	return c.Conn.AsyncWrite(buf.Bytes())
}

func (c *Client) Close() {
	if c.Conn != nil {
		c.Conn.Close()
	}
}
