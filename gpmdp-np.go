// GPMDP-NP
//   Google Play Music Desktop Player - Now Playing
//
// a tool to export the currently playing song (and relevant metadata)
// to a file
//
// (c) 2016 foxbot
//
// Licensed under ISC

package main

import (
	"io"
	"log"

	"golang.org/x/net/websocket"
)

const Url string = "ws://localhost:5672"
const Origin string = "http://localhost"

const Version string = "1.0.0"
const ApiVersion string = "1.1.0"

func main() {
	log.Println("GPMDP-NP", Version)

	_ = loadConfig()

	ws, err := websocket.Dial(Url, "", Origin)
	if err != nil {
		log.Fatal(err)
	}

	listen(ws)
}

func listen(ws *websocket.Conn) {
	log.Println("Entering listen loop...")
	for {
		var msg Message
		err := websocket.JSON.Receive(ws, &msg)
		if err == io.EOF {
			log.Println("Received EOF, closing...")
			return
		} else if err != nil {
			log.Fatal(err)
		} else {
			handleMessage(&msg)
		}
	}
}