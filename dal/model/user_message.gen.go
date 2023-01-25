// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameUserMessage = "user_message"

// UserMessage mapped from table <user_message>
type UserMessage struct {
	MessageID      int64  `gorm:"column:message_id;type:bigint;primaryKey" json:"message_id"`
	UserId1        int64  `gorm:"column:user_id1;type:bigint;not null" json:"user_id1"`
	UserId2        int64  `gorm:"column:user_id2;type:bigint;not null" json:"user_id2"`
	MessageContent string `gorm:"column:message_content;type:varchar(128);not null" json:"message_content"`
}

// TableName UserMessage's table name
func (*UserMessage) TableName() string {
	return TableNameUserMessage
}
