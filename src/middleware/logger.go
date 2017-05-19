package middleware

var (
	Logger ILogger
)

type ILogger interface {
	Println(args ...interface{})
	Printf(format string, args ...interface{})
	LogWithFields(message string, fields map[string]interface{})
	Error(args ...interface{})
	Errorf(format string, args ...interface{})
}

func SetLogger(l ILogger) {
	Logger = l
}