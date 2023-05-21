package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

func reader(r io.Reader) {
	buf := make([]byte, 1024)
	for {
		n, err := r.Read(buf[:])
		if err != nil {
			return
		}
		fmt.Println(string(buf[0:n]))
	}
}

func main() {
	fmt.Println("Starting CLI...")

	sockfile := "/tmp/echo.sock"

	c, err := net.Dial("unix", sockfile)
	if err != nil {
		panic(err)
	}
	defer c.Close()
	go reader(c)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Start typing")

	for {
		fmt.Print("> ")
		text, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}

		// convert CRLF to LF
		text = strings.Replace(text, "\n", "", -1)

		_, err = c.Write([]byte(text))
		if err != nil {
			log.Fatal("write error:", err) // FIXME: errors on the second input
			panic(err)
		}

		time.Sleep(100 * time.Millisecond)
	}
}
