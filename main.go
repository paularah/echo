package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {

	if len(os.Args) != 2 {
		log.Fatal("invalid number of arguments")
	}
	port := fmt.Sprintf(":%s", os.Args[1])
	server := newServer(port)
	server.start()

}

type Server struct {
	localAddr *net.TCPAddr
}

func newServer(port string) *Server {
	address, err := net.ResolveTCPAddr("tcp", port)
	if err != nil {
		log.Fatalf("unable to resolve address for %s err: %v", port, err)
	}

	return &Server{localAddr: address}
}

func (s *Server) start() error {
	listener, err := net.ListenTCP("tcp", s.localAddr)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()
	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			return err
		}
		defer conn.Close()
		go s.handleClient(conn)
	}
}

// creates a bounded buffer, reads and writes the incoming TCP connection to the buffer
func (s *Server) handleClient(conn *net.TCPConn) {
	conn.SetKeepAlive(true)
	for {
		var buffer [300]byte

		buffLen, err := conn.Read(buffer[:])
		if err != nil {
			break
		}

		_, err = conn.Write(buffer[:buffLen])
		if err != nil {
			break
		}
	}

}
