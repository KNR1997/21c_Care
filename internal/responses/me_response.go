package responses

type MeResponse struct {
	Id    uint   `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

func NewMeResponse(id uint, email string, name string) *MeResponse {
	return &MeResponse{
		Id:    id,
		Email: email,
		Name:  name,
	}
}
