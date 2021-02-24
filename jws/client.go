package jws

import (
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"sync"
)

const ClientMaxGroupNum = 100   //最大群组数量
const ClientMaxMsgByteNum = 200 //一次最多发送消息长度

//客户端
type Client struct {
	ClientId    string          //客户端ID
	ServerId    string          //节点服务器ID
	UserId      string          //用户ID
	GroupIdList []string        //群组
	conn        *websocket.Conn //连接
	readCh      chan []byte     //读取消息
	writeCh     chan []byte     //发送消息
	IsRunning   bool            //是否正在运行
	lock        sync.Mutex      //锁

}

//获取客户端ID
func (c *Client) GetClientId() string {
	return c.ClientId
}

//绑定用户ID
func (c *Client) BindUserId(userId string) bool {
	return true
}

//解绑
func (c *Client) UnBindUserId() bool {
	return true
}

//向当前客户端发送信息
func (c *Client) SendMessage(messageId string) bool {
	return true
}

//发送文本信息
func (c *Client) SendTextMessage(textContent string) bool {
	if c.IsRunning {
		c.writeCh <- []byte(textContent)
	}
	return true
}

//加入群组
func (c *Client) JoinGroup(groupId string) bool {
	return true
}

//离开群组
func (c *Client) LeaveGroup(groupId string) bool {
	return true
}

//群发信息,excludeClientIdList 排除这些客户端
func (c *Client) SendToGroup(groupId string, messageId string, excludeClientIdList []string) bool {
	return true
}

//关闭客户端并通知
func (c *Client) Close() bool {
	c.lock.Lock()
	if c.IsRunning {
		for {
			writeChLen := len(c.writeCh)
			if writeChLen == 0 {
				close(c.readCh)
				close(c.writeCh)
				_ = c.conn.Close()
				c.IsRunning = false
				break
			}
		}
	}
	defer c.lock.Unlock()
	return true
}

//读取消息
func (c *Client) ReadMessage() error {
	go func() {
		for {
			_, message, err := c.conn.ReadMessage()
			if err != nil {
				logrus.Println("信息读取错误：" + err.Error())
				break
			}
			c.writeCh <- []byte(message)
		}
	}()
	return nil
}

//发送文本信息
func (c *Client) TestTextMessage() error {
	go func() {
		for {
			select {
			case msg := <-c.writeCh:
				_ = c.conn.WriteMessage(MessageTypeText, msg)
			default:
			}
		}
	}()
	return nil
}

//获取客户端实例
var clientList map[string]*Client

func GetClientList() map[string]*Client {
	return clientList
}

//获取客户端实例
func NewClient(clientId string, serverId string, conn *websocket.Conn) *Client {
	if clientList == nil {
		clientList = make(map[string]*Client)
	}
	var client *Client
	if _, ok := clientList[clientId]; ok {
		client = clientList[clientId]
	} else {
		client = &Client{
			ClientId:    clientId,
			ServerId:    serverId,
			UserId:      "",
			GroupIdList: make([]string, 0),
			conn:        conn,
			readCh:      make(chan []byte, 0),
			writeCh:     make(chan []byte, 0),
			IsRunning:   true,
			lock:        sync.Mutex{},
		}
		clientList[clientId] = client
	}
	return client
}
