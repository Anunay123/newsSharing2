package fetchNews

type ApiResponse struct {
	Heading    string  `json:"heading"`
	Data       []*News `json:"data"`
	PageNumber int     `json:"pageNumber"`
	TotalPages int     `json:"totalPages,omitempty"`
}

type News struct {
	Title     string `json:"title"`
	Url       string `json:"url"`
	Publisher string `json:"publisher"`
}
