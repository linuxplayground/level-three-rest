package common

import (
	"net/http"
)

// ResponseWriter is a wrapper around http.ResponseWriter that provides extra information about
// the response. It is recommended that middleware handlers use this construct to wrap a responsewriter
// if the functionality calls for it.
type ResponseWriter interface {
	http.ResponseWriter
	// Status returns the status code of the response or 200 if the response has
	// not been written (as this is the default response code in net/http)
	Status() int

	// Size returns the size of the response body.
	Size() int
}

// NewResponseWriter creates a ResponseWriter that wraps an http.ResponseWriter
func NewResponseWriter(w http.ResponseWriter) ResponseWriter {
	return &responseWriter{
		ResponseWriter: w,
	}
}

type responseWriter struct {
	http.ResponseWriter
	status      int
	size        int
}

func (w *responseWriter) WriteHeader(s int) {
	w.status = s
	w.ResponseWriter.WriteHeader(s)
}

func (w *responseWriter) Write(b []byte) (int, error) {
	size, err := w.ResponseWriter.Write(b)
	w.size += size
	return size, err
}

func (w *responseWriter) Status() int {
	return w.status
}

func (w *responseWriter) Size() int {
	return w.size
}
