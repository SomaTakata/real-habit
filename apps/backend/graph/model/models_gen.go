// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Follow struct {
	ID       string `json:"id"`
	Follower *User  `json:"follower"`
	Followee *User  `json:"followee"`
}

type Goal struct {
	ID            string    `json:"id"`
	User          *User     `json:"user"`
	Name          string    `json:"name"`
	Description   *string   `json:"description,omitempty"`
	StartDate     string    `json:"startDate"`
	EndDate       string    `json:"endDate"`
	RepeatSetting *string   `json:"repeatSetting,omitempty"`
	IsPublic      bool      `json:"isPublic"`
	Records       []*Record `json:"records"`
}

type Like struct {
	ID      string `json:"id"`
	Post    *Post  `json:"post"`
	User    *User  `json:"user"`
	LikedAt string `json:"likedAt"`
}

type Mutation struct {
}

type Notification struct {
	ID        string `json:"id"`
	User      *User  `json:"user"`
	Type      string `json:"type"`
	Content   string `json:"content"`
	CreatedAt string `json:"createdAt"`
}

type Post struct {
	ID        string  `json:"id"`
	Record    *Record `json:"record"`
	CreatedAt string  `json:"createdAt"`
	IsPublic  bool    `json:"isPublic"`
	Likes     []*Like `json:"likes"`
}

type Query struct {
}

type Record struct {
	ID        string  `json:"id"`
	Goal      *Goal   `json:"goal"`
	Timestamp string  `json:"timestamp"`
	PhotoURL  string  `json:"photoURL"`
	Comment   *string `json:"comment,omitempty"`
	Post      *Post   `json:"post"`
}

type User struct {
	ID             string          `json:"id"`
	Name           string          `json:"name"`
	Email          string          `json:"email"`
	ProfilePicture *string         `json:"profilePicture,omitempty"`
	CreatedAt      string          `json:"createdAt"`
	Goals          []*Goal         `json:"goals"`
	Notifications  []*Notification `json:"notifications"`
}
