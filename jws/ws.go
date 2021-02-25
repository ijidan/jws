package jws

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

const HOST = "localhost"
const PORT = 9999

//地址
var addr = fmt.Sprintf("%s:%d", HOST, PORT)

//升级地址
var upGrader = &websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

//ws类
type webSocket struct {
	upGrader *websocket.Upgrader
	server   *Server
	conn     chan *websocket.Conn
	err      error
	errorMsg string
}

//处理连接
func (ws webSocket) handleConnection(w http.ResponseWriter, r *http.Request) {
	c, err := upGrader.Upgrade(w, r, nil)
	if err != nil {
		ws.err = err
		ws.errorMsg = err.Error()
	}
	//写连接
	go func() {
		ws.conn <- c
	}()
	//读连接
	go func() {
		for {
			select {
			case conn := <-ws.conn:
				addr := conn.RemoteAddr().String()
				serverId := ws.server.ServerId
				client := NewClient(addr, serverId, conn)
				ws.server.AddClient(client)
				_ = client.ReadMessage()
				_ = client.TestTextMessage()
			default:
			}
		}
	}()
}

//启动WS
func StartWSServer() {
	server := NewServer("server_id_1", "server_name_1")
	ws := webSocket{
		upGrader: upGrader,
		server:   server,
		conn:     make(chan *websocket.Conn),
		err:      nil,
		errorMsg: "",
	}
	http.HandleFunc("/", ws.handleConnection)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Println("WS监听失败：" + err.Error())
	}else{
		log.Println("WS监听成功："+addr)
	}
}
