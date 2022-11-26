package fileKeeper

import (
	"fmt"
	"os"
	"time"

	"github.com/shopspring/decimal"
	"go.uber.org/zap"
)

// WriteNewRequest writes new unprocessed request to the .index file
func (f *fileKeeper) WriteNewRequest(account, operation string, funds decimal.Decimal) (string, error) {
	f.mtx.Lock()
	defer f.mtx.Unlock()

	file, err := os.OpenFile(f.filePath, os.O_APPEND|os.O_WRONLY, 0644)

	date := time.Now().Format(time.RFC3339)
	id := OperationHash(account, operation, date)

	record := fmt.Sprintf("id:%s account:%s op:%s amount:%s processed:%v date:%s\n", id, account, operation, funds.String(), false, date)

	_, err = file.WriteString(record)
	if err != nil {
		f.logger.Error("Error when writing file", zap.Error(err))
		return "", err
	}

	file.Close()

	return id, nil
}
