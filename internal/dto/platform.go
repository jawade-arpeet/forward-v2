package dto

type CreatePlatformRequest struct {
	Name    string `json:"name"`
	LogoURL string `json:"logo_url"`
}

type CreatePlatformResponse struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	LogoURL string `json:"logo_url"`
}

func NewCreatePlatformResponse(
	id string,
	name string,
	logoURL string,
) *CreatePlatformResponse {
	return &CreatePlatformResponse{
		ID:      id,
		Name:    name,
		LogoURL: logoURL,
	}
}
