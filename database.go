package blog

import (
	"github.com/dgraph-io/badger"
)

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

func (b *Blog) Read(key string) ([]byte, error) {
	var data []byte
	outerErr := b.database.View(func(txn *badger.Txn) error {
		item, innerErr := txn.Get([]byte(key))
		if innerErr != nil {
			return innerErr
		}
		return item.Value(func(val []byte) error {
			data = append([]byte{}, val...)
			return nil
		})
	})
	return data, outerErr
}

func (b *Blog) Write(key string, data []byte) error {
	return b.database.Update(func(txn *badger.Txn) error {
		return txn.Set([]byte(key), data)
	})
}

func (b *Blog) Scan() (map[string][]byte, error) {
	var data map[string][]byte
	err := b.database.View(func(txn *badger.Txn) error {
		it := txn.NewIterator(badger.DefaultIteratorOptions)
		defer it.Close()
		prefix := []byte("posts/")
		for it.Seek(prefix); it.ValidForPrefix(prefix); it.Next() {
			item := it.Item()
			k := item.Key()
			v, innerErr := item.ValueCopy(nil)
			if innerErr != nil {
				return innerErr
			}
			data[string(k)] = v
		}
		return nil
	})
	return data, err
}
