package dto

type SignUpRequest struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Identity string `json:"identity"`
	Password string `json:"password"`
}

type SwipeRequest struct {
	UserID    string `json:"user_id"`
	TargetID  string `json:"target_id"`
	Direction string `json:"direction"`
	IsPremium bool   `json:"is_premium"`
}
