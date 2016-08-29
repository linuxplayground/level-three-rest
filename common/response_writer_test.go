package common

import (
	"net/http"
	"net/http/httptest"
  "github.com/stretchr/testify/assert"
  "testing"
)

func TestWriteHeader(t *testing.T) {
	w := NewResponseWriter(httptest.NewRecorder())
  w.WriteHeader(http.StatusInternalServerError)
  assert.Equal(t, w.Status(), http.StatusInternalServerError)
}

func TestWriteOnce(t *testing.T) {
	w := NewResponseWriter(httptest.NewRecorder())
	b := []byte("ok")
  l, _ := w.Write(b)
  assert.Equal(t, l, len(b))
}


func TestWriteMultipleTime(t *testing.T) {
	w := NewResponseWriter(httptest.NewRecorder())
	b := []byte("ok")
	c := []byte("Another ok!")
  l, _ := w.Write(b)
	sz, _ := w.Write(c)
  assert.Equal(t, l+sz, len(b) + len(c))
}
