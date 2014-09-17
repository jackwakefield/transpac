package main

import (
	"fmt"
	"hash/fnv"
	"io/ioutil"
	"path"
	"time"

	"github.com/spf13/viper"
)

type cacheItem struct {
	location string
	filename string
}

func newCacheItem(location string) *cacheItem {
	// create a hash for the specified location to be used as the file name
	hash := fnv.New64a()
	hash.Write([]byte(location))

	return &cacheItem{
		location: location,
		filename: fmt.Sprintf("%x", hash.Sum64()),
	}
}

func (cache *cacheItem) exists() bool {
	return fileExists(cache.path())
}

func (cache *cacheItem) outOfDate() bool {
	if lastModified, err := fileModifiedTime(cache.path()); err == nil {
		// determine whether the last time the file was modified was longer than
		// the configured cache length
		return time.Now().Sub(lastModified).Seconds() > float64(viper.GetInt(cacheLengthKey))
	}

	return true
}

func (cache *cacheItem) path() string {
	return path.Join(viper.GetString(cacheDirectoryKey), cache.filename)
}

func (cache *cacheItem) save(data []byte) error {
	return ioutil.WriteFile(cache.path(), data, 0644)
}
