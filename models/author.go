package models

type Author struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Role      string `json:"role"`
	Place     string `json:"place"`
	AvatarURL string `json:"avatar_url"`
}
