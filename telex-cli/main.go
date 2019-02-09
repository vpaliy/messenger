package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"golang.org/x/net/websocket"
)

const channel = "Telex"

type Message struct {
	Id      string      `json:"id"`
	Channel string      `json:"channel"`
	Content interface{} `json:"content"`
}

type CreateChannelAction struct {
	Id      string `json:"id"`
	Channel string `json:"channel"`
}

type ActionRequest struct {
	Type   string      `json:"type"`
	Action interface{} `json:"action"`
}

type JoinAction struct {
	Id      string `json:"id"`
	Channel string `json:"channel"`
}

func createChannel(ws *websocket.Conn) {
	action := CreateChannelAction{"vasyl", channel}
	request := ActionRequest{"create", action}

	if err := websocket.JSON.Send(ws, &request); err != nil {
		log.Fatal("Websocket send error", err)
	}
}

func joinChannel(ws *websocket.Conn) {
	action := JoinAction{"vasyl", channel}
	request := ActionRequest{"join", action}

	if err := websocket.JSON.Send(ws, &request); err != nil {
		log.Fatal("Websocket send error", err)
	}
}

func sendMessage(ws *websocket.Conn) {
	scanner := bufio.NewScanner(os.Stdin)
	message := Message{
		"vasyl", "Telex", "Hello there!",
	}
	fmt.Println("Start typing your messages")
	for scanner.Scan() {
		message.Content = scanner.Text()
		request := ActionRequest{"send", message}

		if err := websocket.JSON.Send(ws, &request); err != nil {
			log.Fatal("Websocket send error", err)
			return
		}
		fmt.Println("Message has been sent.")
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	url := flag.String("address", "ws://localhost:8080/ws", "address of the chat server")
	origin := flag.String("origin", "http://localhost:8080/ws", "origin flag for the websocket client")
	flag.Parse()

	ws, err := websocket.Dial(*url, "", *origin)
	defer ws.Close()
	if err != nil {
		log.Fatal("websocket dial error ", err)
	}
	go func() {
		scanner := bufio.NewScanner(ws)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
		//connection is closed or there is an error, exit
		log.Println("Connection seem to be closed or error occured", scanner.Err())
		os.Exit(0)
	}()

	createChannel(ws)
	joinChannel(ws)

	sendMessage(ws)
}
