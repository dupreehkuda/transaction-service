package fileKeeper

import (
	"os"
	"strings"

	"github.com/shopspring/decimal"
	"go.uber.org/zap"

	i "github.com/dupreehkuda/transaction-service/internal"
)

// GetUnprocessed gets all unprocessed requests (for startup)
func (f *fileKeeper) GetUnprocessed() []i.Job {
	f.mtx.Lock()
	defer f.mtx.Unlock()

	var result []i.Job

	content, err := os.ReadFile(f.filePath)
	if err != nil {
		f.logger.Error("Error reading file", zap.Error(err))
		return result
	}

	requests := strings.Split(string(content), "\n")

	for _, val := range requests {
		if strings.Contains(val, "processed:false") {
			values := strings.Split(val, " ")
			amount, _ := decimal.NewFromString(strings.TrimPrefix(values[3], "amount:"))

			result = append(result, i.Job{
				Id:        strings.TrimPrefix(values[0], "id:"),
				Account:   strings.TrimPrefix(values[1], "account:"),
				Operation: strings.TrimPrefix(values[2], "op:"),
				Amount:    amount,
			})
		}
	}

	return result
}
