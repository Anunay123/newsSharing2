package guardian

import "time"

type NewsData struct {
	Response *Response `json:"response"`
}

func (nd *NewsData) GetResponse() *Response {
	if nd != nil {
		return nd.Response
	}

	return nil
}

//Response
type Response struct {
	Status      string    `json:"status"`
	UserTier    string    `json:"userTier"`
	Total       int       `json:"total"`
	StartIndex  int       `json:"startIndex"`
	PageSize    int       `json:"pageSize"`
	CurrentPage int       `json:"currentPage"`
	Pages       int       `json:"pages"`
	OrderBy     string    `json:"orderBy"`
	Results     []*Result `json:"results"`
}

func (response *Response) GetPages() int {
	if response != nil {
		return response.Pages
	}

	return 0
}

func (response *Response) GetCurrentPage() int {
	if response != nil {
		return response.CurrentPage
	}

	return 0
}

func (response *Response) GetResults() []*Result {
	if response != nil {
		return response.Results
	}

	return nil
}

type Result struct {
	Id                 string    `json:"id"`
	Type               string    `json:"type"`
	SectionId          string    `json:"sectionId"`
	SectionName        string    `json:"sectionName"`
	WebPublicationDate time.Time `json:"webPublicationDate"`
	WebTitle           string    `json:"webTitle"`
	WebUrl             string    `json:"webUrl"`
	ApiUrl             string    `json:"apiUrl"`
	IsHosted           bool      `json:"isHosted"`
	PillarId           string    `json:"pillarId"`
	PillarName         string    `json:"pillarName"`
}

func (result *Result) GetWebUrl() string {
	if result != nil {
		return result.WebUrl
	}

	return ""
}

func (result *Result) GetWebTitle() string {
	if result != nil {
		return result.WebTitle
	}

	return ""
}
