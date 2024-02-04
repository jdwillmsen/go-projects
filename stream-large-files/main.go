package main

import (
	"bytes"
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

type FileServer struct{}

func (fs *FileServer) start() {
	ln, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go fs.readLoop(conn)
	}
}

func (fs *FileServer) readLoop(conn net.Conn) {
	defer conn.Close()

	startTime := time.Now()

	buf := new(bytes.Buffer)
	for {
		var size int64
		err := binary.Read(conn, binary.LittleEndian, &size)
		if err != nil {
			break
		}
		n, err := io.CopyN(buf, conn, size)
		if err != nil {
			log.Fatal(err)
			return
		}
		fmt.Println(buf.Bytes())
		fmt.Printf("received %d bytes over the network\n", n)
	}

	elapsedTime := time.Since(startTime)
	fmt.Printf("readLoop took %s\n", elapsedTime)
}

func sendFile(size int) error {
	startTime := time.Now()

	file := make([]byte, size)
	_, err := io.ReadFull(rand.Reader, file)
	if err != nil {
		return err
	}

	conn, err := net.Dial("tcp", ":3000")
	if err != nil {
		return err
	}

	binary.Write(conn, binary.LittleEndian, int64(size))
	n, err := io.CopyN(conn, bytes.NewReader(file), int64(size))
	if err != nil {
		return err
	}
	fmt.Printf("written %d bytes over the network\n", n)

	elapsedTime := time.Since(startTime)
	fmt.Printf("sendFile took %s\n", elapsedTime)

	return nil
}

func main() {
	go func() {
		time.Sleep(4 * time.Second)
		err := sendFile(20000000)
		if err != nil {
			log.Println(err)
			return
		}
	}()

	server := &FileServer{}
	server.start()
}
