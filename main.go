package main

import (
	"cache-proxy/cache"
	"cache-proxy/proxy"
	"flag"
	"fmt"
	"log"
)

func main() {
	origin := flag.String("origin", "", "URL to fetch")
	flag.Parse()

	// Ensure the URL is provided
	if *origin == "" {
		log.Fatal("origin must contain a URL")
	}

	// Initialize Redis client and cache
	redisClient := cache.NewRedisClient()
	redisCache := cache.NewRedisCache(redisClient)

	// Initialize the proxy with the Redis cache
	cacheProxy := proxy.NewCacheProxy(redisCache)

	// Make the HTTP call (either cached or fetched from origin)
	response, err := cacheProxy.HttpCall(*origin)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	fmt.Println(response)
}
