package fileKeeper

import (
	"os"
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
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		logger.Error("Error opening file", zap.Error(err))
	}
	file.Close()

	return &fileKeeper{
		filePath: filePath,
		mtx:      sync.Mutex{},
		logger:   logger,
	}
}
