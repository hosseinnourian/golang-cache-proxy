## Cache Proxy with Redis (Golang)
This is a simple cache proxy written in Go that caches HTTP requests using Redis. The proxy fetches and caches responses from URLs, and on subsequent requests, retrieves the responses from the cache, thus improving performance. You can also enable a fresh mode to bypass the cache or clear the entire cache using CLI flags.

## Features
Caching HTTP Responses: Responses to HTTP requests are stored in Redis with configurable expiration (default is 60 minutes).
Fresh Mode: Option to bypass the cache and make a fresh HTTP request.
Clear Cache: Option to clear all cached data stored in Redis.
Redis Integration: Utilizes Redis as the caching backend for storing and retrieving responses.

## Project Structure

.
├── cache/
│   └── redis.go    # Redis connection and cache interaction logic
├── proxy/
│   └── proxy.go    # Core proxy logic for handling HTTP requests and caching
└── main.go         # CLI entry point for the application


## Installation
# Prerequisites
    Go 1.16 or higher
    Redis installed and running locally (or change the configuration if Redis is hosted elsewhere)
  
## Clone the Repository

git clone https://github.com/your-username/cache-proxy.git
cd cache-proxy
Install Dependencies
There are no external Go dependencies aside from Redis, which is handled by the go-redis package.

You can install required Go packages via go get:

go get ./...
## Usage
You can run the proxy from the command line, specifying different flags depending on your needs.

1. Fetch URL and Cache Response
Fetches the response for a given URL and caches it in Redis. Subsequent requests to the same URL will be served from the cache.
go run main.go --origin="http://example.com"

2. Fresh Mode (Bypass Cache)
Forces a fresh HTTP request, bypassing the cache entirely. The response will still be cached for future use.
go run main.go --origin="http://example.com" --fresh

3. Clear Entire Cache
Clears all cached data stored in Redis.
go run main.go --clear-cache

## CLI Flags
--origin: (Required) The URL you want to fetch.
--clear-cache: Clears all cached data in Redis.
--fresh: Forces a fresh HTTP request, bypassing the cache.
