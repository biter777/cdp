// Code generated by cdpgen. DO NOT EDIT.

package cachestorage

// CacheID Unique identifier of the Cache object.
type CacheID string

// DataEntry Data entry.
type DataEntry struct {
	RequestURL      string   `json:"requestURL"`      // Request URL.
	ResponseTime    float64  `json:"responseTime"`    // Number of seconds since epoch.
	ResponseHeaders []Header `json:"responseHeaders"` // Response headers
}

// Cache Cache identifier.
type Cache struct {
	CacheID        CacheID `json:"cacheId"`        // An opaque unique id of the cache.
	SecurityOrigin string  `json:"securityOrigin"` // Security origin of the cache.
	CacheName      string  `json:"cacheName"`      // The name of the cache.
}

// Header
type Header struct {
	Name  string `json:"name"`  // No description.
	Value string `json:"value"` // No description.
}

// CachedResponse Cached response
type CachedResponse struct {
	Body string `json:"body"` // Entry content, base64-encoded.
}
