package telematic

type User struct {
	Id    int    `json:"id"`
	Email string `json:"email"`
}

type Permission struct {
	View   bool `json:"view"`
	Edit   bool `json:"edit"`
	Remove bool `json:"remove"`
}

type AuthRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthResponse struct {
	Status      int                   `json:"status"`
	UserApiHash string                `json:"user_api_hash"`
	Permissions map[string]Permission `json:"permissions"`
}
