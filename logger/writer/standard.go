package writer

import (
	"io"
	"os"
)

func StandardErrorWriter() io.Writer {
	return os.Stderr
}

func StandardOutWriter() io.Writer {
	return os.Stdout
}

func Default() io.Writer {
	return StandardErrorWriter()
}
