package main

import (
	"fmt"
	"io"
	"net"
)

func main() {
	//http.HandleFunc("/", handler)
	//http.ListenAndServe(":8081", nil)

	listener, err := net.Listen("tcp", ":443")
	if err != nil {
		panic("connection error:" + err.Error())
	}

	for {
		conn, err := listener.Accept()
		fmt.Println("connected")
		if err != nil {
			fmt.Println("Accept Error:", err)
			continue
		}
		copyConn(conn)
	}
}

func copyConn(src net.Conn) {
	target := src.LocalAddr()
	dst, err := net.Dial(target.Network(), target.String())
	if err != nil {
		panic("Dial Error:" + err.Error())
	}

	done := make(chan struct{})

	go func() {
		defer src.Close()
		defer dst.Close()
		io.Copy(dst, src)
		done <- struct{}{}
	}()

	go func() {
		defer src.Close()
		defer dst.Close()
		io.Copy(src, dst)
		done <- struct{}{}
	}()

	<-done
	<-done
}
