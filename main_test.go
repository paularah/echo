package main

import (
	"fmt"
	"net"
	"testing"
)

const port string = ":4000"

func TestEchoServer(t *testing.T) {
	conn, err := net.Dial("tcp", port)

	if err != nil {
		t.Errorf("got %v expected: %v", err, nil)
	}

	defer conn.Close()

	char := "someting"
	var buff [50]byte

	buffLen, err := conn.Write([]byte(char))
	_, err = conn.Read(buff[:])

	fmt.Println(string(buff[:buffLen]))

	if string(buff[:buffLen]) != char {
		t.Errorf("got %s, want %s", string(buff[:buffLen]), char)
	}

}
