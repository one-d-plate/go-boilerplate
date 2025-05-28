package dto

type RegisterUserRequest struct {
	Name     string `json:"name" validate:"required"`
	Username string `json:"username" validate:"required,min=3,max=20"`
	Password string `json:"password" validate:"required,min=6"`
	Level    string `json:"level" validate:"required"`
}

type LoginUserRequest struct {
	Username  string `json:"username" validate:"required,min=3,max=20"`
	Password  string `json:"password" validate:"required,min=6"`
	IpAddress string `json:"ip_address" validate:"omitempty"`
	ClientId  string `json:"client_id" validate:"omitempty"`
}

type LoginUserResponse struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}
