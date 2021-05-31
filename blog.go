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

func (b *Blog) GetPosts() []*Post {
	// TODO
	posts := make([]*Post, 4)
	posts[0] = NewPost("First Test", "test-one", "This is the first test.")
	posts[1] = NewPost("Second Test", "test-two", "This is the second test.")
	posts[2] = NewPost("Third Test", "test-three", "This is the third test.")
	posts[3] = NewPost("Markdown Test", "test-md", markdown)
	return posts
}

const markdown = `## Section 1

Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Purus semper eget duis at tellus. Lorem donec massa sapien faucibus et molestie ac. Vel fringilla est ullamcorper eget nulla facilisi. Sit amet nulla facilisi morbi tempus iaculis. Gravida cum sociis natoque penatibus et. Fames ac turpis egestas maecenas pharetra. Justo eget magna fermentum iaculis eu non diam. Ac felis donec et odio pellentesque diam volutpat commodo sed. Feugiat nisl pretium fusce id velit. Dictumst vestibulum rhoncus est pellentesque elit ullamcorper dignissim cras tincidunt.

## Section 2

Diam volutpat commodo sed egestas. Pharetra massa massa ultricies mi quis. Nunc scelerisque viverra mauris in aliquam sem fringilla ut morbi. Integer malesuada nunc vel risus commodo viverra maecenas accumsan lacus. Posuere ac ut consequat semper. Odio tempor orci dapibus ultrices in iaculis nunc sed augue. Morbi enim nunc faucibus a pellentesque sit. Ut sem nulla pharetra diam sit amet nisl suscipit adipiscing. Iaculis nunc sed augue lacus viverra vitae congue eu consequat. Amet nisl suscipit adipiscing bibendum est ultricies integer. Amet mattis vulputate enim nulla aliquet porttitor lacus. Mauris a diam maecenas sed enim ut sem. Feugiat vivamus at augue eget arcu dictum varius duis at. Nisi vitae suscipit tellus mauris a diam maecenas sed. Vel fringilla est ullamcorper eget nulla facilisi etiam dignissim. Blandit turpis cursus in hac. Sed sed risus pretium quam vulputate dignissim suspendisse in est.

## Section 3

Commodo elit at imperdiet dui accumsan sit amet. Enim neque volutpat ac tincidunt vitae semper quis lectus nulla. Massa enim nec dui nunc. Molestie nunc non blandit massa enim. Feugiat nibh sed pulvinar proin. Tristique senectus et netus et malesuada. At tempor commodo ullamcorper a lacus vestibulum. Mattis molestie a iaculis at erat pellentesque. Ultrices sagittis orci a scelerisque. Amet volutpat consequat mauris nunc congue nisi. Amet justo donec enim diam vulputate ut.

## Section 4

Vulputate eu scelerisque felis imperdiet. At erat pellentesque adipiscing commodo elit at imperdiet. Neque gravida in fermentum et. Mi eget mauris pharetra et ultrices neque. Elementum curabitur vitae nunc sed velit dignissim. Ac turpis egestas maecenas pharetra. Lectus nulla at volutpat diam ut venenatis tellus in. Duis ut diam quam nulla. Ut etiam sit amet nisl. Laoreet suspendisse interdum consectetur libero id faucibus nisl tincidunt eget. Leo in vitae turpis massa sed. Eu tincidunt tortor aliquam nulla. Elementum eu facilisis sed odio. Et ligula ullamcorper malesuada proin libero. Porta lorem mollis aliquam ut porttitor leo a diam. Posuere morbi leo urna molestie at. Augue mauris augue neque gravida in. Urna id volutpat lacus laoreet non curabitur gravida arcu. Gravida rutrum quisque non tellus orci ac auctor augue.

## Section 5

Lectus proin nibh nisl condimentum id venenatis a. Semper quis lectus nulla at volutpat. Hendrerit gravida rutrum quisque non tellus. Ut diam quam nulla porttitor. Sem et tortor consequat id porta. Odio eu feugiat pretium nibh ipsum consequat. Et magnis dis parturient montes nascetur ridiculus mus mauris. Turpis cursus in hac habitasse platea dictumst. Odio ut sem nulla pharetra. Viverra tellus in hac habitasse platea dictumst vestibulum rhoncus. Pharetra diam sit amet nisl suscipit. Tellus id interdum velit laoreet id donec ultrices. Condimentum vitae sapien pellentesque habitant morbi tristique senectus et. Pharetra vel turpis nunc eget lorem dolor sed viverra ipsum.`
