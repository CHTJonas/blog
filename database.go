package blog

import (
	"bytes"
	"encoding/gob"
	"strings"

	"github.com/dgraph-io/badger"
)

const databasePrefix = "posts/"

func (b *Blog) OpenDB(path string) error {
	opts := badger.DefaultOptions(path)
	db, err := badger.Open(opts)
	if err != nil {
		return err
	}
	b.database = db
	return nil
}

func (b *Blog) CloseDB() {
	b.database.Close()
}

func (b *Blog) LoadPost(slug string) (*Post, error) {
	var post Post
	outerErr := b.database.View(func(txn *badger.Txn) error {
		key := []byte(databasePrefix + slug)
		item, innerErr := txn.Get(key)
		if innerErr != nil {
			return innerErr
		}
		return item.Value(func(data []byte) error {
			buf := bytes.NewBuffer(data)
			dec := gob.NewDecoder(buf)
			return dec.Decode(&post)
		})
	})
	return &post, outerErr
}

func (b *Blog) SavePost(slug string, post *Post) error {
	return b.database.Update(func(txn *badger.Txn) error {
		var buf bytes.Buffer
		enc := gob.NewEncoder(&buf)
		err := enc.Encode(post)
		if err != nil {
			return err
		}
		key := []byte(databasePrefix + slug)
		return txn.Set(key, buf.Bytes())
	})
}

func (b *Blog) ListPosts() (*map[string]*Post, error) {
	posts := make(map[string]*Post)
	outerErr := b.database.View(func(txn *badger.Txn) error {
		it := txn.NewIterator(badger.DefaultIteratorOptions)
		defer it.Close()
		prefix := []byte(databasePrefix)
		for it.Seek(prefix); it.ValidForPrefix(prefix); it.Next() {
			var post Post
			item := it.Item()
			slug := strings.TrimPrefix(string(item.Key()), databasePrefix)
			if innerErr := item.Value(func(data []byte) error {
				buf := bytes.NewBuffer(data)
				dec := gob.NewDecoder(buf)
				return dec.Decode(&post)
			}); innerErr != nil {
				return innerErr
			}
			posts[slug] = &post
		}
		return nil
	})
	return &posts, outerErr
}
