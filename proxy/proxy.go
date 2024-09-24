package proxy

import (
	"cache-proxy/cache"
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

type CacheProxy struct {
	Cache *cache.RedisCache
}

func NewCacheProxy(redisCache *cache.RedisCache) *CacheProxy {
	return &CacheProxy{Cache: redisCache}
}

func (p *CacheProxy) HttpCall(origin string) (string, error) {

	ctx := context.Background()

	value, err := p.Cache.Get(ctx, origin)
	if err == nil {
		fmt.Println("Read from cache...")
		return value, nil
	}
	if err != nil && err.Error() != fmt.Sprintf("cache miss for key: %s", origin) {
		fmt.Println("Redis error:", err)
	}

	fmt.Println("Cache miss. Fetching from origin...")

	response, err := http.Get(origin)
	if err != nil {
		return "", fmt.Errorf("error fetching origin %s: %v", origin, err)
	}
	defer response.Body.Close()

	// Read response body
	byteBody, err := io.ReadAll(response.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response body: %v", err)
	}
	stringResponse := string(byteBody)

	// Cache the response with an expiration of 5 minutes
	expiration := 60 * time.Minute
	err = p.Cache.Set(ctx, origin, stringResponse, expiration)
	if err != nil {
		fmt.Println("Error caching response:", err)
	}

	return stringResponse, nil

}
