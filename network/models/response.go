// Author: BeYoung
// Date: 2023/1/26 0:13
// Software: GoLand

package models

type FeedResponse struct {
	NextTime   int64   `json:"next_time"`
	StatusCode int64   `json:"status_code"`
	StatusMsg  string  `json:"status_msg"`
	VideoList  []Video `json:"video_list"`
}

type Video struct {
	Author        User   `json:"author"`
	CommentCount  int64  `json:"comment_count"`
	CoverURL      string `json:"cover_url"`
	FavoriteCount int64  `json:"favorite_count"`
	ID            int64  `json:"id"`
	IsFavorite    bool   `json:"is_favorite"`
	PlayURL       string `json:"play_url"`
	Title         string `json:"title"`
}

type RegisterResponse struct {
	StatusCode int64  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
	Token      string `json:"token"`
	UserID     int64  `json:"user_id"`
}

type LoginResponse struct {
	StatusCode int64  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
	Token      string `json:"token"`
	UserID     int64  `json:"user_id"`
}

type InfoResponse struct {
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

type PublishResponse struct {
	StatusCode int64  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}

type PublishListResponse struct {
	StatusCode int64   `json:"status_code"`
	StatusMsg  string  `json:"status_msg"`
	VideoList  []Video `json:"video_list"`
}

type FavoriteResponse struct {
	StatusCode int64  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}

type FavoriteListResponse struct {
	StatusCode string  `json:"status_code"`
	StatusMsg  string  `json:"status_msg"`
	VideoList  []Video `json:"video_list"`
}

type CommentResponse struct {
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

type CommentListResponse struct {
	Comments   []Comment `json:"comment_list"`
	StatusCode int64     `json:"status_code"`
	StatusMsg  string    `json:"status_msg"`
}

type FollowResponse struct {
	StatusCode int64  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}

type FollowListResponse struct {
	StatusCode string `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
	UserList   []User `json:"user_list"`
}

type FollowerListResponse struct {
	StatusCode string `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
	UserList   []User `json:"user_list"`
}

type FriendListResponse struct {
	StatusCode string `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
	UserList   []User `json:"user_list"`
}

type MessageResponse struct {
	StatusCode int64  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}

type MessageListResponse struct {
	MessageList []MessageList `json:"message_list"`
	StatusCode  string        `json:"status_code"`
	StatusMsg   string        `json:"status_msg"`
}

type MessageList struct {
	Content    string `json:"content"`
	CreateTime string `json:"create_time"`
	ID         int64  `json:"id"`
}
