// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameFollower = "follower"

// Follower mapped from table <follower>
type Follower struct {
	UserID   int64 `gorm:"column:user_id;type:bigint;primaryKey" json:"user_id"`
	ToUserID int64 `gorm:"column:to_user_id;type:bigint;primaryKey" json:"to_user_id"`
	Flag     int64 `gorm:"column:flag;type:tinyint(1);not null" json:"flag"`
}

// TableName Follower's table name
func (*Follower) TableName() string {
	return TableNameFollower
}