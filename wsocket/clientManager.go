package wsocket

import (
	"fmt"
)

var (
	clientManager = NewClientManager() // 管理者

	serverIp   string
	serverPort string
)

type ClientManager struct {
	Clients map[*Client]bool //全部连接

	Users      map[string]*Client // 登陆的用户
	Msg        chan string
	Register   chan *Client // 连接连接处理
	Unregister chan *Client // 断开连接处理程序

}

func NewClientManager() (clientManager *ClientManager) {

	clientManager = &ClientManager{
		Clients:    make(map[*Client]bool, 1000),
		Users:      make(map[string]*Client, 1000),
		Msg:        make(chan string, 1000),
		Register:   make(chan *Client, 1000),
		Unregister: make(chan *Client, 1000),
	}
	return

}

func (c *ClientManager) EventRegister(client *Client) {

}

func (c *ClientManager) DealCenter() {

	for {

		select {
		case conn := <-c.Register:
			c.EventRegister(conn)
		case msg := <-c.Msg:
			fmt.Println(msg)
		}
		fmt.Println(".....")

	}

}
