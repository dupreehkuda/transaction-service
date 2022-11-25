package fileKeeper

import (
	"go.uber.org/zap"
)

type fileKeeper struct {
	filePath string
	logger   *zap.Logger
}

func (f fileKeeper) WriteNewRequest() {
	//TODO implement me
	panic("implement me")
}

func (f fileKeeper) UpdateRequest() {
	//TODO implement me
	panic("implement me")
}

// New creates new instance of fileKeeper
func New(filePath string, logger *zap.Logger) *fileKeeper {
	return &fileKeeper{
		filePath: filePath,
		logger:   logger,
	}
}
