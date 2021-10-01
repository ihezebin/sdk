package excel

import (
	"github.com/pkg/errors"
	"github.com/xuri/excelize/v2"
	"io"
	"sync"
)

// Excel excel
type Excel struct {
	kernel    *excelize.File
	sizeCache *sync.Map // cache sheet size for later use.
}

// Size record for excel sheet size
type Size struct {
	Row    int
	Column int
}

// New create a new excel
func New() *Excel {
	return &Excel{kernel: excelize.NewFile(), sizeCache: new(sync.Map)}
}

func OpenFile(filename string, opt ...Options) (*Excel, error) {
	file, err := excelize.OpenFile(filename, opt...)
	if err != nil {
		return nil, err
	}
	return &Excel{kernel: file, sizeCache: new(sync.Map)}, nil
}

func OpenReader(r io.Reader, opt ...Options) (*Excel, error) {
	file, err := excelize.OpenReader(r, opt...)
	if err != nil {
		return nil, err
	}
	return &Excel{kernel: file, sizeCache: new(sync.Map)}, nil
}

// Rows return rows by sheet index
func (e *Excel) Rows(index int) (*excelize.Rows, error) {
	sheet := e.kernel.GetSheetName(index)
	return e.kernel.Rows(sheet)
}

// Size returns the row count and max column count of a sheet
func (e *Excel) Size(index int) (row int, column int, err error) {
	sheet := e.kernel.GetSheetName(index)
	if val, ok := e.sizeCache.Load(sheet); ok {
		size := val.(Size)
		return size.Row, size.Column, nil
	}
	row, column, err = e.getSheetSize(sheet)
	if err != nil {
		return 0, 0, err
	}
	e.sizeCache.Store(sheet, Size{Row: row, Column: column})
	return
}

func (e *Excel) getSheetSize(sheet string) (row int, column int, err error) {
	rows, err := e.kernel.Rows(sheet)
	if err != nil {
		return 0, 0, err
	}
	for rows.Next() {
		cols, err := rows.Columns()
		if err != nil {
			return 0, 0, err
		}
		if column < len(cols) {
			column = len(cols)
			cols = cols[:0] // truncate cols
		}
		row++
	}
	return
}

// GetRows get all rows with row and column limit.
// Default Row limit is 8192. when exceed will set hasMore=true.
// Default Column limit is 1000. when exceed will cause ErrTooManyColumn
func (e *Excel) GetRows(index int, options ...ReadOption) ([][]string, bool, error) {
	config := &ReadConfig{
		MaxRow:    defaultMaxRow,
		MaxColumn: defaultMaxColumn,
	}
	for _, option := range options {
		option(config)
	}
	rows, err := e.Rows(index)
	if err != nil {
		return nil, false, err
	}

	results := make([][]string, 0, 64)
	count := 0
	hasMore := false
	for rows.Next() {
		count++
		if count > config.MaxRow {
			hasMore = true
			break
		}
		row, err := rows.Columns()
		if err != nil {
			return nil, hasMore, err
		}
		if len(row) > config.MaxColumn {
			return nil, hasMore, errors.New("too many column in row")
		}
		results = append(results, row)
	}
	return results, hasMore, nil
}
