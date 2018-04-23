package blog

// Author model
type Author struct {
	ID       int
	Name     string
	Email    string
	username string
}

// Tag model
type Tag struct {
	ID   int
	Name string
	Slug string
}

// Post Model
type Post struct {
	ID     int     `json:"ID"`
	Author *Author `json:"author"`
	Tags   []*Tag  `json:"tags"`
}
