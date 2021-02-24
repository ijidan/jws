package jws

import "sync"

//工具类
type Tools struct {
	wg   sync.WaitGroup
	lock sync.Mutex //锁
}

//获取所有的客户端
func (t *Tools) GetAllClientList() []*Client {
	serverList := GetServerList()
	allClientList := make([]*Client, 0)
	for _, server := range serverList {
		clientList := server.GetClientList()
		if len(clientList) > 0 {
			for _, client := range clientList {
				allClientList = append(allClientList, client)
			}
		}
	}
	return allClientList
}

//获取所有的客户端总数
func (t *Tools) GetAllClientCnt() int64 {
	allClientList := t.GetAllClientList()
	return int64(len(allClientList))
}

//根据客户端ID查询客户端
func (t *Tools) GetClientByClientId(clientId string) *Client {
	allClientList := t.GetAllClientList()
	if len(allClientList) > 0 {
		for _, client := range allClientList {
			currentClientId := client.ClientId
			if currentClientId == clientId {
				return client
			}
		}
	}
	return nil
}

//获取所有在线客户端
func (t *Tools) GetAllOnlineClientList() []*Client {
	return nil
}

//获取所有在线客户端总数
func (t *Tools) GetAllOnlineClientCnt() int64 {
	return 0
}

//判断客户端是否在线
func (t *Tools) IsOnline(client *Client) bool {
	return true
}

//获取所有连接数
func (t *Tools) GetAllClientCount() int64 {
	return 0
}

//关闭客户端并通知
func (t *Tools) CloseClient(client *Client, messageId string) bool {
	return true
}

//关闭客户端
func (t *Tools) KickOff(client *Client, textContent string) bool {
	if len(textContent) > 0 {
		client.SendTextMessage(textContent)
	}
	defer client.Close()
	return true
}

//向客户端发送信息
func (t *Tools) SendMessageToClient(client *Client, messageId string) bool {
	return true
}

//发送文本信息
func (t *Tools) SendTextMessageToClient(client *Client, textContent string) bool {
	return client.SendTextMessage(textContent)
}

// 向指定客户端或者排除客户端之后的所有客户端发送信息
func (t *Tools) SendMessageToAll(message string, clientList []*Client, excludeClientList []*Client) bool {
	return true
}

//群发信息,excludeClientList 排除这些客户端
func (t *Tools) SendMessageToGroup(groupId string, messageId string, excludeClientList []*Client) bool {
	return true
}

//获取工具实例
func NewTools() *Tools {
	tools := &Tools{
		wg:   sync.WaitGroup{},
		lock: sync.Mutex{},
	}
	return tools
}
