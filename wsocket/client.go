package wsocket

import "github.com/gorilla/websocket"

type Client struct {
	Addr          string          //客户地址
	Conn          *websocket.Conn //连接
	UserId        int             //id
	HeartbeatTime uint64          // 用户上次心跳时间
	LoginTime     uint64          // 登录时间 登录以后才有
}

// 创建客户端
func NewClient(addr string, conn *websocket.Conn) (client *Client) {

	client = &Client{
		Addr: addr,
		Conn: conn,
	}
	return
}
