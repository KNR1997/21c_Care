package responses

type LoginResponse struct {
	AccessToken  string    `json:"accessToken"`
	RefreshToken string    `json:"refreshToken"`
	Exp          int64     `json:"exp"`
	Permissions  [1]string `json:"permissions"`
	Role         string    `json:"role"`
}

func NewLoginResponse(token, refreshToken string, exp int64, permissions [1]string, role string) *LoginResponse {
	return &LoginResponse{
		AccessToken:  token,
		RefreshToken: refreshToken,
		Exp:          exp,
		Permissions:  permissions,
		Role:         role,
	}
}
