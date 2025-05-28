package dto

type RegisterClientRequest struct {
	Name         string `json:"name" validate:"required,min=3"`
	Description  string `json:"description" validate:"omitempty"`
	RedirectURI  string `json:"redirect_uri" validate:"omitempty,url"`
	ContactEmail string `json:"contact_email" validate:"omitempty,email"`
}

type RegisterClientResponse struct {
	Name      string `json:"name"`
	ApiKey    string `json:"api_key"`
	ApiSecret string `json:"api_secret"`
}

type GrantTokenRequest struct {
	ApiKey    string `json:"api_key" validate:"required"`
	ApiSecret string `json:"api_secret" validate:"required"`
	IpAddress string `json:"ip_address" validate:"omitempty"`
}

type GrantTokenResponse struct {
	GrantToken string `json:"grant_token"`
	ExpiresIn  int64  `json:"expires_in"` // dalam detik
}
