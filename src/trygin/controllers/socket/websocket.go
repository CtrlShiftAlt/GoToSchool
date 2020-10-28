package socket

import (
	"errors"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

// Connection .
type Connection struct {
	wsConn           *websocket.Conn
	inChan           chan []byte // 读取websocket的消息的channel
	outChan          chan []byte // 给websocket写消息的channel
	closeChan        chan byte   // 关闭通道信号
	mutex            sync.Mutex
	isClosed         bool          // closeChan状态
	heartbeatTimeOut time.Duration // 超时时间
	heartbeatTime    time.Time     // 消息时间
}

// Close 关闭websocket并设置状态
func (conn *Connection) Close() {
	// 数据库更新用户为offline
	conn.wsConn.Close()
	conn.mutex.Lock()
	if !conn.isClosed {
		close(conn.closeChan)
		conn.isClosed = true
	}
	conn.mutex.Unlock()
}

// WriteMessage 写消息 把消息放入通道中outChan
func (conn *Connection) WriteMessage(data []byte) (err error) {
	select {
	case conn.outChan <- data:
	case <-conn.closeChan:
		err = errors.New("WriteMessage connection is closed")
	}
	return
}

// ReadMessage 读消息 从数通道中读取inChan
func (conn *Connection) ReadMessage() (data []byte, err error) {
	select {
	case data = <-conn.inChan:
		// 更新最后消息时间
		conn.heartbeatTime = time.Now()
	case <-conn.closeChan:
		err = errors.New("ReadMessage connection is closed")
	}
	return
}

// 监听读消息 读取客户端消息，放入通道inChan
func (conn *Connection) readLoop() {
	var (
		data []byte
		err  error
	)
	for {
		if _, data, err = conn.wsConn.ReadMessage(); err != nil {
			conn.Close()
			return
		}
		select {
		case conn.inChan <- data:
		case <-conn.closeChan:
			return
		}
	}
}

// 监听写消息 在通道outChan 取数据，推送给客户端
func (conn *Connection) writeLoop() {
	var (
		data []byte
		err  error
	)
	for {
		select {
		case data = <-conn.outChan:
		case <-conn.closeChan:
			return
		}
		if err = conn.wsConn.WriteMessage(websocket.TextMessage, data); err != nil {
			conn.Close()
			return
		}
	}
}

// 2秒定时器检测是否超时
func (conn *Connection) checkHeartbeatTimeout() {
	t := time.NewTicker(time.Second * 2)
	defer t.Stop()
	for {
		// fmt.Println("check HeartbeatTimeout Ticker running...")
		<-t.C
		if conn.isClosed == true {
			return
		}
		// 现在大于消息超时时间
		if time.Now().After(conn.heartbeatTime.Add(conn.heartbeatTimeOut)) {
			conn.Close()
			return
		}
	}
}

// InitConnection 初始化
func InitConnection(wsConn *websocket.Conn) (conn *Connection, err error) {
	conn = &Connection{
		wsConn:           wsConn,
		inChan:           make(chan []byte, 1000),
		outChan:          make(chan []byte, 1000),
		closeChan:        make(chan byte, 1),
		heartbeatTimeOut: time.Second * time.Duration(30),
		heartbeatTime:    time.Now(),
	}
	// 启动读协程
	go conn.readLoop()
	// 启动写协程
	go conn.writeLoop()
	// 定时器检查是否超时
	go conn.checkHeartbeatTimeout()

	// 初始化完成
	// 这时候可以设置会员为在线 online
	return
}
