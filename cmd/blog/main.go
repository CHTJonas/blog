package main

import (
	"github.com/CHTJonas/blog"
)

var CHTJonasBlog = &blog.Blog{
	URL:         "https://blog.charliejonas.co.uk",
	Title:       "CHTJonas' Blog",
	Description: "todo",
	OwnerName:   "Charlie Jonas",
	OwnerEmail:  "charlie@charliejonas.co.uk",
	YearCreated: 2021,
}

func main() {
	CHTJonasBlog.OpenDB("db")
	defer CHTJonasBlog.CloseDB()
	CHTJonasBlog.SeedPosts()
	CHTJonasBlog.Serve("localhost:8182")
}
