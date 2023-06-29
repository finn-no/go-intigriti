package v2

type User struct {
	UserID    string `json:"userId"`
	UserName  string `json:"userName"`
	AvatarURL string `json:"avatarUrl"`
	Role      string `json:"role"`
	Email     string `json:"email"`
}
