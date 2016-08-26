package common

import(
  "net/http"
  "net/http/httptest"
  "github.com/stretchr/testify/assert"
  "testing"
  "github.com/nu7hatch/gouuid"
)

// httptest.NewRecorder isn't compatible with NewRequestWriter
// Need to figure this out
// func TestRequestTimerHandler(t *testing.T) {
//   req, err := http.NewRequest("GET", "/level-three-rest", nil)
//   if err != nil {
//       t.Fatal(err)
//   }
//
//   rr := httptest.NewRecorder()
//   handler := testHandler(rr, req)
//
//   RequestTimerHandler(handler).ServeHTTP(rr, req)
//   assert.Equal(t, rr.Code, http.StatusOK)
// }

func TestTraceIdInHeader(t *testing.T) {
  req, err := http.NewRequest("GET", "/level-three-rest", nil)
  h := "test-header"
  req.Header.Set(TraceHeader, h)
  if err != nil {
      t.Fatal(err)
  }

  rr := httptest.NewRecorder()
  handler := testHandler(rr, req)

  TraceHandler(handler).ServeHTTP(rr, req)
  assert.Equal(t, rr.Header().Get(TraceHeader), h)
}

func TestNoTraceIdInHeader(t *testing.T) {
  req, err := http.NewRequest("GET", "/level-three-rest", nil)
  if err != nil {
      t.Fatal(err)
  }

  rr := httptest.NewRecorder()
  handler := testHandler(rr, req)

  TraceHandler(handler).ServeHTTP(rr, req)
  h := rr.Header().Get(TraceHeader)
  assert.NotEmpty(t, h)
  u, err := uuid.ParseHex(h)
  assert.NotNil(t, u)
  assert.Nil(t, err)
}

// func TestResponseWriterHandler(t *testing.T) {
//   req, err := http.NewRequest("GET", "/level-three-rest", nil)
//   if err != nil {
//       t.Fatal(err)
//   }
//
//   rr := httptest.NewRecorder()
//   handler := testHandler(rr, req)
//
//   ResponseWriterHandler(handler).ServeHTTP(rr, req)
//   assert.IsType(t, ResponseWriter, rr)
// }

func testHandler(w http.ResponseWriter, r *http.Request) http.Handler {
  fn := func (w http.ResponseWriter, r *http.Request)  {
    w.WriteHeader(http.StatusOK)
  }
  return http.HandlerFunc(fn)
}
