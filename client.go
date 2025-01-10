package main

import (
	"net"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/ssh"
)

func main() {
	log.Infof("Client started")

	netConn, err := net.Dial("tcp", "127.0.0.1:30000")
	if err != nil {
		log.Fatal(err)
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

	msg := strings.Repeat("a", 50)
	for {
		log.Infof("Writing message '%+v' to channel...\n", msg)
		_, err := c.Write([]byte(msg))

		if err != nil {
			log.Fatalf("Error writing to channel: %+v\n", err)
		}
		time.Sleep(2 * time.Millisecond)   //**************************
	}

}
