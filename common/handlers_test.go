package common

import(
  "net/http"
  "net/http/httptest"
  "github.com/stretchr/testify/assert"
  "testing"
  "github.com/nu7hatch/gouuid"
)

func TestRequestTimerHandler(t *testing.T) {
  req, err := http.NewRequest("GET", "/level-three-rest", nil)
  if err != nil {
      t.Fatal(err)
  }

  rr := httptest.NewRecorder()
  handler := testHandler(rr, req)

  RequestTimerHandler(handler).ServeHTTP(rr, req)
  assert.Equal(t, rr.Code, http.StatusOK)
}

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

func TestRecoveryHandlerNoError(t *testing.T) {
  req, err := http.NewRequest("GET", "/level-three-rest", nil)
  if err != nil {
      t.Fatal(err)
  }

  rr := httptest.NewRecorder()
  handler := testHandler(rr, req)

  RecoveryHandler(handler).ServeHTTP(rr, req)
  assert.Equal(t, rr.Code, http.StatusOK)
}

func TestRecoveryHandlerError(t *testing.T) {
  req, err := http.NewRequest("GET", "/level-three-rest", nil)
  if err != nil {
      t.Fatal(err)
  }

  rr := httptest.NewRecorder()
  handler := testPanicHandler(rr, req)

  RecoveryHandler(handler).ServeHTTP(rr, req)
  assert.Equal(t, rr.Code, http.StatusInternalServerError)
}

func testHandler(w http.ResponseWriter, r *http.Request) http.Handler {
  fn := func (w http.ResponseWriter, r *http.Request)  {
    w.WriteHeader(http.StatusOK)
  }
  return http.HandlerFunc(fn)
}

func testPanicHandler(w http.ResponseWriter, r *http.Request) http.Handler {
  fn := func (w http.ResponseWriter, r *http.Request)  {
    panic("oh no!")
  }
  return http.HandlerFunc(fn)
}
