package source

import "github.com/miiren/mbox/cache"

var Cache *cache.MCache

func InitCache() {
	Cache = cache.New(10)
}
