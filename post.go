package blog

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/feeds"
	"github.com/russross/blackfriday"
)

type Post struct {
	UUID        string
	Title       string
	Created     time.Time
	Updated     time.Time
	AuthorName  string
	AuthorEmail string
	Content     string
	Tags        []string
	Draft       bool
	NoIndex     bool
}

func NewDraft(title, content string) *Post {
	p := NewPost(title, content)
	p.Draft = true
	return p
}

func NewPost(title, content string) *Post {
	now := time.Now()
	return &Post{
		UUID:        uuid.New().String(),
		Title:       title,
		Created:     now,
		Updated:     now,
		AuthorName:  "Charlie Jonas",
		AuthorEmail: "charlie@charliejonas.co.uk,",
		Content:     content,
	}
}

func (p *Post) GetLink(slug string) *feeds.Link {
	return &feeds.Link{
		Href: fmt.Sprintf("https://blog.charliejonas.co.uk/post/%s", slug),
	}
}

func (p *Post) GetAuthor() *feeds.Author {
	return &feeds.Author{
		Name:  p.AuthorName,
		Email: p.AuthorEmail,
	}
}

func (p *Post) GetHTML() string {
	return string(blackfriday.MarkdownCommon([]byte(p.Content)))
}
