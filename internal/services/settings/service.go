package settings

import "context"

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

type Settings struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (s *Service) Get(ctx context.Context) (map[string]interface{}, error) {
	return map[string]interface{}{
		"id":       1,
		"name":     "ABC clinic",
		"language": "en",
	}, nil
}
