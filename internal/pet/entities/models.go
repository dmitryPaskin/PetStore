package entities

type Pet struct {
	Id       int       `json:"id"`
	Category *Category `json:"category"`
	Name     string    `json:"name"`
	Tags     []*Tag    `json:"tags"`
	Status   string    `json:"status"`
}

type Category struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Tag struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
