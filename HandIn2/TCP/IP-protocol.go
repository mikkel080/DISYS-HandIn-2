package main

import (
	"fmt"
	"time"
)

func client(client_seq_ch chan int, client_ack_ch chan int, server_seq_ch chan int, server_ack_ch chan int) {
	seq := 0
	ack := 0

	server_seq_ch <- seq

	ack_temp := <-client_ack_ch
	if ack_temp == seq+1 {
		seq_temp := <-client_seq_ch
		if seq_temp == seq {
			seq = seq_temp
			ack = ack_temp
		}
	}

	fmt.Printf("Client recieved second handshake with sequence = %[2]d and acknowledgement = %[1]d", ack, seq)
}

func server(server_seq_ch chan int, server_ack_ch chan int, client_seq_ch chan int, client_ack_ch chan int) {
	ack := 1

	seq := <-server_seq_ch
	fmt.Printf("Server recieved first handshake with sequence = %[1]d and sent acknowledgement = %[2]d", seq, ack)
	client_seq_ch <- seq
	client_ack_ch <- ack

}

func main() {
	client_seq_ch := make(chan int)
	client_ack_ch := make(chan int)

	server_seq_ch := make(chan int)
	server_ack_ch := make(chan int)

	go client(client_seq_ch, client_ack_ch, server_seq_ch, server_ack_ch)
	fmt.Println("Clint created.")

	go server(server_seq_ch, server_ack_ch, client_seq_ch, client_ack_ch)
	fmt.Println("Server created.")

	time.Sleep(10 * time.Second)
}
