package majsoul_api

import (
	"bytes"
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

const (
	SocketNull     = 0
	SocketNotify   = 1
	SocketRequest  = 2
	SocketResponse = 3
	MaxSocketIndex = 60006
)

type socket struct {
	dialer *websocket.Dialer
	conn   *websocket.Conn

	service  string
	index    uint16
	handlers []func(msg []byte)
}

func (s *socket) init(service string) (err error) {
	s.service = ".lq." + service
	s.index = 0
	s.dialer = &websocket.Dialer{
		Proxy:            http.ProxyFromEnvironment,
		HandshakeTimeout: 45 * time.Second,
	}
	s.conn, _, err = s.dialer.Dial("wss://mj-srv-5.majsoul.com:4101", http.Header{})
	s.handlers = make([]func(msg []byte), MaxSocketIndex)
	return
}

func (s *socket) Listen() {
	for {
		_, p, err := s.conn.ReadMessage()
		if err != nil {
			panic(err)
			return
		}

		var index = p[1] + p[2]<<8

		var respType = p[0]
		switch respType {
		case SocketNull, SocketRequest:
			// TODO: error
		case SocketNotify:
			fmt.Println("NOTIFY NOT IMPLEMENTED!")
			p = p[1:]
			// TODO: handle notify
		case SocketResponse:
			p = p[3:]
			s.handlers[index](unWrap(p))
			s.handlers[index] = nil
		}
	}
}

func (s *socket) sendRPC(method string, packet []byte, handler func(msg []byte)) error {
	s.index = (s.index + 1) % MaxSocketIndex

	buffer := &bytes.Buffer{}
	s.encodeHeader(buffer, SocketRequest)
	s.encodePacket(buffer, method, packet)

	s.handlers[s.index] = handler
	return s.conn.WriteMessage(websocket.BinaryMessage, buffer.Bytes())
}

func (s *socket) encodeHeader(bw *bytes.Buffer, t byte) {
	bw.WriteByte(t)
	bw.WriteByte(byte(255 & s.index))
	bw.WriteByte(byte(s.index >> 8))
}

func (s *socket) encodePacket(bw *bytes.Buffer, method string, packet []byte) {
	bw.Write(wrap(s.service+"."+method, packet))
}
