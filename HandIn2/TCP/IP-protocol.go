package main

import (
	"fmt"
	"time"
)

func client(client_ch chan int, server_ch chan int) {
	seq := 0

	server_ch <- seq
	ack := <-client_ch

	fmt.Printf("Client recieved second handshake with sequence = %[2]d and acknowledgement = %[1]d", ack, seq)
}

func server(server_ch chan int, client_ch chan int) {
	ack := 1

	seq := <-server_ch
	fmt.Printf("Server recieved first handshake with sequence = %[1]d and sent acknowledgement = %[2]d", seq, ack)
	client_ch <- seq
	client_ch <- ack

}

func main() {
	client_ch := make(chan int)

	server_ch := make(chan int)

	go client(client_ch, server_ch)
	fmt.Println("Clint created.")

	go server(server_ch, client_ch)
	fmt.Println("Server created.")

	time.Sleep(10 * time.Second)
}
