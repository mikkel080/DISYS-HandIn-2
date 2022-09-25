package main

import (
	"fmt"
	"time"
)

func client(client_seq_ch chan int, client_ack_ch chan int, server_seq_ch chan int, server_ack_ch chan int) {
	seq := 0
	ack := 0

	//Send first handshake
	fmt.Printf("Client sent first handshake.\n")
	server_seq_ch <- seq

	//Reviece second handshake
	ack_temp := <-client_ack_ch
	if ack_temp == seq+1 {
		seq_temp := <-client_seq_ch
		if seq_temp == seq {
			seq = seq_temp
			ack = ack_temp
			fmt.Printf("Client recieved second handshake with sequence = %[2]d and acknowledgement = %[1]d\n", ack, seq)
			seq = seq + 1
		}
	}

	//Send third handshake
	fmt.Printf("Client sent third handshake.\n")
	server_seq_ch <- seq
	server_ack_ch <- ack
}

func server(server_seq_ch chan int, server_ack_ch chan int, client_seq_ch chan int, client_ack_ch chan int) {
	ack := 1

	//Recieve first handshake
	seq := <-server_seq_ch
	fmt.Printf("Server recieved first handshake with sequence = %[1]d and sent acknowledgement = %[2]d\n", seq, ack)

	//Send second handshake
	client_ack_ch <- ack
	time.Sleep(1 * time.Second)
	client_seq_ch <- seq
	fmt.Printf("Server sent second handshake.\n")

	//Recieve third handshake
	seq_temp := <-server_seq_ch
	if seq_temp == seq+1 {
		ack_temp := <-server_ack_ch
		if ack_temp == ack {
			seq = seq_temp
			ack = ack_temp
			fmt.Printf("Server recieved third handshake with sequence = %[1]d and acknowledgement = %[2]d\n", seq, ack)
			ack = ack + 1
		}
	}
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
