package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/satori/go.uuid"
	"net/http"
)

type ClientManager struct {
	clients    map[*Client]bool //所有连接的 client
	broadcast  chan []byte      //广播消息
	register   chan *Client     //有 client 连接进来时进行通信的 channel
	unregister chan *Client     //有 client 关闭时进行通信的 channel
}

type Client struct {
	id     string          //client 分配的 id
	socket *websocket.Conn //socket
	send   chan []byte     //发送的消息
}

func (c *Client) read() {
	defer func() {
		manager.unregister <- c
		c.socket.Close()
	}()

	for {
		_, message, err := c.socket.ReadMessage() //通过 socket 读取数据，如果发生错误表示 client 发生的了关闭
		if err != nil {
			manager.unregister <- c //通知注册的关闭的 channel
			c.socket.Close()
			break //退出循环
		}
		jsonMessage, _ := json.Marshal(&Message{Sender: c.id, Content: string(message)})
		fmt.Println(jsonMessage)
		manager.broadcast <- jsonMessage
	}
}

func (c *Client) write() {
	defer func() {
		c.socket.Close()
	}()

	for {
		select {
		case message, ok := <-c.send:
			if !ok {
				c.socket.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			c.socket.WriteMessage(websocket.TextMessage, message)
		}
	}
}

type Message struct {
	Sender    string `json:"sender,omitempty"`
	Recipient string `json:"recipient,omitempty"`
	Content   string `json:"content,omitempty"`
}

func (manager *ClientManager) start() {
	for {
		select {
		case conn := <-manager.register:
			manager.clients[conn] = true
			jsonMessage, _ := json.Marshal(&Message{Content: "A new socket connected"})
			manager.send(jsonMessage, conn) //发送消息时忽略当前的 client，表示当前 clien 连接成功用来通知其他 client
		case conn := <-manager.unregister:
			if _, ok := manager.clients[conn]; ok {
				close(conn.send)
				delete(manager.clients, conn)
				jsonMessage, _ := json.Marshal(&Message{Content: "/A socket has disconnected."})
				manager.send(jsonMessage, conn)
			}
		case message := <-manager.broadcast:
			for conn := range manager.clients {
				select {
				case conn.send <- message:
				default:
					close(conn.send)
					delete(manager.clients, conn)
				}
			}
		}
	}
}

func (manager *ClientManager) send(message []byte, ignore *Client) {
	for conn := range manager.clients {
		if conn != ignore {
			conn.send <- message
		}
	}
}

func wsPage(res http.ResponseWriter, req *http.Request) {
	conn, error := (&websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}).Upgrade(res, req, nil)
	if error != nil {
		http.NotFound(res, req)
		return
	}
	client := &Client{id: uuid.NewV4().String(), socket: conn, send: make(chan []byte)}

	manager.register <- client

	go client.read()
	go client.write()
}

var manager = ClientManager{
	clients:    make(map[*Client]bool),
	broadcast:  make(chan []byte),
	register:   make(chan *Client),
	unregister: make(chan *Client),
}

func main() {

	fmt.Println("Starting application...")
	go manager.start()
	http.HandleFunc("/ws", wsPage)
	http.ListenAndServe(":12345", nil)
}
