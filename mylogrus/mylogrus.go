package mylogrus

import (
	"bytes"

	"github.com/sirupsen/logrus"
)

func LogHello() string {
	var buffer bytes.Buffer
	logger := logrus.New()
	logger.Out = &buffer
	logger.Info("Hello")
	return buffer.String()
}
