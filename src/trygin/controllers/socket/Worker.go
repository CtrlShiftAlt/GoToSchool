package socket

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var connPool = make(map[uint]*Connection)

// Ws websocket连接
func Ws(C *gin.Context) {
	var (
		uid uint
	)

	// 校验 JWT
	{
		tokenStr := C.Query("token")
		mySigningKey := []byte("AllYourBase")
		// MyClaims .
		type MyClaims struct {
			jwt.StandardClaims
			UID uint `json:"uid"`
		}
		token, err := jwt.ParseWithClaims(tokenStr, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
			return mySigningKey, nil
		})
		if err != nil {
			fmt.Printf("ParseWithClaims failed, err: %v\n", err)
			return
		}
		claims, ok := token.Claims.(*MyClaims)
		if !ok || !token.Valid {
			return
		}
		fmt.Println("ID:", claims.UID, " 正连接websocket...")
		uid = claims.UID
	}

	var (
		upgrader = websocket.Upgrader{
			//允许跨域
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		}
		wsConn *websocket.Conn
		err    error
		conn   *Connection
		data   []byte
	)

	// 提升http为socket
	wsConn, err = upgrader.Upgrade(C.Writer, C.Request, nil)
	if err != nil {
		fmt.Printf("upgrade failed,err: %v\n", err)
		return
	}
	defer wsConn.Close()

	// 初始化长连接
	if conn, err = InitConnection(wsConn); err != nil {
		conn.Close()
		return
	}

	// 接到消息 前端每个用户发起订阅，订阅连接相应的频道 （最好加上token校验用户）
	connPool[uid] = conn

	type message struct {
		Event string `json:"event"`
	}

	for {
		if data, err = conn.ReadMessage(); err != nil {
			fmt.Printf("rd %v\n", err)
			conn.Close()
			return
		}

		var msg message
		if err = json.Unmarshal(data, &msg); err != nil {
			fmt.Printf("json %v\n", err)
			conn.Close()
			return
		}

		switch msg.Event {
		case "ping":
			pong := message{
				Event: "pong",
			}
			pongJSON, _ := json.Marshal(pong)
			if err = conn.WriteMessage(pongJSON); err != nil {
				conn.Close()
				return
			}
		case "bookUser":
			// 订阅发送给xx的消息
		case "bookGroup":
			// 订阅群组消息
		case "bookGrobal":
			// 订阅全局消息
		default:
			if err = conn.WriteMessage([]byte("undefine Event: " + msg.Event)); err != nil {
				conn.Close()
				return
			}
		}

	}

}
