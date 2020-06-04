package yar

import (
	"github.com/flyhope/go-yar/comm"
	"github.com/sirupsen/logrus"
)

type Level logrus.Level

const (
	LevelTrace   Level = Level(logrus.TraceLevel)
	LevelDebug   Level = Level(logrus.DebugLevel)
	LevelInfo    Level = Level(logrus.InfoLevel)
	LevelWarning Level = Level(logrus.WarnLevel)
	LevelFatal   Level = Level(logrus.FatalLevel)
	LevelPanic   Level = Level(logrus.PanicLevel)
)

// 设置日志输出级别
func SetLevel(level Level) {
	comm.Log.SetLevel(logrus.Level(level))
}
