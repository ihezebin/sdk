package writer

import (
	"fmt"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"io"
	"time"
)

const (
	defaultRotateTime       = time.Hour * 24
	defaultRotateExpireTime = time.Hour * 24 * 3
)

type rotateFileWriter struct {
	kernel     *rotatelogs.RotateLogs
	filename   string
	rotateTime time.Duration
	expireTime time.Duration
}

func DefaultRotateFileWriter(filename string) io.Writer {
	return createRotateFileWriter(filename, defaultRotateTime, defaultRotateExpireTime)
}

// NewRotateFileWriter Rotate output and split local log files to avoid accumulation of excessive log content in a log file
// 轮转输出, 分割本地日志文件, 避免一个日志文件中堆积过多日志内容
func NewRotateFileWriter(filename string, rotateTime time.Duration, expireTime time.Duration) io.Writer {
	return createRotateFileWriter(filename, rotateTime, expireTime)
}

/*  createRotateFileWriter
WithLinkName 为最新的日志建立软连接
WithRotationTime 设置日志分割的时间，隔多久分割一次
WithMaxAge 和 WithRotationCount二者只能设置一个
  WithMaxAge 设置文件清理前的最长保存时间
  WithRotationCount 设置文件清理前最多保存的个数
*/
func createRotateFileWriter(filename string, rotateTime time.Duration, expireTime time.Duration) io.Writer {
	kernel, err := rotatelogs.New(
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
		//rotatelogs.WithLinkName(filename),
		// Time to wait until old logs are purged. By default no logs are purged,
		// which certainly isn't what you want. Note: Remember to use time.Duration values.
		// 设置最大保存时间
		rotatelogs.WithMaxAge(expireTime),
		// The number of files should be kept. By default, this option is disabled.
		// Note: MaxAge should be disabled by specifing WithMaxAge(-1) explicitly.
		//rotatelogs.WithRotationCount(1),
		// Ensure a new file is created every time New() is called.
		// If the base file name already exists, an implicit rotation is performed.
		rotatelogs.ForceNewFile(),
	)
	if err != nil {
		panic(fmt.Sprintf("create rotate file writer failed: %s", err.Error()))
	}
	return &rotateFileWriter{kernel, filename, rotateTime, expireTime}
}

func (writer *rotateFileWriter) Write(p []byte) (n int, err error) {
	return writer.kernel.Write(p)
}
