package jws

import "sync"

//获取服务器实例
var serverList map[string]*Server

//获取服务器列表
func GetServerList() []*Server {
	serverCnt := len(serverList)
	if serverCnt == 0 {
		return nil
	}
	var cleanServerList []*Server
	for _, value := range serverList {
		cleanServerList = append(cleanServerList, value)
	}
	return cleanServerList
}

//获取实例
func NewServer(serverId string, serverName string) *Server {
	if serverList == nil {
		serverList = make(map[string]*Server)
	}
	var server *Server
	if _, ok := serverList[serverId]; ok {
		server = serverList[serverId]
	} else {
		server = &Server{
			ServerId:   serverId,
			ServerName: serverName,
			ClientList: make([]*Client, 0),
		}
		server = &Server{
			ServerId:   serverId,
			ServerName: serverName,
			ClientList: make([]*Client, 0),
			IsRunning:  true,
			wg:         sync.WaitGroup{},
			lock:       sync.Mutex{},
		}
		serverList[serverId] = server
	}
	return server
}

//根据ID查询服务器
func GetServerByServerId(serverId string) *Server {
	if len(serverList) > 0 {
		if server, ok := serverList[serverId]; ok {
			return server
		}
	}
	return nil
}

//通过服务器发送信息到对应的客户端
func SendTextMessageToServer(server *Server, textContent string) bool {
	wg := sync.WaitGroup{}
	clientCnt := server.GetClientCnt()
	if clientCnt > 0 {
		wg.Add(int(clientCnt))
		clientList := server.GetClientList()
		for _, client := range clientList {
			go func() {
				client.SendTextMessage(textContent)
				wg.Done()
			}()
		}
	}
	wg.Wait()
	return true
}
