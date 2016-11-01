package main

import (
	"fmt"
	"log"
	"sync/atomic"

	"github.com/boltdb/bolt"
	"github.com/calebamiles/github-client/client"
)

func main() {
	var numPagesInCache uint32

	cacheFile := "/tmp/test-github-client-caching.boltdb"
	c, err := client.New("calebamiles", "test-caching-repo", "", cacheFile)
	if err != nil {
		panic(err)
	}

	_, err = c.FetchPullRequests()
	if err != nil {
		panic(err)
	}

	_, err = c.FetchPullRequests()
	if err != nil {
		panic(err)
	}

	err = c.Done()
	if err != nil {
		panic(err)
	}

	db, err := bolt.Open(cacheFile, 0600, nil)
	if err != nil {
		panic(err)
	}

	db.View(func(tx *bolt.Tx) error {
		pages := tx.Bucket([]byte("pages"))
		pages.ForEach(func(k, v []byte) error {
			atomic.AddUint32(&numPagesInCache, 1)
			log.Println("cahed key: " + string(k))
			return nil
		})

		return nil
	})

	fmt.Printf("there are %d pages in the cache", numPagesInCache)
}
