package common

import (
  "net/http"
  "time"
  "github.com/nu7hatch/gouuid"
  "log"
  "github.com/bhavikkumar/level-three-rest/logger"
)

const(
  // The default trace header is "X-Request-ID" which can be used to correlate
  // requests between the client and server. This is added to every log message
  // when the logger package is used.
  TraceHeader = "X-Request-ID"
)

// Trace handler ensures that the trace id is available even
// if the client did not present one. In the case the client does not present
// one, then a UUID will be generated. The TraceId is used by the logger
// package.
func TraceHandler(next http.Handler) http.Handler  {
  fn := func(w http.ResponseWriter, r *http.Request)  {
    logger.TraceId = r.Header.Get(TraceHeader)
    if logger.TraceId == "" {
      u, err := uuid.NewV4()
      if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        return
      }
      logger.TraceId = u.String()
    }
    w.Header().Set(TraceHeader, logger.TraceId)
    next.ServeHTTP(w, r)
  }
  return http.HandlerFunc(fn)
}

// Logs the time taken for the entire call to be completed.
func RequestTimerHandler(next http.Handler) http.Handler {
  fn := func (w http.ResponseWriter, r *http.Request)  {
    time.Now()
    // Use our ResponseWriter wrapper.
    w = NewResponseWriter(w)
    next.ServeHTTP(w,r)
    res := w.(ResponseWriter)
    log.Println("Size:", res.Size(),"Status:",res.Status())
  }
  return http.HandlerFunc(fn)
}

func RecoveryHandler(next http.Handler) http.Handler  {
  fn := func(w http.ResponseWriter, r *http.Request) {
    defer func() {
       if err := recover(); err != nil {
         http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
       }
     }()
     next.ServeHTTP(w, r)
   }
  return http.HandlerFunc(fn)
}
