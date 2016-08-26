package common

import (
  "net/http"
  "time"
  "github.com/nu7hatch/gouuid"
)

const(
  // The default trace header is "X-Request-ID" which can be used to correlate
  // requests between the client and server. This is added to every log message
  // when the logger package is used.
  TraceHeader = "X-Request-ID"
)

// Trace handler ensures that the trace header is present on all requests even
// if the client did not present one. In the case the client does not present
// one, then a UUID will be generated.
func TraceHandler(next http.Handler) http.Handler  {
  fn := func(w http.ResponseWriter, r *http.Request)  {
    if r.Header.Get(TraceHeader) == "" {
      u, _ := uuid.NewV4()
      r.Header.Set(TraceHeader, u.String())
    }
    w.Header().Set(TraceHeader,r.Header.Get(TraceHeader))
    next.ServeHTTP(w, r)
  }
  return http.HandlerFunc(fn)
}

// Logs the time taken for the entire call to be completed.
func RequestTimerHandler(next http.Handler) http.Handler {
  fn := func (w http.ResponseWriter, r *http.Request)  {
    time.Now()
  }
  return http.HandlerFunc(fn)
}
