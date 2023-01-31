package dto

type User struct {
	Id            uint   `json:"id"`
	Name          string `json:"name"`
	FollowCount   uint   `json:"follow_count"`
	FollowerCount uint   `json:"follower_count"`
	IsFollow      bool   `json:"is_follow"`
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
	Id      uint   `json:"id"`
	Content string `json:"content"`
	// yyyy-MM-dd HH:MM:ss
	CreateTime string `json:"create_time"`
}
