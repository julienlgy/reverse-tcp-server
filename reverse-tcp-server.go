package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"net"
	"os"
	"strings"
	"time"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		arguments[1] = "8569"
		return
	}

	PORT := ":" + arguments[1]
	l, err := net.Listen("tcp4", PORT)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()
	rand.Seed(time.Now().Unix())

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go handleConnection(c)
	}
}

func handleConnection(c net.Conn) {
	fmt.Printf("Serving %s\n", c.RemoteAddr().String())
	for {
		reader := bufio.NewReader(os.Stdin)

		fmt.Printf("[%s] >>", c.RemoteAddr().String())
		toExec, err := reader.ReadBytes('\n')
		if err != nil {
			fmt.Println("read error")
			break
		}

		temp := strings.TrimSpace(string(toExec))
		if temp == "STOP" {
			break
		}
		fmt.Fprintf(c, "%s\n", temp)
		msg, err := bufio.NewReader(c).ReadString('°')
		if err != nil {
			fmt.Println("read client error")
			break
		}

		fmt.Printf("[%s] << \n%s\n", c.RemoteAddr().String(), msg)
	}
	//connection closed
	c.Close()
}
