package client

// Logger is a minimal logging interface consumers can implement and pass via
// Config.Logger. Methods mirror common leveled printf-style logging.
type Logger interface {
	Debugf(format string, v ...interface{})
	Infof(format string, v ...interface{})
	Warnf(format string, v ...interface{})
	Errorf(format string, v ...interface{})
}
