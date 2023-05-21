package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	fmt.Println("Starting daemon...")

	sockfile := "/tmp/echo.sock"

	socket, err := net.Listen("unix", sockfile)
	if err != nil {
		log.Fatal(err)
	}

	// Cleanup the sockfile.
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		os.Remove(sockfile)
		os.Exit(1)
	}()

	for {
		// Accept an incoming connection.
		conn, err := socket.Accept()
		if err != nil {
			log.Fatal(err)
		}

		// Handle the connection in a separate goroutine.
		go func(conn net.Conn) {
			defer conn.Close()
			// Create a buffer for incoming data.
			buf := make([]byte, 4096)

			// Read data from the connection.
			n, err := conn.Read(buf)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Printf("Received input: %s\n", string(buf[:n]))

			// reply
			reply := fmt.Sprintf("Daemon time is %s", time.Now().Format(time.RFC3339))
			fmt.Printf("Replying with '%s' ...\n", reply)
			_, err = conn.Write([]byte(reply))
			if err != nil {
				log.Fatal(err)
			}
		}(conn)
	}

	fmt.Println("Daemon exited")
}
