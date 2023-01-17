package controller

type Resp struct {
	StatusCode int    `json:"statusCode"`
	StatusMsg  string `json:"statusMsg,omitempty"`
}

type User struct {
	Id            uint   `json:"id"`
	Name          string `json:"name"`
	FollowCount   uint   `json:"followCount"`
	FollowerCount uint   `json:"followerCount"`
	IsFollow      uint   `json:"isFollow"`
}

type Video struct {
	Id            uint   `json:"id"`
	Author        User   `json:"author"`
	PlayUrl       string `json:"playUrl"`
	CoverUrl      string `json:"coverUrl"`
	FavoriteCount uint   `json:"favoriteCount"`
	CommentCount  uint   `json:"commentCount"`
	IsFavorite    uint   `json:"isFavorite"`
	Title         string `json:"title"`
}

type Comment struct {
	Id      uint   `json:"id"`
	User    User   `json:"user"`
	Content string `json:"content"`
	//mm-dd
	CreateDate string `json:"createDate"`
}
