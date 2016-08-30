// Logger implements a custom logging mechanism which can be used within handlers
// Based on https://github.com/sadlil/gologger/blob/v0/gologger.go
package logger

// import(
//   "log"
//   "fmt"
//
// )
//
// const (
// 	CONSOLE string = "console"
// 	FILE string = "file"
// )
//
//
// The unique trace id to be logged for each log event. This allows us to trace entire
// end user request.
var TraceId string
//
// type Logger struct {
//   LoggerType string
//   log.Logger
// }
//
// func NewLogger(loggerType string)  {
//   l := log.New()
// }
