package scanblock

import "net/http"

// ResponseWriter is used to wrap given response writers.
type ResponseWriter struct {
	http.ResponseWriter
	// HTTP status codes to block, listed as start->end pair, e.g. the 400,499 for 400->401,410->499 inclusive
	BlockStatusCodeStart int
	BlockStatusCodeEnd   int
	cacheEntry           *CacheEntry
}

// WriteHeader adds custom handling to the wrapped WriterHeader method.
func (rw *ResponseWriter) WriteHeader(code int) {
	// If the request returns with a 4xx code, increase the scan requests counter.
	if code >= rw.BlockStatusCodeStart && code <= rw.BlockStatusCodeEnd {
		rw.cacheEntry.ScanRequests.Add(1)
	}

	// Continue with the original method.
	rw.ResponseWriter.WriteHeader(code)
}
