package main

import (
	"fmt"
	"time"
)

func client(client_seq_ch chan int, client_ack_ch chan int, server_seq_ch chan int, server_ack_ch chan int) {
	seq := 0
	ack := 0

	//Send first handshake
	server_seq_ch <- seq
	fmt.Printf("Client sent first handshake.\n")

	//Reviece second handshake
	ack_temp := <-client_ack_ch
	if ack_temp == seq+1 {
		seq_temp := <-client_seq_ch
		if seq_temp == seq {
			seq = seq_temp
			ack = ack_temp
			fmt.Printf("Client recieved second handshake with sequence = %[2]d and acknowledgement = %[1]d\n", ack, seq)
			seq++
		}
	}

	//Send third handshake
	fmt.Printf("Client sent third handshake.\n")
	server_seq_ch <- seq
	server_ack_ch <- ack

	//Send packages
	for i := 0; i < 10; i++ {
		ack_temp = <-client_ack_ch
		if ack_temp == ack+1 {
			ack = ack_temp
			fmt.Printf("Client recieved acknowledgement = %d.\n", ack)
			seq++
			fmt.Printf("Client sent fictional data with sequence = %d.\n", seq)
			server_seq_ch <- seq
		}
	}

}

func server(server_seq_ch chan int, server_ack_ch chan int, client_seq_ch chan int, client_ack_ch chan int) {
	ack := 1

	//Recieve first handshake
	seq := <-server_seq_ch
	time.Sleep(500)
	fmt.Printf("Server recieved first handshake with sequence = %[1]d and sent acknowledgement = %[2]d\n", seq, ack)

	//Send second handshake
	client_ack_ch <- ack
	time.Sleep(500)
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
			ack++
			client_ack_ch <- ack
		}
	}

	//Recieve packages
	for i := 0; i < 10; i++ {
		seq_temp = <-server_seq_ch
		if seq_temp == seq+1 {
			seq = seq_temp
			ack++
			fmt.Printf("Server recieved sequence = %[1]d and sent acknowledgement = %[2]d\n", seq, ack)
			client_ack_ch <- ack
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
