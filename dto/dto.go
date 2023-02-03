package dto

type User struct {
	Id            uint   `json:"id"`
	Name          string `json:"name"`
	FollowCount   uint   `json:"follow_count"`
	FollowerCount uint   `json:"follower_count"`
	IsFollow      bool   `json:"is_follow"`
	Avatar        string `json:"avatar,omitempty"`
}

type Video struct {
	Id            uint   `json:"id"`
	Author        User   `json:"author"`
	PlayUrl       string `json:"play_url"`
	CoverUrl      string `json:"cover_url"`
	FavoriteCount uint   `json:"favorite_count"`
	CommentCount  uint   `json:"comment_count"`
	IsFavorite    bool   `json:"is_favorite"`
	Title         string `json:"title"`
}

type Comment struct {
	Id      uint   `json:"id"`
	User    User   `json:"user"`
	Content string `json:"content"`
	// mm-dd
	CreateDate string `json:"create_date"`
}

type Message struct {
	Id         uint   `json:"id"`
	Content    string `json:"content"`
	FromUserId uint   `json:"from_user_id"`
	ToUserId   uint   `json:"to_user_id"`
	// yyyy-MM-dd HH:MM:ss
	CreateTime string `json:"create_time"`
}

const (
	RECEIVE = 0
	SEND    = 1
)

type FriendUser struct {
	User
	Message string `json:"message"`
	// message消息的类型，0 => 当前请求用户接收的消息， 1 => 当前请求用户发送的消息
	MsgType int64 `json:"msgType"`
}
