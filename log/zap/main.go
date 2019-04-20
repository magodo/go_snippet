package main

import (
	"go.uber.org/zap"
)

func main() {
	tryPresets()
}

func tryPresets() {
	l := zap.NewExample()
	l.Info("example info", zap.String("foo", "bar"))

	l, _ = zap.NewDevelopment()
	l.Info("development info", zap.String("foo", "bar"))

	l, _ = zap.NewProduction()
	l.Info("production info", zap.String("foo", "bar"))

	s := l.Sugar()
	s.Infof("log from %s", "sugar")
}
