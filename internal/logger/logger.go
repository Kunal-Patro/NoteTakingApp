package logger

import (
	"sync"

	"github.com/sirupsen/logrus"
)

var (
	log  *logrus.Logger
	once sync.Once
	mu   sync.Mutex
)

func Init() {
	once.Do(func() {
		log = logrus.New()
		log.SetFormatter(&logrus.JSONFormatter{
			TimestampFormat: "2006-01-02T15:04:05.999999999Z",
		})
	})
}

func SetLevel(level string) {
	mu.Lock()
	defer mu.Unlock()

	if log == nil {
		Init()
	}

	lvl, err := logrus.ParseLevel(level)
	if err != nil {
		log.WithError(err).Warn("falling back to default info level")
		return
	}

	log.SetLevel(lvl)
}

func WithService(service string) *logrus.Entry {
	mu.Lock()
	defer mu.Unlock()

	if log == nil {
		Init()
	}

	return log.WithField("service", service)
}

func init() {
	Init()
}
