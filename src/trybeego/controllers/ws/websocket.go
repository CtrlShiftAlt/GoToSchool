package ws

import (
	"fmt"
	"log"
	"net/http"

	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"
)

// WebsocketController ...
type WebsocketController struct {
	beego.Controller
}

// 连接的客户端，把每个客户端都放进来
var clients = make(map[*websocket.Conn]bool)

// 广播频道（通道）
var broadcast = make(chan Message)

// 配置升级程序（升级位websocket）
var upgrader = websocket.Upgrader{}

// Message ...
type Message struct {
	// Data ...
	Data interface{} `json:"data"`
}

// WsServer ...
func (W *WebsocketController) WsServer() {
	// 解决跨域问题（微信小程序）
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}
	// 升级将HTTP服务器连接升级到WebSocket协议
	// responseHeader包含在对客户端升级的响应中请求。
	// 使用responseHeader指定Cookie（设置Cookie）和应用程序协商的子目录（Sec WebSocket协议）
	// 如果升级失败，则升级将向客户端答复一个HTTP错误
	ws, err := upgrader.Upgrade(W.Ctx.ResponseWriter, W.Ctx.Request, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer ws.Close()
	// 将当前客户端放入map中
	clients[ws] = true
	m := Message{
		Data: 0,
	}
	// 把消息写入通道
	broadcast <- m
	W.EnableRender = false // Beego不启用渲染
	var s Message
	for {
		// 接收客户端的消息
		err := ws.ReadJSON(&s)
		if err != nil {
			log.Printf("页面可能断开啦 ws.ReadJSON error: %v", err.Error())
			delete(clients, ws) // 删除map中的客户端
			break
		} else {
			// 接受消息 业务逻辑
			fmt.Println("接受到从页面上反馈回来的消息 ", s)
		}
	}
}

func init() {
	go handleMessage()
}

func handleMessage() {
	for {
		// 读取通道中的消息
		msg := <-broadcast
		fmt.Println("clients len ", len(clients))
		// 循环map客户端
		for client := range clients {
			// 把通道中的消息发送给客户端
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("client.WriteJSON error: %v", err)
				client.Close()          // 关闭
				delete(clients, client) // 删除map中的客户端
			}
		}
	}
}
