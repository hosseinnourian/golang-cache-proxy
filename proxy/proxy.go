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

func (p *CacheProxy) HttpCall(origin string, fresh bool) (string, error) {
	ctx := context.Background()


	if !fresh {
		value, err := p.Cache.Get(ctx, origin)
		if err == nil {
			fmt.Println("Read from cache...")
			return value, nil
		}

		// Log any Redis errors
		if err.Error() != fmt.Sprintf("cache miss for key: %s", origin) {
			fmt.Println("Redis error:", err)
		} else {
			fmt.Println("Cache miss. Fetching from origin...")
		}
	} else {
		fmt.Println("Fresh mode enabled. Bypassing cache...")
	}

	// Make the HTTP request (shared by both freshCopy and cache miss cases)
	return p.fetchAndCacheResponse(ctx, origin)

}
func (p *CacheProxy) fetchAndCacheResponse(ctx context.Context, origin string) (string, error) {
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

	// Cache the response with an expiration of 60 minutes
	expiration := 60 * time.Minute
	err = p.Cache.Set(ctx, origin, stringResponse, expiration)
	if err != nil {
		fmt.Println("Error caching response:", err)
	}

	return stringResponse, nil
}