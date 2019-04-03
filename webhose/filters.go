package webhose

// News news and blog articles
type News struct {
	Posts            []Posts `json:"posts"`
	TotalResults     int     `json:"total_results"`
	MoreResultsAvail int     `json:"moreResultsAvailable"`
	Next             string  `json:"next"`
	RequestsLeft     int     `json:"requestsLeft"`
}

// Posts articles fetched
type Posts struct {
	Thread         Thread   `json:"thread"`
	UUID           string   `json:"uuid"`
	URL            string   `json:"url"`
	OrdInThread    int      `json:"ord_in_thread"`
	Author         string   `json:"author"`
	Published      string   `json:"published"`
	Title          string   `json:"title"`
	Text           string   `json:"text"`
	HighlightText  string   `json:"highlightText"`
	HighlightTitle string   `json:"highlightTitle"`
	Language       string   `json:"language"`
	ExternalLinks  []string `json:"external_links"`
	Entities       Entities `json:"entities"`
}

// Entities entities extracted
type Entities struct {
	Persons       []Entity
	Organizations []Entity
	Locations     []Entity
}

// Entity entity extracted
type Entity struct {
	Name      string `json:"name"`
	Sentiment string `json:"sentiment"`
}

// Thread details on thread
type Thread struct {
	UUID              string   `json:"uuid"`
	URL               string   `json:"url"`
	SiteFull          string   `json:"site_full"`
	Site              string   `json:"site"`
	SiteSection       string   `json:"site_section"`
	SiteCategories    []string `json:"site_categories"`
	SectionTitle      string   `json:"section_title"`
	Title             string   `json:"title"`
	TitleFull         string   `json:"title_full"`
	Published         string   `json:"published"`
	RepliesCount      int      `json:"replies_count"`
	ParticipantsCount int      `json:"participants_count"`
	SiteType          string   `json:"site_type"`
	Country           string   `json:"country"`
	SpamScore         int      `json:"spam_score"`
	MainImage         string   `json:"main_image"`
	PerformanceScore  int      `json:"performance_score"`
	DomainRank        int      `json:"domain_rank"`
	Social            Social   `json:"social"`
}

// Social social interactions
type Social struct {
	Facebook    Facebook
	Gplus       Shares
	Pinterest   Shares
	LinkedIn    Shares
	StumbleUpon Shares
	Vk          Shares
}

// Facebook facebook data
type Facebook struct {
	Likes    int `json:"likes"`
	Comments int `json:"comments"`
	Shares   int `json:"shares"`
}

// Shares number of social shares
type Shares struct {
	Shares int `json:"shares"`
}
