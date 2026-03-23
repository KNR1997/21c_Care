package responses

type RefreshTokenResponse struct {
	AccessToken string `json:"accessToken"`
}

func NewRefreshTokenResponse(accessToken string) *RefreshTokenResponse {
	return &RefreshTokenResponse{
		AccessToken: accessToken,
	}
}
