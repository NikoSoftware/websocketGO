package wsocket

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
)

func StartWebsocket() {

	http.HandleFunc("/ws", WsEndpoint)

	http.ListenAndServe(":7300", nil)

}

func WsEndpoint(w http.ResponseWriter, r *http.Request) {
	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {

		fmt.Println(err)
	}

	client := NewClient(conn.RemoteAddr().String(), conn)

	go ReadMsg(client)

	go WriteMsg(client)

	conn.WriteMessage(1, []byte("Hi Client!"))
}

func WriteMsg(c *Client) {
	//msg := &model.Message{}

}

func ReadMsg(c *Client) {

	for {
		_, p, err := c.Conn.ReadMessage()

		fmt.Println("收到消息", string(p))
		if err != nil {

			fmt.Println(err)
		}

	}

}
