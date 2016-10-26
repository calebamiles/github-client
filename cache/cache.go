package cache

import (
	"crypto/md5"
	"fmt"
	"io"
	"time"

	"github.com/boltdb/bolt"
)

type PageCache interface {
	KeyForPage(pageURL string) (cacheKey string, exists bool)
	FetchPageByKey(cacheKey string) (page []byte, err error)
	AddPage(pageURL string, cacheKey string, page []byte) error
	Open() error
	Close() error
}

// New returns a new PageCache or an error relating to opening the datastore
func New(datastorePath string) PageCache {

	return nil
}

const (
	pagesBucket = "pages"
	keysBucket  = "keys"
)

type pageCache struct {
	dbPath string
	db     *bolt.DB
}

func (c *pageCache) Open() error {
	db, err := bolt.Open(c.dbPath, 0600, &bolt.Options{Timeout: 10 * time.Second})
	if err != nil {
		return err
	}

	db.Update(func(tx *bolt.Tx) error {
		var updateErr error
		_, updateErr = tx.CreateBucketIfNotExists([]byte(pagesBucket))
		if updateErr != nil {
			return fmt.Errorf("create pages bucket: %s", updateErr)
		}

		_, updateErr = tx.CreateBucketIfNotExists([]byte(keysBucket))
		if updateErr != nil {
			return fmt.Errorf("create keys bucket: %s", updateErr)
		}

		return nil
	})

	c.db = db

	return nil
}

func (c *pageCache) Close() error {
	return c.db.Close()
}

func (c *pageCache) AddPage(pageURL string, cacheKey string, page []byte) error {
	c.db.Update(func(tx *bolt.Tx) error {
		keys := tx.Bucket([]byte(keysBucket))
		pages := tx.Bucket([]byte(pagesBucket))

		hash := md5.New()
		_, err := io.WriteString(hash, pageURL)
		if err != nil {
			return err
		}

		err = keys.Put(hash.Sum(nil), []byte(cacheKey))

		return err
	})

}
