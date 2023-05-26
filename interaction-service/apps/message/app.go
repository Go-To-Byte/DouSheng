// @Author: Hexiaoming 2023/2/15
package message

import (
	"github.com/go-playground/validator/v10"
	"time"
)

const (
	AppName = "message"
)

var (
	validate = validator.New()
)

func NewMessagePo(req *ChatMessageActionRequest) *MessagePo {
	return &MessagePo{
		ToUserId:   req.ToUserId,
		FromUserId: 0,
		Content:    req.Content,
		CreatedAt:  time.Now().Unix(),
	}
}

func (r *ChatMessageListRequest) Validate() error {
	return validate.Struct(r)
}

func NewChatMessageListRequest() *ChatMessageListRequest {
	return &ChatMessageListRequest{}
}

func NewChatMessageListResponse() *ChatMessageListResponse {
	return &ChatMessageListResponse{}
}

// Validate 发送聊天消息 相关
func (r *ChatMessageActionRequest) Validate() error {
	return validate.Struct(r)
}

func NewChatMessageActionRequest() *ChatMessageActionRequest {
	return &ChatMessageActionRequest{}
}

func NewChatMessageActionResponse() *ChatMessageActionResponse {
	return &ChatMessageActionResponse{}
}

// TableName 指明表名 -> gorm 参数映射
func (*MessagePo) TableName() string {
	return "message"
}
