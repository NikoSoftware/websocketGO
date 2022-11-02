package wsocket

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
	"runtime/debug"
	"websocketGO/model"
)

func StartWebsocket() {

	go clientManager.DealCenter()
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

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("write stop", string(debug.Stack()), r)
		}

	}()

	defer func() {
		c.Conn.Close()
	}()

	c.Conn.SetCloseHandler(func(code int, text string) error {
		fmt.Println("连接关闭码：", code) // 断开连接时将打印code和text
		return nil
	})

	for {
		_, p, err := c.Conn.ReadMessage()

		fmt.Println("收到消息", string(p))
		if err != nil {
			fmt.Println(err)
			return
		}
		msg := &model.Message{}
		err = json.Unmarshal(p, msg)
		if err != nil {
			fmt.Println(err)
			continue
		}

		if msg.Cmd == 1 {

			clientManager.Msg <- msg.Msg
		}

	}

}
