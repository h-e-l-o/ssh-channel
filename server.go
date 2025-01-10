package main

import (
	"fmt"
	"net"
        "time"

	"golang.org/x/crypto/ssh"
)

func main() {

	hostKey := `-----BEGIN OPENSSH PRIVATE KEY-----
b3BlbnNzaC1rZXktdjEAAAAABG5vbmUAAAAEbm9uZQAAAAAAAAABAAAAMwAAAAtzc2gtZW
QyNTUxOQAAACC1p6++3ifTGI/pZvrjflrDg/WgvX2Sea0g6BJ+ibdDoAAAAJg/YSkbP2Ep
GwAAAAtzc2gtZWQyNTUxOQAAACC1p6++3ifTGI/pZvrjflrDg/WgvX2Sea0g6BJ+ibdDoA
AAAEAmNUR8Je/cLvuCRkIEl8EYr0Y/4xMezReHzKFt+oI8RrWnr77eJ9MYj+lm+uN+WsOD
9aC9fZJ5rSDoEn6Jt0OgAAAAEGRhdmlkZUBkYXZpZGUtZGEBAgMEBQ==
-----END OPENSSH PRIVATE KEY-----`

	sshConfig := ssh.ServerConfig{
		PasswordCallback: func(c ssh.ConnMetadata, pass []byte) (*ssh.Permissions, error) {
			if c.User() == "foo" && string(pass) == "bar" {
				return nil, nil
			}
			return nil, fmt.Errorf("password rejected for %q", c.User())
		},
	}

	privateBytes := []byte(hostKey)
	private, _ := ssh.ParsePrivateKey(privateBytes)
	sshConfig.AddHostKey(private)

	fmt.Println("Listening...")
	listener, _ := net.Listen("tcp", "127.0.0.1:30000")
	netConn, _ := listener.Accept()
	_, chans, _, _ := ssh.NewServerConn(netConn, &sshConfig)

	newChannel := <-chans

	c, _, _ := newChannel.Accept()

	buf := make([]byte, 4096)

	for {
		n, err := c.Read(buf)
		if err != nil {
			panic(fmt.Sprintf("Error reading from channel: %+v", err))
		}
		fmt.Printf("Got message: length %+v, msg: %+v\n", n, string(buf))
                time.Sleep(1000 * time.Millisecond)
		buf = nil

	}
}
