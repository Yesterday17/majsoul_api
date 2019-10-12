package majsoul_api

import (
	"bytes"
	"context"
	"fmt"
	"nhooyr.io/websocket"
)

const (
	socketNull     = 0
	socketNotify   = 1
	socketRequest  = 2
	socketResponse = 3
	maxSocketIndex = 60006
)

type socketClient struct {
	conn *websocket.Conn

	service  string
	index    uint16
	handlers []func(msg []byte)
}

func (s *socketClient) init(service string, url string) (err error) {
	s.service = ".lq." + service
	s.index = 0

	s.conn, _, err = websocket.Dial(context.Background(), url, nil)
	s.handlers = make([]func(msg []byte), maxSocketIndex)
	return
}

func (s *socketClient) Listen() {
	for {
		_, p, err := s.conn.Read(context.Background())
		if err != nil {
			panic(err)
			return
		}

		var index = p[1] + p[2]<<8

		var respType = p[0]
		switch respType {
		case socketNull, socketRequest:
			// TODO: error
		case socketNotify:
			fmt.Println("NOTIFY NOT IMPLEMENTED!")
			p = p[1:]
			// TODO: handle notify
		case socketResponse:
			p = p[3:]
			s.handlers[index](unWrap(p))
			s.handlers[index] = nil
		}
	}
}

func (s *socketClient) sendRPC(method string, packet []byte, handler func(msg []byte)) error {
	s.index = (s.index + 1) % maxSocketIndex

	buffer := &bytes.Buffer{}
	s.encodeHeader(buffer, socketRequest)
	s.encodePacket(buffer, method, packet)

	s.handlers[s.index] = handler
	return s.conn.Write(context.Background(), websocket.MessageBinary, buffer.Bytes())
}

func (s *socketClient) encodeHeader(bw *bytes.Buffer, t byte) {
	bw.WriteByte(t)
	bw.WriteByte(byte(255 & s.index))
	bw.WriteByte(byte(s.index >> 8))
}

func (s *socketClient) encodePacket(bw *bytes.Buffer, method string, packet []byte) {
	bw.Write(wrap(s.service+"."+method, packet))
}
