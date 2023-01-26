// Author: BeYoung
// Date: 2023/1/26 0:13
// Software: GoLand

package data

type Feed struct {
	NextTime   int64       `json:"next_time"`
	StatusCode int64       `json:"status_code"`
	StatusMsg  string      `json:"status_msg"`
	VideoList  []VideoList `json:"video_list"`
}

type VideoList struct {
	Author        Author `json:"author"`
	CommentCount  int64  `json:"comment_count"`
	CoverURL      string `json:"cover_url"`
	FavoriteCount int64  `json:"favorite_count"`
	ID            int64  `json:"id"`
	IsFavorite    bool   `json:"is_favorite"`
	PlayURL       string `json:"play_url"`
	Title         string `json:"title"`
}

type Author struct {
	FollowCount   int64  `json:"follow_count"`
	FollowerCount int64  `json:"follower_count"`
	ID            int64  `json:"id"`
	IsFollow      bool   `json:"is_follow"`
	Name          string `json:"name"`
}

type UserRegister struct {
	StatusCode int64  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
	Token      string `json:"token"`
	UserID     int64  `json:"user_id"`
}

type UserLogin struct {
	StatusCode int64  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
	Token      string `json:"token"`
	UserID     int64  `json:"user_id"`
}

type UserInfo struct {
	StatusCode int64  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
	User       User   `json:"user"`
}

type User struct {
	FollowCount   int64  `json:"follow_count"`
	FollowerCount int64  `json:"follower_count"`
	ID            int64  `json:"id"`
	IsFollow      bool   `json:"is_follow"`
	Name          string `json:"name"`
}

type PublishAction struct {
	StatusCode int64  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}

type PublishList struct {
	StatusCode int64       `json:"status_code"`
	StatusMsg  string      `json:"status_msg"`
	VideoList  []VideoList `json:"video_list"`
}

type FavoriteAction struct {
	StatusCode int64  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}

type FavoriteList struct {
	StatusCode string      `json:"status_code"`
	StatusMsg  string      `json:"status_msg"`
	VideoList  []VideoList `json:"video_list"`
}

type CommentAction struct {
	Comment    Comment `json:"comment"`
	StatusCode int64   `json:"status_code"`
	StatusMsg  string  `json:"status_msg"`
}

type Comment struct {
	Content    string `json:"content"`
	CreateDate string `json:"create_date"`
	ID         int64  `json:"id"`
	User       User   `json:"user"`
}

type CommentList struct {
	Comments   []Comment `json:"comment_list"`
	StatusCode int64     `json:"status_code"`
	StatusMsg  string    `json:"status_msg"`
}

type RelationAction struct {
	StatusCode int64  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}

type RelationFollowList struct {
	StatusCode string     `json:"status_code"`
	StatusMsg  string     `json:"status_msg"`
	UserList   []UserList `json:"user_list"`
}

type UserList struct {
	FollowCount   int64  `json:"follow_count"`
	FollowerCount int64  `json:"follower_count"`
	ID            int64  `json:"id"`
	IsFollow      bool   `json:"is_follow"`
	Name          string `json:"name"`
}

type RelationFollowerList struct {
	StatusCode string     `json:"status_code"`
	StatusMsg  string     `json:"status_msg"`
	UserList   []UserList `json:"user_list"`
}

type RelationFriendList struct {
	StatusCode string     `json:"status_code"`
	StatusMsg  string     `json:"status_msg"`
	UserList   []UserList `json:"user_list"`
}

type MessageAction struct {
	StatusCode int64  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}

type MessageChat struct {
	MessageList []MessageList `json:"message_list"`
	StatusCode  string        `json:"status_code"`
	StatusMsg   string        `json:"status_msg"`
}

type MessageList struct {
	Content    string `json:"content"`
	CreateTime string `json:"create_time"`
	ID         int64  `json:"id"`
}
