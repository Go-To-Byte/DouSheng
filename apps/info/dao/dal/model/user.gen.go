// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameUser = "user"

// User mapped from table <user>
type User struct {
	ID       int64  `gorm:"column:id;type:bigint;primaryKey" json:"id"`
	Username string `gorm:"column:username;type:varchar(16);not null" json:"username"`
	Passwd   string `gorm:"column:passwd;type:char(128);not null" json:"passwd"`
}

// TableName User's table name
func (*User) TableName() string {
	return TableNameUser
}