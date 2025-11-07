package timepad

// Структуры только для нужных полей
type PosterImage struct {
	DefaultURL    string `json:"default_url"`
	UploadcareURL string `json:"uploadcare_url"`
}

type Organization struct {
	Name string `json:"name"`
}

type Category struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Event struct {
	StartsAt     string       `json:"starts_at"`
	Name         string       `json:"name"`
	URL          string       `json:"url"`
	PosterImage  PosterImage  `json:"poster_image"`
	Organization Organization `json:"organization"`
	Categories   []Category   `json:"categories"`
}

type Response struct {
	Total  int     `json:"total"`
	Values []Event `json:"values"`
}
