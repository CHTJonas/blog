package blog

import (
	"fmt"
	"time"

	"github.com/dgraph-io/badger"
	"github.com/gorilla/feeds"
)

type Blog struct {
	URL         string
	Title       string
	Description string
	OwnerName   string
	OwnerEmail  string
	YearCreated int
	database    *badger.DB
}

func (b *Blog) Serve(addr string) (*Server, error) {
	s := NewServer()
	s.b = b
	return s, s.Start(addr)
}

func (b *Blog) GetLink() *feeds.Link {
	return &feeds.Link{
		Href: b.URL,
	}
}

func (b *Blog) GetAuthor() *feeds.Author {
	return &feeds.Author{
		Name:  b.OwnerName,
		Email: b.OwnerEmail,
	}
}

func (b *Blog) GetCopyright() string {
	years := fmt.Sprint(b.YearCreated)
	nowYear := fmt.Sprint(time.Now().Year())
	if nowYear != years {
		years = fmt.Sprintf("%s-%s", years, nowYear)
	}
	return fmt.Sprintf("Copyright (c) %s %s", years, b.OwnerName)
}

// TODO remove this
func (b *Blog) GetPosts() []*Post {
	p := make([]*Post, 0, 20)
	posts, err := b.ListPosts()
	if err != nil {
		fmt.Println(err)
		return nil
	}
	for slug, post := range *posts {
		post.Slug = slug
		p = append(p, post)
	}
	return p
}
