package majsoul_api

import (
	"bytes"
	"context"
	"github.com/Yesterday17/majsoul_api/lq"
	"log"
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

	service   string
	index     uint16
	handlers  []func(msg []byte)
	listeners map[string][]func(wrapper lq.Wrapper)
}

func (s *socketClient) init(service string, url string) (err error) {
	s.service = ".lq." + service
	s.index = 0

	s.conn, _, err = websocket.Dial(context.Background(), url, nil)
	s.handlers = make([]func(msg []byte), maxSocketIndex)
	s.listeners = make(map[string][]func(wrapper lq.Wrapper))
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
			log.Fatalf("Unexpected socket type: %d", respType)
		case socketNotify:
			p = p[1:]
			data := unWrap(p)
			for _, listener := range s.listeners[data.Name] {
				listener(data)
			}
		case socketResponse:
			p = p[3:]
			s.handlers[index](unWrap(p).Data)
			s.handlers[index] = nil
		}
	}
}

// TODO: make it private
func (s *socketClient) AddListener(method string, listener func(wrapper lq.Wrapper)) {
	arr, ok := s.listeners[".lq."+method]
	if !ok {
		s.listeners[".lq."+method] = make([]func(wrapper lq.Wrapper), 8)
		arr = s.listeners[method]
	}
	s.listeners[method] = append(arr, listener)
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
