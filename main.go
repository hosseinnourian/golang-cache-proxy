package main

import (
	"cache-proxy/cache"
	"cache-proxy/proxy"
	"context"
	"flag"
	"fmt"
	"log"
)

func main() {
	origin := flag.String("origin", "", "URL to fetch")
	clearCache := flag.Bool("clear-cache", false, "Clear all cache")
	flag.Parse()


	redisClient := cache.NewRedisClient()
	redisCache := cache.NewRedisCache(redisClient)


	if *clearCache {
		err := redisCache.FlushCache(context.TODO())
		if err != nil {
			log.Fatalf("Failed to clear cache: %v", err)
		}
		return
	}

	// Ensure the URL is provided
	if *origin == "" {
		log.Fatal("origin must contain a URL")
	}

	// Initialize Redis client and cache


	// Initialize the proxy with the Redis cache
	cacheProxy := proxy.NewCacheProxy(redisCache)

	// Make the HTTP call (either cached or fetched from origin)
	response, err := cacheProxy.HttpCall(*origin)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	fmt.Println(response)
}
