package fileKeeper

import (
	"go.uber.org/zap"
)

type fileKeeper struct {
	filePath string
	logger   *zap.Logger
}

func (f fileKeeper) WriteNewRequest() error {
	//TODO implement me
	return nil
}

func (f fileKeeper) UpdateRequest() error {
	//TODO implement me
	return nil
}

// New creates new instance of fileKeeper
func New(filePath string, logger *zap.Logger) *fileKeeper {
	return &fileKeeper{
		filePath: filePath,
		logger:   logger,
	}
}
