package fileKeeper

import (
	"os"
	"strings"

	"go.uber.org/zap"
)

// UpdateRequest updates existing request in the .index file
func (f *fileKeeper) UpdateRequest(id string) {
	f.mtx.Lock()
	defer f.mtx.Unlock()

	content, err := os.ReadFile(f.filePath)
	if err != nil {
		f.logger.Error("Error reading file", zap.Error(err))
		return
	}

	requests := strings.Split(string(content), "\n")

	for i, val := range requests {
		if strings.Contains(val, id) {
			requests[i] = strings.Replace(val, "false", "true", 1)
		}
	}

	err = os.WriteFile(f.filePath, []byte(strings.Join(requests, "\n")), 0644)
	if err != nil {
		f.logger.Error("Error writing file", zap.Error(err))
		return
	}
}
