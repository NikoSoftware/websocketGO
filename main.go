package main

import (
	"websocketGO/wsocket"
)

func main() {

	wsocket.StartWebsocket()

	manager := wsocket.NewClientManager()
	manager.DealCenter()

}
