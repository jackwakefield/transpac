package main

import (
	"fmt"
	"hash/fnv"
	"io/ioutil"
	"time"
)

type cacheItem struct {
	location string
	hash     uint64
}

func newCacheItem(location string) *cacheItem {
	// create a hash for the specified location to be used as the file name
	hash := fnv.New64a()
	hash.Write([]byte(location))

	return &cacheItem{
		location: location,
		hash:     hash.Sum64(),
	}
}

func (cache *cacheItem) exists() bool {
	return fileExists(cache.path())
}

func (cache *cacheItem) outOfDate() bool {
	if lastModified, err := fileModifiedTime(cache.path()); err == nil {
		// determine whether the last time the file was modified was longer than
		// the configured cache length
		return time.Now().Sub(lastModified).Seconds() > float64(*cacheLength)
	}

	return true
}

func (cache *cacheItem) path() string {
	return fmt.Sprintf("%s/%x", *cacheDirectory, cache.hash)
}

func (cache *cacheItem) save(data []byte) error {
	return ioutil.WriteFile(cache.path(), data, 0644)
}
