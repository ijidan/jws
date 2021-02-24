package jws

const GroupMaxClientNum = 1000 //群员最大数量

//群主
type Group struct {
	groupId      string      //群组ID
	groupName    string      //群组名称
	clientIdList []*Client   //客户端ID列表
	readMsgCh    chan []byte //接收消息
}

//获取群组ID
func (g *Group) GetGroupId() string {
	return g.groupId
}

//获取群组名称
func (g *Group) GetGroupName() string {
	return g.groupName
}

//向当前客户端发送信息
func (g *Group) SendMessage(messageId string) bool {
	return true
}

//获取群组所有客户端
func (g *Group) GetAllClientList(groupId string) []*Client {
	return g.clientIdList
}

//获取群组所有客户端总数
func (g *Group) GetAllClientCnt(groupId string) int64 {
	return int64(len(g.clientIdList))
}

//获取群组在线客户端
func (g *Group) GetOnlineClientList(groupId string) []string {
	return nil
}

//获取群组在线客户端总数
func (g *Group) GetOnlineClientCnt(groupId string) int64 {
	return 0
}

//获取群组实例
func NewGroup(groupId string, groupName string) *Group {
	group := &Group{
		groupId:      groupId,
		groupName:    groupName,
		clientIdList: make([]*Client, GroupMaxClientNum),
		readMsgCh:    nil,
	}
	return group
}
