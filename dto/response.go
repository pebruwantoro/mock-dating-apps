package dto

type SignUpResponse struct {
	Id string `json:"id"`
}

type LoginResponse struct {
	Id    string `json:"id"`
	Token string `json:"token"`
}

type SwipeResponse struct {
	Id string `json:"id"`
}

type PurchasePremiumPackageResponse struct {
	Id string `json:"id"`
}

type BaseResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
