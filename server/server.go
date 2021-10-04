package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

var Messages = make(chan Message)
var Clients = make(map[string]Client)

type Client struct {
	Incoming chan string
	Outgoing chan string
	Reader   *bufio.Reader
	Writer   *bufio.Writer
	Conn     net.Conn
}

func (client *Client) Read() {
	for {
		line, _ := client.Reader.ReadString('\n')
		client.Incoming <- line
	}
}

func (client *Client) Write() {
	for data := range client.Outgoing {
		client.Writer.WriteString(data)
		client.Writer.Flush()
	}
}

func (client *Client) Listen() {
	go client.Read()
	go client.Write()
}

func NewClient(connection net.Conn) *Client {
	writer := bufio.NewWriter(connection)
	reader := bufio.NewReader(connection)

	client := &Client{
		Incoming: make(chan string),
		Outgoing: make(chan string),
		Reader:   reader,
		Writer:   writer,
		Conn:     connection,
	}

	client.Listen()
	return client
}

type Message struct {
	Text   string
	Client Client
}

type Group struct {
	Clients []Client
}

func broadcast() {
	for {
		select {
		case msg := <-Messages:
			for _, conn := range Clients {
				if msg.Client.Conn.RemoteAddr().String() == conn.Conn.RemoteAddr().String() {
					continue
				}
				fmt.Fprintln(conn.Conn, fmt.Sprintf("%s: %s", conn.Conn.RemoteAddr().String(), msg.Text))
			}
		}
	}
}

func handler(conn net.Conn) {
	Clients[conn.RemoteAddr().String()] = *NewClient(conn)

	Messages <- Message{Text: "joined.", Client: Clients[conn.RemoteAddr().String()]}

	input := bufio.NewScanner(conn)
	for input.Scan() {
		Messages <- Message{Text: input.Text(), Client: Clients[conn.RemoteAddr().String()]}
	}

	Messages <- Message{Text: "has left.", Client: Clients[conn.RemoteAddr().String()]}

	//Delete client form map
	delete(Clients, conn.RemoteAddr().String())

	conn.Close() // NOTE: ignoring network errors
}

func main() {
	listen, err := net.Listen("tcp", "localhost:9090")
	if err != nil {
		log.Fatal(err)
	}

	go broadcast()

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handler(conn)
	}
}
