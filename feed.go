package blog

import (
	"time"

	"github.com/gorilla/feeds"
)

func (b *Blog) GetRSS() (string, error) {
	feed := &feeds.Feed{
		Link:        b.GetLink(),
		Title:       b.Title,
		Description: b.Description,
		Author:      b.GetAuthor(),
		Created:     time.Now().UTC(),
		Copyright:   b.GetCopyright(),
	}
	posts, err := b.ListPosts()
	if err != nil {
		return "", err
	}
	feed.Items = make([]*feeds.Item, len(*posts))
	i := 0
	for slug, p := range *posts {
		if p.Draft {
			continue
		}
		feed.Items[i] = &feeds.Item{
			Id:      p.UUID,
			Title:   p.Title,
			Link:    p.GetLink(slug),
			Created: p.Created,
			Updated: p.Updated,
			Author:  p.GetAuthor(),
			Content: p.GetHTML(),
		}
		i++
	}
	return feed.ToRss()
}
