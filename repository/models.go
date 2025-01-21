package repository

import "time"

type User struct {
	UUID      string    `json:"uuid"`
	Email     string    `json:"email"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	IsPremium bool      `json:"is_premium"`
	CreatedAt time.Time `json:"created_at"`
}

type Swipe struct {
	UUID      string    `json:"uuid"`
	UserID    string    `json:"user_id"`
	TargetID  string    `json:"target_id"`
	Direction string    `json:"direction"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
