package blog

import (
	"fmt"
	"strings"
)

func (b *Blog) GetSitemap() string {
	var builder strings.Builder
	builder.WriteString(`<?xml version="1.0" encoding="UTF-8"?>` + "\n")
	builder.WriteString(`<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">` + "\n")
	posts, err := b.ListPosts()
	if err != nil {
		panic(err)
	}
	for slug, post := range *posts {
		if post.Draft {
			continue
		}
		builder.WriteString("\t<url>\n")
		builder.WriteString(fmt.Sprintf("\t\t<loc>%s/%s</loc>\n", b.URL, slug))
		builder.WriteString(fmt.Sprintf("\t\t<lastmod>%s</lastmod>\n", post.Updated.Format("2006-01-02T15:04:05-07:00")))
		builder.WriteString("\t</url>\n")
	}
	builder.WriteString("</urlset>")
	return builder.String()
}
