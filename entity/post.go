package entity

type Post struct {
	ID    string
	Title string
	Text  string
}

func CreatePost(id, title, text string) Post {
	return Post{
		ID:    id,
		Title: title,
		Text:  text,
	}
}
