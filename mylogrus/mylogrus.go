package mylogrus

import (
	"bytes"

	"github.com/sirupsen/logrus"
)

func LogString(str string) string {
	var buffer bytes.Buffer
	logger := logrus.New()
	logger.Out = &buffer
	logger.Info(str)
	return buffer.String()
}
