package cache

import (
	"io/ioutil"
	"os"
	"strings"
	"time"
)

const Location = "/var/cache/"

type CacheItem struct {
	TTL int64
	Key string
}

func NewCache(endpoint string, params ...string) CacheItem {
	cacheName := endpoint + "_" + strings.Join(params, "_")
	return CacheItem{
		Key: Location + cacheName,
		TTL: 3600, // 1 soat
	}
}

func (c CacheItem) Get() (bool, string) {
	stats, err := os.Stat(c.Key)
	if err != nil {
		return false, ""
	}
	age := time.Now().Unix() - stats.ModTime().Unix()
	if age <= c.TTL {
		cache, err := ioutil.ReadFile(c.Key)
		if err == nil {
			return true, string(cache)
		}
	}
	return false, ""
}

func (c CacheItem) Set(data string) bool {
	err := ioutil.WriteFile(c.Key, []byte(data), 0644)
	return err == nil
}

func (c CacheItem) Clear() bool {
	err := os.Remove(c.Key)
	return err == nil
}
