package jws

import (
	"github.com/bwmarrin/snowflake"
	"time"
)

//消息类型

const (
	MessageTypeText   = 1  //文本
	MessageTypeBinary = 2  //二进制
	MessageTypeImg    = 2  //图片
	MessageTypeVideo  = 3  //视频
	MessageTypeClose  = 8  //关闭
	MessageTypePing   = 9  //ping
	MessageTypePong   = 10 //pong
)

//消息
type Message struct {
	id         string
	msgType    int64
	fromUserId int64
	toUserId   int64
	createTime int64
}

//文本消息
type TextMessage struct {
	Message
	textContent string
}

//图片消息
type videoMessage struct {
	Message
	imgPath string
	imgUrl  string
	imgSize int64
	imgExt  string
}

//视频消息
type VideoMessage struct {
	Message
	videoUrl  string
	videoSize int64
	videoExt  string
}

//生成消息ID
func genMessageId() string {
	node, _ := snowflake.NewNode(1)
	id := string(node.Generate().Int64())
	return "msg-" + id
}

//获取消息实例
func NewMessage(msgType int64, fromUserId int64, toUserId int64, createTime int64) *Message {
	node, _ := snowflake.NewNode(1)
	id := string(node.Generate().Int64())
	if createTime == 0 {
		createTime = int64(time.Now().Unix())
	}
	message := &Message{
		id:         id,
		msgType:    msgType,
		fromUserId: fromUserId,
		toUserId:   toUserId,
		createTime: createTime,
	}
	return message
}

//获取文本消息实例
func NewTextMessage(fromUserId int64, toUserId int64, createTime int64, textContent string) *TextMessage {
	id := genMessageId()
	if createTime == 0 {
		createTime = int64(time.Now().Unix())
	}
	textMessage := &TextMessage{
		textContent: textContent,
	}
	//设置属性
	textMessage.id = id
	textMessage.msgType = MessageTypeText
	textMessage.fromUserId = fromUserId
	textMessage.toUserId = toUserId
	textMessage.createTime = createTime
	return textMessage
}

//获取图片消息实例
func NewImgMessage(fromUserId int64, toUserId int64, createTime int64, imgPath string, imgUrl string, imgSize int64, imgExt string) *videoMessage {
	id := genMessageId()
	if createTime == 0 {
		createTime = int64(time.Now().Unix())
	}
	imgMessage := &videoMessage{
		imgPath: imgPath,
		imgUrl:  imgUrl,
		imgSize: imgSize,
		imgExt:  imgExt,
	}
	//设置属性
	imgMessage.id = id
	imgMessage.msgType = MessageTypeImg
	imgMessage.fromUserId = fromUserId
	imgMessage.toUserId = toUserId
	imgMessage.createTime = createTime

	return imgMessage
}

//获取视频消息实例
func NewVideoMessage(fromUserId int64, toUserId int64, createTime int64, videoUrl string, videoSize int64, videoExt string) *VideoMessage {
	id := genMessageId()
	if createTime == 0 {
		createTime = int64(time.Now().Unix())
	}
	videoMessage := &VideoMessage{
		videoUrl:  videoUrl,
		videoSize: videoSize,
		videoExt:  videoExt,
	}
	//设置属性
	videoMessage.id = id
	videoMessage.msgType = MessageTypeVideo
	videoMessage.fromUserId = fromUserId
	videoMessage.toUserId = toUserId
	videoMessage.createTime = createTime

	return videoMessage
}
