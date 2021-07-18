package logger

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"io"
	"os"
	"time"
)

func OStdErrWriter() io.Writer {
	return os.Stderr
}

func OStdWriter() io.Writer {
	return os.Stdout
}

func RotateFileWriter(filename string, rotateTime time.Duration, expireTime time.Duration) io.Writer {
	writer, err := rotatelogs.New(
		// The pattern used to generate actual log file names.
		// You should use patterns using the strftime (3) format.
		// 分割后的文件名称
		filename+".%Y%m%d%H%M%S",
		// Interval between file rotation. By default logs are rotated every 86400 seconds.
		// Note: Remember to use time.Duration values.
		// 设置日志切割时间间隔
		rotatelogs.WithRotationTime(rotateTime),
		// Path where a symlink for the actual log file is placed.
		// This allows you to always check at the same location for log files even if the logs were rotated
		// 生成软链，指向最新日志文件
		//rotatelogs.WithLinkName("./"),
		// Time to wait until old logs are purged. By default no logs are purged,
		// which certainly isn't what you want. Note: Remember to use time.Duration values.
		// 设置最大保存时间(7天)
		rotatelogs.WithMaxAge(expireTime),
		// The number of files should be kept. By default, this option is disabled.
		// Note: MaxAge should be disabled by specifing WithMaxAge(-1) explicitly.
		//rotatelogs.WithRotationCount(1),
		// Ensure a new file is created every time New() is called.
		// If the base file name already exists, an implicit rotation is performed.
		rotatelogs.ForceNewFile(),
	)
	if err != nil {
		panic("rotate file writer create failed")
	}
	return writer
}

func FileWriter(filename string) io.Writer {
	return nil
}
