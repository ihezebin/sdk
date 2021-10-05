package excel

import "github.com/xuri/excelize/v2"

const (
	defaultMaxRow    = 2 << 12
	defaultMaxColumn = 1000
)

type Option = excelize.Options

type ReadConfig struct {
	MaxRow    int
	MaxColumn int
}

type ReadOption func(opt *ReadConfig)

func newReadConfig(options ...ReadOption) ReadConfig {
	readConfig := ReadConfig{
		MaxRow:    defaultMaxRow,
		MaxColumn: defaultMaxColumn,
	}
	for _, option := range options {
		option(&readConfig)
	}
	return readConfig
}

// WithMaxRow set the max row to read when get all rows
func WithMaxRow(maxRow int) ReadOption {
	return func(opt *ReadConfig) {
		opt.MaxRow = maxRow
	}
}

// WithMaxColumn set the max column to read when get all rows
func WithMaxColumn(maxColumn int) ReadOption {
	return func(opt *ReadConfig) {
		opt.MaxColumn = maxColumn
	}
}
