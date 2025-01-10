package main

import (
	"net"
        "fmt"
	//"strings"
	"time"

	"golang.org/x/crypto/ssh"
)

func main() {
	fmt.Println("Client started")

	netConn, err := net.Dial("tcp", "127.0.0.1:30000")
	if err != nil {
		panic(err)
	}

	sshConfig := ssh.ClientConfig{
		User: "foo",
		Auth: []ssh.AuthMethod{
			ssh.Password("bar"),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	sshConn, _, _, _ := ssh.NewClientConn(netConn, "127.0.0.1:30000", &sshConfig)

	c, _, _ := sshConn.OpenChannel("control", []byte{})

	//msg := strings.Repeat("a", 50)
        count := 0
	for {
	        msg := fmt.Sprintf("%d a", count)
		fmt.Printf("Writing message '%+v' to channel...\n", msg)
		n, err := c.Write([]byte(msg))
                count += 1

		if err != nil {
			panic(fmt.Sprintf("Error writing to channel: %+v", err))
		}
                if n != len(msg) {
                        panic(fmt.Sprintf("Wrote %d bytes instead of %d\n", n, len(msg)))
                }
		time.Sleep(300 * time.Millisecond)   //**************************
	}

}
