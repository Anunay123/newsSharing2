package nyTimes

type NewsData struct {
	Status    string    `json:"status"`
	Copyright string    `json:"copyright"`
	Response  *Response `json:"response"`
}

func (newsData *NewsData) GetResponse() *Response {
	if newsData != nil {
		return newsData.Response
	}
	return nil
}

type Response struct {
	Docs []*Doc `json:"docs"`
	Meta *Meta  `json:"meta"`
}

func (response *Response) GetDocs() []*Doc {
	if response != nil {
		return response.Docs
	}
	return nil
}

type Meta struct {
	Hits   int `json:"hits"`
	Offset int `json:"offset"`
	Time   int `json:"time"`
}

type Doc struct {
	Abstract string `json:"abstract"`
	WebUrl   string `json:"web_url"`
	//Snippet  string `json:"snippet"`
	//LeadParagraph string `json:"lead_paragraph"`
	//PrintSection string `json:"print_section,omitempty"`
	//PrintPage    string `json:"print_page,omitempty"`
	//Source       string `json:"source"`
	//Multimedia     []*MultiMediaData `json:"multimedia"`
	Headline *HeadLine `json:"headline"`
	//Keywords       []*Keyword        `json:"keywords"`
	PubDate      string `json:"pub_date"`
	DocumentType string `json:"document_type"`
	//NewsDesk       string  `json:"news_desk"`
	//SectionName    string  `json:"section_name"`
	//SubsectionName string  `json:"subsection_name,omitempty"`
	//Byline         *ByLine `json:"byline"`
	TypeOfMaterial string `json:"type_of_material"`
	//Id             string  `json:"_id"`
	WordCount int `json:"word_count"`
	//Uri            string  `json:"uri"`
}

func (doc *Doc) GetWebUrl() string {
	if doc != nil {
		return doc.WebUrl
	}
	return ""
}

func (doc *Doc) GetHeadline() *HeadLine {
	if doc != nil {
		return doc.Headline
	}
	return nil
}

type MultiMediaData struct {
	Rank     int     `json:"rank"`
	Subtype  string  `json:"subtype"`
	Caption  string  `json:"caption"`
	Credit   string  `json:"credit"`
	Type     string  `json:"type"`
	Url      string  `json:"url"`
	Height   int     `json:"height"`
	Width    int     `json:"width"`
	Legacy   *Legacy `json:"legacy"`
	SubType  string  `json:"subType"`
	CropName string  `json:"crop_name"`
}

type Legacy struct {
	Xlarge          string `json:"xlarge,omitempty"`
	Xlargewidth     int    `json:"xlargewidth,omitempty"`
	Xlargeheight    int    `json:"xlargeheight,omitempty"`
	Thumbnail       string `json:"thumbnail,omitempty"`
	Thumbnailwidth  int    `json:"thumbnailwidth,omitempty"`
	Thumbnailheight int    `json:"thumbnailheight,omitempty"`
	Widewidth       int    `json:"widewidth,omitempty"`
	Wideheight      int    `json:"wideheight,omitempty"`
	Wide            string `json:"wide,omitempty"`
}

type HeadLine struct {
	Main string `json:"main"`
	//Kicker        string `json:"kicker"`
	//ContentKicker string `json:"content_kicker"`
	PrintHeadline string `json:"print_headline"`
	//Name          string `json:"name"`
	//Seo           string `json:"seo"`
	//Sub           string `json:"sub"`
}

func (headLine *HeadLine) GetHeadLine() string {
	if headLine != nil {
		return headLine.PrintHeadline
	}

	return ""
}

type Keyword struct {
	Name  string `json:"name"`
	Value string `json:"value"`
	Rank  int    `json:"rank"`
	Major string `json:"major"`
}

type ByLine struct {
	Original     string        `json:"original"`
	Person       []*PersonData `json:"person"`
	Organization *string       `json:"organization"`
}
type PersonData struct {
	Firstname    string `json:"firstname"`
	Middlename   string `json:"middlename"`
	Lastname     string `json:"lastname"`
	Qualifier    string `json:"qualifier"`
	Title        string `json:"title"`
	Role         string `json:"role"`
	Organization string `json:"organization"`
	Rank         int    `json:"rank"`
}
