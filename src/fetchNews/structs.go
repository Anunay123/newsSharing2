package fetchNews

type ApiResponse struct {
	Heading      string  `json:"heading"`
	Data         []*News `json:"data"`
	PageNumber   int     `json:"pageNumber"`
	TotalPages   int     `json:"totalPages,omitempty"`
	ErrorMessage string  `json:"errorMessage"`
}

type News struct {
	Title     string `json:"title"`
	Url       string `json:"url"`
	Publisher string `json:"publisher"`
}
