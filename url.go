package url_short

type URL struct {
	ID    string
	Value string
}

type URLReq struct {
	URL string `json:"url"`
}

type URLRes struct {
	URL string `json:"url"`
}
