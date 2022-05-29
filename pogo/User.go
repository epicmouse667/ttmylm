package pogo

type User struct {
	Id              int    `json:"id,omitempty"`
	Name            string `json:"name,omitempty"`
	FollowCount     int    `json:"follow_count,omitempty"`
	FollowerCount   int    `json:"follower_count,omitempty"`
	IsFollow        bool   `json:"is_follow,omitempty"`
	Avatar          string `json:"avatar,omitempty"`
	Signature       string `json:"signature"`
	BackgroundImage string `json:"background_image"`
}
