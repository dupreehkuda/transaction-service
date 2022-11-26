package fileKeeper

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"sync"

	"go.uber.org/zap"
)

type fileKeeper struct {
	filePath string
	mtx      sync.Mutex
	logger   *zap.Logger
}

// New creates new instance of fileKeeper
func New(filePath string, logger *zap.Logger) *fileKeeper {
	return &fileKeeper{
		filePath: filePath,
		mtx:      sync.Mutex{},
		logger:   logger,
	}
}

// OperationHash returns short hash string
func OperationHash(account, operation, date string) string {
	coupled := fmt.Sprintf("%s%s%s", account, operation, date)
	hsha := sha1.Sum([]byte(coupled))

	return hex.EncodeToString(hsha[:5])
}
