package server

import (
	"fmt"
	"net"
)

type Message struct {
	From string
	Msg  string
}

type server struct {
	addr   string
	ln     net.Listener
	quitch chan struct{}
	Msgch  chan Message
}

func NewServer(addr string) *server {
	return &server{
		addr:   addr,
		quitch: make(chan struct{}),
		Msgch:  make(chan Message),
	}
}

func (s *server) Start() error {
	ln, err := net.Listen("tcp", s.addr)
	if err != nil {
		return err
	}
	defer ln.Close()
	s.ln = ln
	go s.Accept()
	<-s.quitch
	close(s.Msgch)
	return nil
}

func (s *server) Accept() {
	for {
		conn, err := s.ln.Accept()
		if err != nil {
			fmt.Println("accept error")
			continue
		}
		fmt.Println("connected to ", conn.RemoteAddr().String())
		go s.Read(conn)
	}
}

func (s *server) Read(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 2048)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("error in read")
			continue
		}
		// msg := buf[:n]/
		// fmt.Println(string(msg))
		s.Msgch <- Message{
			From: conn.RemoteAddr().String(),
			Msg:  string(buf[:n]),
		}
	}
}
