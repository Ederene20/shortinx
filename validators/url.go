package validators

type CreateUrlInput struct {
	LongUrl string `json:"long_url" binding:"required"`
}