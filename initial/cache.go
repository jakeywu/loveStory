package initial

import (
	"sc-git.com/beeApi/utils/cache"
)

func initCache() {
	cache.InitRedis()
}
