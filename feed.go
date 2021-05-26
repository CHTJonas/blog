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
	posts := b.GetPosts()
	feed.Items = make([]*feeds.Item, len(posts))
	for i, p := range posts {
		if p.Draft {
			continue
		}
		feed.Items[i] = &feeds.Item{
			Id:      p.UUID,
			Title:   p.Title,
			Link:    p.GetLink(),
			Created: p.Created,
			Updated: p.Updated,
			Author:  p.GetAuthor(),
			Content: p.GetHTML(),
		}
	}
	return feed.ToRss()
}
