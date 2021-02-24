package jws

import "sync"

const ServerMaxClientNum = 1000 //群员最大数量

//服务器
type Server struct {
	ServerId   string    //服务器ID
	ServerName string    //服务器名称
	ClientList []*Client //客户端ID列表
	IsRunning  bool      //是否正在运行
	wg         sync.WaitGroup
	lock       sync.Mutex //锁
}

//新增客户端
func (s *Server) AddClient(client *Client) bool {
	s.ClientList = append(s.ClientList, client)
	return true
}

//删除客户端
func (s *Server) RemoveClient(client *Client) bool {
	return true
}

//获取客户端列表
func (s *Server) GetClientList() []*Client {
	return s.ClientList
}

//获取客户端数量
func (s *Server) GetClientCnt() int64 {
	cnt:=len(s.ClientList)
	return int64(cnt)
}

//关闭服务端
func (s *Server) Close() bool {
	s.lock.Lock()
	defer s.lock.Unlock()
	if s.IsRunning {
		clientCnt := s.GetClientCnt()
		s.wg.Add(int(clientCnt))
		for _, client := range s.ClientList {
			go func() {
				client.Close()
				s.wg.Done()
			}()
		}
		s.wg.Wait()
		s.IsRunning = false
	}
	return true
}
