package comm

import "github.com/sirupsen/logrus"

var Log = logrus.New()

func init() {
	// 配置日志
	var formatter logrus.Formatter
	formatter = &logrus.TextFormatter{
		ForceColors:     true ,
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	}
	Log.SetFormatter(formatter)
	Log.SetLevel(logrus.WarnLevel)
}

