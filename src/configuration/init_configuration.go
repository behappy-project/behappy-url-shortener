package configuration

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
	"time"
)

// 初始化log
func init() {
	/* 日志轮转相关函数
	WithLinkName(linkName), // 生成软链，指向最新日志文件
	WithMaxAge(-1), // 文件最大保存时间
	WithRotationCount(50), // 最多文件数 WithMaxAge 与WithRotationCount 二选一
	WithRotationTime(-1), // 日志切割时间间隔
	WithRotationSize(8*1024*1024), // 日志切割大小 WithRotateTime 与WithRotationSize 二选一
	*/
	logPath := "/var/log/behappy_url_shortener/shortener"
	rotationCount := uint(7)
	rotationTime := time.Duration(1)
	rotationSize := int64(10 * 1024 * 1024)
	writer, _ := rotatelogs.New(
		logPath+".%Y%m%d",
		rotatelogs.WithLinkName(logPath),
		rotatelogs.WithRotationCount(rotationCount),
		rotatelogs.WithRotationTime(rotationTime),
		rotatelogs.WithRotationSize(rotationSize),
	)
	//设置输出样式，自带的只有两种样式logrus.JSONFormatter{}和logrus.TextFormatter{}
	logrus.SetFormatter(&logrus.TextFormatter{})
	//设置output,默认为stderr,可以为任何io.Writer，比如文件*os.File
	logrus.SetOutput(writer)
	//设置最低loglevel
	logrus.SetLevel(logrus.InfoLevel)
}
